package main

import (
	"fmt"
	"os"

	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/pile"
	"github.com/gcla/gowid/widgets/styled"
	"github.com/gcla/gowid/widgets/text"
)

func main() {
	palette := gowid.Palette{
		"green": gowid.MakePaletteEntry(gowid.ColorWhite, gowid.ColorGreen),
		"red":   gowid.MakePaletteEntry(gowid.ColorWhite, gowid.ColorRed),
	}

	stxt1 := styled.New(text.New("one"), gowid.MakePaletteRef("green"))
	stxt2 := styled.New(text.New("two"), gowid.MakePaletteRef("red"))

	pl := pile.NewWithDim(gowid.RenderWithWeight{W: 1}, stxt1, stxt2)
	/* Above code, `pl := pile.NewWithDim(...)`, is same as follows. */
	// pl := pile.New([]gowid.IContainerWidget{
	// 	&gowid.ContainerWidget{IWidget: stxt1, D: gowid.RenderWithWeight{W: 1}},
	// 	&gowid.ContainerWidget{IWidget: stxt2, D: gowid.RenderWithWeight{W: 1}},
	// })

	app, err := gowid.NewApp(gowid.AppArgs{
		View:    pl,
		Palette: palette,
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	app.SimpleMainLoop()
}
