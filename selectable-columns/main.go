package main

import (
	"fmt"
	"os"

	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/columns"
	"github.com/gcla/gowid/widgets/selectable"
	"github.com/gcla/gowid/widgets/styled"
	"github.com/gcla/gowid/widgets/text"
)

func main() {
	palette := gowid.Palette{
		"focus": gowid.MakePaletteEntry(gowid.ColorBlack, gowid.ColorGreen),
	}

	fixed := gowid.RenderFixed{}

	// text.Widget is not selectable.
	// selectable.Widget makes non-selectable widget selectable.
	widgets := []gowid.IContainerWidget{
		&gowid.ContainerWidget{
			IWidget: styled.NewFocus(selectable.New(text.New(" one ")), gowid.MakePaletteRef("focus")),
			D:       fixed,
		},
		&gowid.ContainerWidget{
			IWidget: styled.NewFocus(selectable.New(text.New(" two ")), gowid.MakePaletteRef("focus")),
			D:       fixed,
		},
		&gowid.ContainerWidget{
			IWidget: styled.NewFocus(selectable.New(text.New(" three ")), gowid.MakePaletteRef("focus")),
			D:       fixed,
		},
	}

	// Selected widget of the columns can be changed by pressing Left or Right.
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
