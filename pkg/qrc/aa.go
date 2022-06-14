package qrc

import (
	"bufio"
	"fmt"
	"io"

	"github.com/mgutz/ansi"
	"github.com/qpliu/qrencode-go/qrencode"
	"golang.org/x/term"
)

func PrintAA(w_in io.Writer, grid *qrencode.BitGrid, inverse bool) {
	// Buffering required for Windows (go-colorable) support
	w := bufio.NewWriterSize(w_in, 1024)

	reset := ansi.ColorCode("reset")
	black := ansi.ColorCode(":black")
	white := ansi.ColorCode(":white")
	if inverse {
		black, white = white, black
	}

	height := grid.Height()
	width := grid.Width()
	line := white + fmt.Sprintf("%*s", width*2+2, "") + reset + "\n"
	ww, _, _ := term.GetSize(0)
	shift := (ww - width) / 2 / 8 // as a tap
	shiftStr := ""
	for i := 0; i < shift-1; i++ {
		shiftStr += "\t"
	}
	fmt.Fprint(w, shiftStr, line)
	for y := 0; y < height; y++ {
		fmt.Fprint(w, shiftStr, white, " ")
		color_prev := white
		for x := 0; x < width; x++ {
			if grid.Get(x, y) {
				if color_prev != black {
					fmt.Fprint(w, black)
					color_prev = black
				}
			} else {
				if color_prev != white {
					fmt.Fprint(w, white)
					color_prev = white
				}
			}
			fmt.Fprint(w, "  ")
		}
		fmt.Fprint(w, white, " ", reset, "\n")
		w.Flush()
	}
	fmt.Fprint(w, shiftStr, line)
	w.Flush()
}
