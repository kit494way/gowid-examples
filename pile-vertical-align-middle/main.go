package main

import (
	"fmt"
	"os"

	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/divider"
	"github.com/gcla/gowid/widgets/pile"
	"github.com/gcla/gowid/widgets/text"
	"github.com/gcla/gowid/widgets/vpadding"
)

func main() {
	flow := gowid.RenderFlow{}

	pl := pile.New([]gowid.IContainerWidget{
		&gowid.ContainerWidget{IWidget: divider.NewAscii(), D: flow},
		&gowid.ContainerWidget{IWidget: text.New("vertical align middle"), D: flow},
		&gowid.ContainerWidget{IWidget: divider.NewAscii(), D: flow},
	})

	view := vpadding.New(pl, gowid.VAlignMiddle{}, flow)

	app, err := gowid.NewApp(gowid.AppArgs{View: view})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	app.SimpleMainLoop()
}
