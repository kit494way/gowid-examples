package main

import (
	"fmt"
	"os"

	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/list"
	"github.com/gcla/gowid/widgets/styled"
	"github.com/gcla/gowid/widgets/text"
)

func main() {
	palette := gowid.Palette{
		"red":   gowid.MakePaletteEntry(gowid.ColorRed, gowid.ColorBlack),
		"green": gowid.MakePaletteEntry(gowid.ColorGreen, gowid.ColorBlack),
		"blue":  gowid.MakePaletteEntry(gowid.ColorBlue, gowid.ColorBlack),
	}

	widgets := []gowid.IWidget{
		styled.New(text.New("red"), gowid.MakePaletteRef("red")),
		styled.New(text.New("green"), gowid.MakePaletteRef("green")),
		styled.New(text.New("blue"), gowid.MakePaletteRef("blue")),
	}

	lst := list.New(list.NewSimpleListWalker(widgets))

	app, err := gowid.NewApp(gowid.AppArgs{
		View:    lst,
		Palette: palette,
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	app.SimpleMainLoop()
}
