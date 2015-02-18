package main

import "goxgl/context"

func main() {
	context.OnInit = onInit
	context.OnSize = onSize
	context.OnDraw = onDraw
	context.OnKeyDown = onKeyDown
	context.OnKeyUp = onKeyUp
	if context.InitDisplay("goxgl", 0, 0, 600, 600) == 0 {
		start()
		context.MainLoop()
	}
}
