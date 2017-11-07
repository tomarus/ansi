package ansi

import (
	"fmt"
	"testing"
)

func TestAnsi(t *testing.T) {
	fmt.Println(Clear, Top)
	fmt.Println(Color(White, Red) + "Color " + Inverse + " Test" + Reset)
	fmt.Println(Underline + "Color Palette" + Reset + " format " + Bold + "Foreground,Background" + Reset)
	for x := 0; x < 8; x++ {
		fmt.Printf("%s %d ", Color(x), x)
		for y := 0; y < 8; y++ {
			fmt.Printf("%s %d,%d  ", Color(x, y), x, y)
		}
		fmt.Println(Reset)
	}
	for x := 0; x < 8; x++ {
		fmt.Print("   ")
		for y := 60; y < 68; y++ {
			fmt.Printf("%s %d,%d ", Color(x, y), x, y)
		}
		fmt.Println(Reset)
	}
}
