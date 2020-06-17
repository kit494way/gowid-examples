package main

import (
	"fmt"
	"os"

	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/columns"
	"github.com/gcla/gowid/widgets/holder"
	"github.com/gcla/gowid/widgets/list"
	"github.com/gcla/gowid/widgets/pile"
	"github.com/gcla/gowid/widgets/selectable"
	"github.com/gcla/gowid/widgets/styled"
	"github.com/gcla/gowid/widgets/text"
)

type Tab struct {
	key string
	gowid.IContainerWidget
}

func NewTab(key string, style gowid.ICellStyler) *Tab {
	c := &gowid.ContainerWidget{
		IWidget: styled.NewFocus(selectable.New(text.New(fmt.Sprintf(" %s ", key))), style),
		D:       gowid.RenderWithWeight{W: 1},
	}
	return &Tab{
		key:              key,
		IContainerWidget: c,
	}
}

func main() {
	palette := gowid.Palette{
		"focus": gowid.MakePaletteEntry(gowid.ColorBlack, gowid.ColorDarkGray),
		"green": gowid.MakePaletteEntry(gowid.ColorWhite, gowid.ColorGreen),
		"red":   gowid.MakePaletteEntry(gowid.ColorWhite, gowid.ColorRed),
	}

	// Selected tab can be changed by processing Right, Left, Ctrl-F or Ctrl-B.
	tabs := columns.New([]gowid.IContainerWidget{
		NewTab("one", gowid.MakePaletteRef("focus")),
		&gowid.ContainerWidget{IWidget: text.New("|"), D: gowid.RenderWithUnits{U: 1}}, // This is divider, not selectable.
		NewTab("two", gowid.MakePaletteRef("focus")),
		&gowid.ContainerWidget{IWidget: text.New("|"), D: gowid.RenderWithUnits{U: 1}},
		NewTab("three", gowid.MakePaletteRef("focus")),
	})

	tabConts := make(map[string]gowid.IWidget)

	tabConts["one"] = styled.New(text.New("Hello World"), gowid.MakePaletteRef("green"))

	tabConts["two"] = list.New(list.NewSimpleListWalker([]gowid.IWidget{
		text.New("one"),
		text.New("two"),
	}))

	tabConts["three"] = styled.New(text.New("Tab 3"), gowid.MakePaletteRef("red"))

	// Tab contents holder.
	// SubWidget of tabCont is changed when selected tab is changed.
	tabCont := holder.New(tabConts["one"])

	tabs.OnFocusChanged(gowid.WidgetCallback{Name: "focus", WidgetChangedFunction: func(app gowid.IApp, widget gowid.IWidget) {
		if w, ok := widget.(*columns.Widget); ok {
			if tab, ok := w.SubWidgets()[w.Focus()].(*Tab); ok {
				tabCont.SetSubWidget(tabConts[tab.key], app)
			}
		}
	}})

	pl := pile.New([]gowid.IContainerWidget{
		// Show tabs in first row.
		&gowid.ContainerWidget{IWidget: tabs, D: gowid.RenderFlow{}},

		// Show tab contents in remaining space.
		&gowid.ContainerWidget{IWidget: tabCont, D: gowid.RenderWithWeight{W: 1}},
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
