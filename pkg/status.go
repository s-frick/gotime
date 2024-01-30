package gotime

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/s-frick/go-time-track/pkg/internal/color"
)

func Status(ctx context.Context) {
	gt := readState(ctx)
	if gt != nil && gt.state.Start > 0 {
		f := gt.generateFrame(time.Now())
		s := stats(f)

		fmt.Printf("Frame %s started %s ago. (%s)\n", gt.state.Tags, s.sinceStarted.String(), color.Blue().Sprint(f.Start.Format("02.01.2006 15:04")))
		os.Exit(0)
	}
	fmt.Printf("No frame started.\n")
	os.Exit(0)
}
