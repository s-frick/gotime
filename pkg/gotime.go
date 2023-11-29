package gotime

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/s-frick/go-time-track/pkg/internal/color"
	"github.com/s-frick/go-time-track/pkg/internal/utils"
)

const (
	truncToMinute = (1 * time.Minute)
	dformat       = "15 h and 4 m"
)

var (
	bigBang = time.Unix(0, 0).UTC()
)

type State struct {
	Start int64 `json:"start"`
	Tags  []Tag `json:"tags"`
}

type Tag string

func (t Tag) String() string {
	return color.Cyan().Sprint(string(t))
}

type Time struct {
	*time.Time
}

func (t Time) String() string {
	return color.Green().Sprint(t.Time.Format("15:04"))
}

type ID string

type GoTime struct {
	state State
	tags  []Tag
}

type Frame struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	Tags  []Tag     `json:"tags"`
	ID    ID        `json:"id"`
}

type Frames []Frame

type Options struct {
	At string
}

type InternOption struct {
	At time.Time
}

func readState(ctx context.Context) *GoTime {
	home := ctx.Value("gotimeDir")
	stateFile := fmt.Sprintf("%s/state", home)

	rawState, err := os.ReadFile(stateFile)
	var oldState State
	if err != nil {
		slog.Debug("reading state file %s: %s", stateFile, err)
	}
	err = json.Unmarshal(rawState, &oldState)
	if err != nil {
		slog.Debug("unmarshal state, rawState %s: %s", string(rawState), err)
		oldState = State{}
	}

	rawTags, err := os.ReadFile(fmt.Sprintf("%s/tags", home))
	var tags []Tag
	if err != nil {
		tags = make([]Tag, 0)
	}
	err = json.Unmarshal(rawTags, &tags)
	if err != nil {
		fmt.Printf("can't unmarshal tags file: %s", err)
	}

	g := &GoTime{
		state: oldState,
		tags:  tags,
	}

	return g
}

func saveToFile(ctx context.Context, file string, v any) {
	home := ctx.Value("gotimeDir")
	path := fmt.Sprintf("%s/%s", home, file)

	stateJson, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = f.Write(stateJson)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func removeState(ctx context.Context) {
	home := ctx.Value("gotimeDir")
	path := fmt.Sprintf("%s/%s", home, "state")
	err := os.Remove(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Start(ctx context.Context, t []Tag, opt Options) {
	g := readState(ctx)

	gt := time.Now()
	if opt.At != "" {
		gt = parseAtOption(opt.At)
	}

	var newTags []Tag
	var knownTags []Tag
outer:
	for _, newTag := range t {
		for _, tag := range g.tags {
			if tag == newTag {
				knownTags = append(knownTags, tag)
				continue outer
			}
		}
		if utils.Confirmation(utils.ConfirmationRequest(fmt.Sprintf("Create new tag? \"%s\"", newTag))) {
			newTags = append(newTags, newTag)
		} else {
			newTag = Tag(utils.Ask(utils.Question(fmt.Sprintf("What should we \"%s\" replace with? ", newTag))))
			newTags = append(newTags, newTag)
		}
	}

	if g.state.Start != 0 {
		f := g.stop(ctx, InternOption{At: gt})
		gt = f.End
	}

	state := State{
		Start: gt.Unix(),
		Tags:  append(knownTags, newTags...),
	}
	g.state = state

	saveToFile(ctx, "state", g.state)
	saveToFile(ctx, "tags", append(g.tags, newTags...))

	fmt.Printf("Starting frame %s at %s\n", fmt.Sprint(append(knownTags, newTags...)), Time{Time: &gt})
}

func parseAtOption(atStr string) time.Time {
	at, err := time.ParseInLocation("15:04", atStr, time.Local)
	if err != nil {
		return time.Now()
	}
	y, m, d := time.Now().Date()

	return at.AddDate(y, int(m)-1, d-1)
}

func Stop(ctx context.Context, opt Options) {
	at := parseAtOption(opt.At)
	g := readState(ctx)
	g.stop(ctx, InternOption{At: at})
}

func generateID(f Frame) string {
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%s%s%s", f.Start, f.End, f.Tags)))
	return hex.EncodeToString(h.Sum(nil))
}

type Timespan time.Duration

func (t Timespan) Format(format string) string {
	return time.Unix(0, 0).UTC().Add(time.Duration(t)).Format(format)
}
func (t Timespan) String() string {
	return color.Green().Sprint(t.Format(dformat))
}

func (g *GoTime) stop(ctx context.Context, opt InternOption) Frame {
	end := opt.At
	f := Frame{Start: time.Unix(g.state.Start, 0).Truncate(truncToMinute), End: end.Truncate(truncToMinute), Tags: g.state.Tags}
	id := generateID(f)
	f.ID = ID(id)

	s := f.Start
	e := f.End
	d := e.Sub(s)
	now := time.Unix(time.Now().Unix(), 0)
	sinceStarted := now.Sub(s)

	saveFrame(ctx, f)
	removeState(ctx)
	sinceStartedFmt := Timespan(sinceStarted).String()
	dFmt := Timespan(d).String()

	fmt.Printf("Stopping frame %s at %s, started %s ago and lasted %s minutes.\n", fmt.Sprint(f.Tags), Time{Time: &f.End}, sinceStartedFmt, dFmt)
	return f
}

func saveFrame(ctx context.Context, f Frame) {
	home := ctx.Value("gotimeDir")
	framesFile := fmt.Sprintf("%s/frames", home)

	rawFrames, err := os.ReadFile(framesFile)
	var frames Frames
	if err != nil {
		slog.Debug("reading frames file %s: %s", framesFile, err)
	}
	err = json.Unmarshal(rawFrames, &frames)
	if err != nil {
		slog.Debug("unmarshal frames, rawState %s: %s", string(rawFrames), err)
		frames = make([]Frame, 0)
	}

	frames = append(frames, f)

	saveToFile(ctx, "frames", frames)
}
