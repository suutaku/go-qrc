package qrc

import (
	"os"

	"github.com/fumiyas/go-tty"
	"github.com/mattn/go-colorable"
	"github.com/qpliu/qrencode-go/qrencode"
)

func ShowQR(text string, inverse bool) error {
	grid, err := qrencode.Encode(text, qrencode.ECLevelL)
	if err != nil {
		return err
	}
	da1, err := tty.GetDeviceAttributes1(os.Stdout)
	if err == nil && da1[tty.DA1_SIXEL] {
		PrintSixel(os.Stdout, grid, inverse)
	} else {
		stdout := colorable.NewColorableStdout()
		PrintAA(stdout, grid, inverse)
	}
	return nil
}
