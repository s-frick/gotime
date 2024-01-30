package gotime

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/s-frick/go-time-track/pkg/internal/color"
)

type JsonLogFrame struct {
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	ID       ID        `json:"id"`
	Tags     []Tag     `json:"tags"`
	Duration float64   `json:"duration"`
}

func NewJsonLogFrame(frame Frame) *JsonLogFrame {
	stats := stats(frame)

	d := time.Duration.Seconds(time.Duration(stats.duration))
	return &JsonLogFrame{
		Start:    frame.Start,
		End:      frame.End,
		ID:       frame.ID,
		Tags:     frame.Tags,
		Duration: d,
	}
}

type CsvLogFrame struct {
	Start    time.Time `csv:"start"`
	End      time.Time `csv:"end"`
	ID       ID        `csv:"id"`
	Tags     string    `csv:"tags"`
	Duration float64   `csv:"duration"`
}

func NewCsvLogFrame(frame Frame) *CsvLogFrame {
	stats := stats(frame)

	var sb strings.Builder
	for i, tag := range frame.Tags {
		sb.WriteString(string(tag))
		if i != len(frame.Tags)-1 {
			sb.WriteString(" ")
		}
	}

	d := time.Duration.Seconds(time.Duration(stats.duration))
	return &CsvLogFrame{
		Start:    frame.Start,
		End:      frame.End,
		ID:       frame.ID,
		Tags:     sb.String(),
		Duration: d,
	}
}

type LogType int

const (
	PrettyLog LogType = iota
	JsonLog
	CsvLog
)

func Log(ctx context.Context) {
	frames := readFrames(ctx)

	switch GetLogType(ctx) {
	case PrettyLog:
		prettyLog(frames)
	case JsonLog:
		jsonLog(frames)
	case CsvLog:
		csvLog(frames)
	}
}

func jsonLog(frames Frames) {
	for _, frame := range frames {
		out := NewJsonLogFrame(frame)

		stateJson, err := json.Marshal(out)
		if err != nil {
			log.Fatalf("Failed writing json log, %s", err)
		}

		_, err = os.Stdout.Write(stateJson)
		os.Stdout.WriteString("\n")
		if err != nil {
			log.Fatalf("Failed writing json log, %s", err)
		}
	}
}

func csvLog(frames Frames) {
	fs := make([]*CsvLogFrame, 0)
	for _, frame := range frames {
		fs = append(fs, NewCsvLogFrame(frame))
	}

	err := gocsv.MarshalFile(fs, os.Stdout)
	if err != nil {
		log.Fatalf("Failed writing csv log, %s", err)
	}
}

func writePrettyFrameSet(cur time.Time, sum time.Duration, log string) {
	fmt.Printf("\n%s%s%s%s\n", strings.Repeat(color.Red().Sprint("_"), 12), color.Red().Sprint("[ ", cur.Format(dateFormat), " ]"), strings.Repeat(color.Red().Sprint("_"), 12), color.Red().Sprint("[ ", sum.String(), " ]"))
	fmt.Print(log)
}

func prettyLog(frames Frames) {
	cur := bigBang
	var log strings.Builder
	sum := time.Duration(0)
	for i, frame := range frames {
		if cur.Equal(bigBang) {
			cur = frame.Start.Truncate(24 * time.Hour)
		}

		if cur != frame.Start.Truncate(24*time.Hour) {
			writePrettyFrameSet(cur, sum, log.String())

			cur = bigBang
			log.Reset()
			sum = time.Duration(0)
		}

		stats := stats(frame)
		start := frame.Start.Format(timeFormat)
		end := frame.End.Format(timeFormat)

		log.WriteString(color.Gray().Sprintf("  (%s)", frame.ID[:7]))
		log.WriteString("      " + color.Purple().Sprint(start))
		log.WriteString("  ~>  ")
		log.WriteString(color.Purple().Sprint(end) + "    ")
		log.WriteString(stats.duration.String())
		log.WriteString(fmt.Sprintf("    %s", frame.Tags))
		log.WriteString("\n")

		sum = sum + time.Duration(stats.duration)

		if i == len(frames)-1 {
			writePrettyFrameSet(cur, sum, log.String())
			fmt.Println()
		}
	}
}
