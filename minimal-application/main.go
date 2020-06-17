package main

import (
	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/text"
)

func main() {
	txt := text.New("Hello World")
	app, _ := gowid.NewApp(gowid.AppArgs{View: txt})
	app.SimpleMainLoop()
}
