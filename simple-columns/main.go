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

	fixed := gowid.RenderFixed{}

	widgets := []gowid.IContainerWidget{
		&gowid.ContainerWidget{
			IWidget: styled.New(text.New(" one "), gowid.MakePaletteRef("green")),
			D:       fixed,
		},
		&gowid.ContainerWidget{
			IWidget: styled.New(text.New(" two "), gowid.MakePaletteRef("red")),
			D:       fixed,
		},
		&gowid.ContainerWidget{
			IWidget: styled.New(text.New(" three "), gowid.MakePaletteRef("green")),
			D:       fixed,
		},
	}

	cols := columns.New(widgets)

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
