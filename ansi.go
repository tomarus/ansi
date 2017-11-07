package ansi

// https://en.wikipedia.org/wiki/ANSI_escape_code
// http://ascii-table.com/ansi-escape-sequences-vt-100.php

import "fmt"

const (
	Black = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// only for bg
const (
	LightBlack = iota + 60
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightMagenta
	LightCyan
	LightWhite
)

const Escape = "\033["

const Reset = Escape + "0m"
const Bold = Escape + "1m"
const Underline = Escape + "4m"
const Inverse = Escape + "7m"

const ClearUp = Escape + "J"
const ClearDown = Escape + "1J"
const Clear = Escape + "2J"

const Top = Escape + "f"

func Fg(color int) string { return fmt.Sprintf("%s%dm", Escape, 30+color) }
func Bg(color int) string { return fmt.Sprintf("%s%dm", Escape, 40+color) }
func Y(y int) string      { return fmt.Sprintf("%s%df", Escape, y) }
func X(x int) string      { return fmt.Sprintf("%s%dC", Escape, x) }
func Pos(x, y int) string { return fmt.Sprintf("%s%df%s%dC", Escape, y, Escape, x) }
func UpN(n int) string    { return fmt.Sprintf("%s%dA", Escape, n) }
func DownN(n int) string  { return fmt.Sprintf("%s%dB", Escape, n) }
func RightN(n int) string { return fmt.Sprintf("%s%dC", Escape, n) }
func LeftN(n int) string  { return fmt.Sprintf("%s%dD", Escape, n) }

func Color(fg int, bgopt ...int) string {
	if len(bgopt) == 1 {
		bg := bgopt[0]
		if bg >= 0 {
			return fmt.Sprintf("%s%d;%dm", Escape, 30+fg, 40+bg)
		}
	}
	return fmt.Sprintf("%s%dm", Escape, 30+fg)
}
