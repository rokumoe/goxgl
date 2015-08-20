package context

import "C"

const (
	KEYCODE_LEFT = iota
	KEYCODE_RIGHT
	KEYCODE_UP
	KEYCODE_DOWN
	KEYCODE_SHIFT
	KEYCODE_UNKNOWN
)

var (
	OnInit    (func())
	OnExit    (func())
	OnSize    (func(w int32, h int32))
	OnDraw    (func())
	OnKeyDown (func(kc int))
	OnKeyUp   (func(kc int))
)

//export onInit
func onInit() {
	f := OnInit
	if f != nil {
		f()
	}
}

//export onExit
func onExit() {
	f := OnExit
	if f != nil {
		f()
	}
}

//export onSize
func onSize(w, h int32) {
	f := OnSize
	if f != nil {
		f(w, h)
	}
}

//export onDraw
func onDraw() {
	f := OnDraw
	if f != nil {
		f()
	}
}

//export onKeyDown
func onKeyDown(kc int32) {
	f := OnKeyDown
	if f != nil {
		f(int(kc))
	}
}

//export onKeyUp
func onKeyUp(kc int32) {
	f := OnKeyUp
	if f != nil {
		f(int(kc))
	}
}
