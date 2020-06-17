package main

import (
	"fmt"
	"os"

	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/grid"
	"github.com/gcla/gowid/widgets/styled"
	"github.com/gcla/gowid/widgets/text"
)

func main() {
	palette := gowid.Palette{
		"green": gowid.MakePaletteEntry(gowid.ColorWhite, gowid.ColorGreen),
	}

	widgets := make([]gowid.IWidget, 0)
	for i := 0; i < 20; i++ {
		widgets = append(
			widgets,
			styled.New(
				text.New(fmt.Sprintf("%d", i), text.Options{Align: gowid.HAlignMiddle{}}),
				gowid.MakePaletteRef("green"),
			),
		)
	}

	grd := grid.New(
		widgets,
		20, // column width
		2,  // width between columns
		1,  // width between rows
		gowid.HAlignMiddle{})

	app, err := gowid.NewApp(gowid.AppArgs{
		View:    grd,
		Palette: palette,
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	app.SimpleMainLoop()
}
