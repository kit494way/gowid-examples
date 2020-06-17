package main

import (
	"fmt"
	"os"

	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/list"
	"github.com/gcla/gowid/widgets/selectable"
	"github.com/gcla/gowid/widgets/styled"
	"github.com/gcla/gowid/widgets/text"
)

func main() {
	palette := gowid.Palette{
		"focus": gowid.MakePaletteEntry(gowid.ColorBlack, gowid.ColorGreen),
	}

	// text.Widget is not selectable.
	// selectable.Widget makes non-selectable widget selectable.
	widgets := []gowid.IWidget{
		styled.NewFocus(selectable.New(text.New("one")), gowid.MakePaletteRef("focus")),
		styled.NewFocus(selectable.New(text.New("two")), gowid.MakePaletteRef("focus")),
		styled.NewFocus(selectable.New(text.New("three")), gowid.MakePaletteRef("focus")),
	}

	// Selected item of the list can be changed by pressing Down, Up, Ctrl-N or Ctrl-P.
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
