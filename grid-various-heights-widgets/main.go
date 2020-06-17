package main

import (
	"fmt"
	"os"

	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/grid"
	"github.com/gcla/gowid/widgets/pile"
	"github.com/gcla/gowid/widgets/styled"
	"github.com/gcla/gowid/widgets/text"
)

func main() {
	palette := gowid.Palette{
		"green": gowid.MakePaletteEntry(gowid.ColorWhite, gowid.ColorGreen),
	}

	// Height of each widget is between 1 and 5.
	widgets := make([]gowid.IWidget, 0)
	for i := 1; i <= 20; i++ {
		n := i % 5
		if n == 0 {
			n = 5
		}
		ws := make([]gowid.IContainerWidget, 0)
		for j := 1; j <= n; j++ {
			ws = append(ws,
				&gowid.ContainerWidget{
					IWidget: styled.New(
						text.New(fmt.Sprintf(""), text.Options{Align: gowid.HAlignMiddle{}}),
						gowid.MakePaletteRef("green"),
					),
					D: gowid.RenderFlow{},
				})
		}
		widgets = append(widgets, pile.New(ws))
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
