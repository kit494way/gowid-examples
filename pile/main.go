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
		"green":     gowid.MakePaletteEntry(gowid.ColorWhite, gowid.ColorGreen),
		"red":       gowid.MakePaletteEntry(gowid.ColorWhite, gowid.ColorRed),
		"blue":      gowid.MakePaletteEntry(gowid.ColorWhite, gowid.ColorBlue),
		"ligtGreen": gowid.MakePaletteEntry(gowid.ColorWhite, gowid.ColorLightGreen),
	}

	flow := gowid.RenderFlow{}

	pl := pile.New([]gowid.IContainerWidget{
		&gowid.ContainerWidget{IWidget: styled.New(text.New("pile"), gowid.MakePaletteRef("green")), D: flow},
		&gowid.ContainerWidget{IWidget: styled.New(text.New("is"), gowid.MakePaletteRef("red")), D: flow},
		&gowid.ContainerWidget{IWidget: styled.New(text.New("vertical"), gowid.MakePaletteRef("blue")), D: flow},
		&gowid.ContainerWidget{IWidget: styled.New(text.New("stack"), gowid.MakePaletteRef("ligtGreen")), D: flow},
	})

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
