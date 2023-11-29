package color

import (
	"fmt"
	"runtime"
)

var (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
	white  = "\033[97m"
)

func init() {
	fmt.Print()
	if runtime.GOOS == "windows" {
		reset = ""
		red = ""
		green = ""
		yellow = ""
		blue = ""
		purple = ""
		cyan = ""
		gray = ""
		white = ""
	}
}

type TermColor interface {
	Sprintf(format string, a ...any) string

	Sprintln(a ...any) string

	Sprint(a ...any) string

	Printf(format string, a ...any) (n int, err error)

	// Println(a ...any) (n int, err error)
	//
	// Print(a ...any) (n int, err error)
}

type Colorizer struct {
	color string
}

func (c Colorizer) Sprintf(format string, a ...any) string {
	cFormat := fmt.Sprintf("%s%s%s", c.color, format, reset)
	return fmt.Sprintf(cFormat, a...)
}
func (c Colorizer) Sprintln(a ...any) string {
	strings := fmt.Sprintf("%s%s%s", c.color, fmt.Sprint(a...), reset)
	return fmt.Sprintln(strings)
}

func (c Colorizer) Printf(format string, a ...any) (n int, err error) {
	cFormat := fmt.Sprintf("%s%s%s", c.color, format, reset)
	return fmt.Printf(cFormat, a...)
}
func (c Colorizer) Sprint(a ...any) string {
	strings := fmt.Sprintf("%s%s%s", c.color, fmt.Sprint(a...), reset)
	return fmt.Sprint(strings)
}

func Red() TermColor {
	return Colorizer{color: red}
}
func Green() TermColor {
	return Colorizer{color: green}
}
func Yellow() TermColor {
	return Colorizer{color: yellow}
}
func Blue() TermColor {
	return Colorizer{color: blue}
}
func Purple() TermColor {
	return Colorizer{color: purple}
}
func Cyan() TermColor {
	return Colorizer{color: cyan}
}
func Gray() TermColor {
	return Colorizer{color: gray}
}
func White() TermColor {
	return Colorizer{color: white}
}
