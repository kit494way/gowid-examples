package main

import (
	"fmt"
	"os"

	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/list"
	"github.com/gcla/gowid/widgets/pile"
	"github.com/gcla/gowid/widgets/selectable"
	"github.com/gcla/gowid/widgets/styled"
	"github.com/gcla/gowid/widgets/text"
)

func main() {
	palette := gowid.Palette{
		"default": gowid.MakePaletteEntry(gowid.ColorWhite, gowid.ColorBlack),
		"focus":   gowid.MakePaletteEntry(gowid.ColorBlack, gowid.ColorGreen),
	}

	txt := text.New("")
	stxt := styled.New(txt, gowid.MakePaletteRef("default"))

	// text.Widget is not selectable.
	// selectable.Widget makes non-selectable widget selectable.
	// Selected item of the list can be changed by pressing Down, Up, Ctrl-N or Ctrl-P.
	widgets := []gowid.IWidget{
		styled.NewFocus(selectable.New(text.New("one")), gowid.MakePaletteRef("focus")),
		styled.NewFocus(selectable.New(text.New("two")), gowid.MakePaletteRef("focus")),
		styled.NewFocus(selectable.New(text.New("three")), gowid.MakePaletteRef("focus")),
	}

	lst := list.New(list.NewSimpleListWalker(widgets))

	// This is called when selected item of the list is changed.
	lst.OnFocusChanged(gowid.WidgetCallback{Name: "focus", WidgetChangedFunction: func(app gowid.IApp, widget gowid.IWidget) {
		tw := widget.(*styled.Widget).SubWidget().(*selectable.Widget).SubWidget().(*text.Widget)
		txt.SetContent(app, tw.Content())
	}})

	pl := pile.New([]gowid.IContainerWidget{
		&gowid.ContainerWidget{IWidget: stxt, D: gowid.RenderFlow{}},
		&gowid.ContainerWidget{IWidget: lst, D: gowid.RenderFlow{}},
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
