package main

import (
	"fmt"
	"os"

	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/columns"
	"github.com/gcla/gowid/widgets/styled"
	"github.com/gcla/gowid/widgets/text"
)

func main() {
	palette := gowid.Palette{
		"green": gowid.MakePaletteEntry(gowid.ColorWhite, gowid.ColorGreen),
		"red":   gowid.MakePaletteEntry(gowid.ColorWhite, gowid.ColorRed),
	}

	cols := columns.NewWithDim(gowid.RenderWithWeight{W: 1},
		styled.New(text.New("one"), gowid.MakePaletteRef("green")),
		styled.New(text.New("two"), gowid.MakePaletteRef("red")),
	)

	/* Above code, `cols : = columns.NewWithDim(...)`, is same as follows. */
	// widgets := []gowid.IContainerWidget{
	// 	&gowid.ContainerWidget{
	// 		IWidget: styled.New(text.New("one"), gowid.MakePaletteRef("green")),
	// 		D:       gowid.RenderWithWeight{W: 1},
	// 	},
	// 	&gowid.ContainerWidget{
	// 		IWidget: styled.New(text.New("two"), gowid.MakePaletteRef("red")),
	// 		D:       gowid.RenderWithWeight{W: 1},
	// 	},
	// }
	// cols := columns.New(widgets)

	app, err := gowid.NewApp(gowid.AppArgs{
		View:    cols,
		Palette: palette,
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	app.SimpleMainLoop()
}
