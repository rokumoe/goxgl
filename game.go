package main

import (
	"goxgl/context"
	"goxgl/gl"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Box struct {
	bad    bool
	dirty  bool
	x      float32
	y      float32
	w      float32
	h      float32
	vx     float32
	vy     float32
	color  [3]float32
	vertex [8]gl.Float
}

func (p *Box) setColor(r, g, b float32) {
	p.color[0] = r
	p.color[1] = g
	p.color[2] = b
}

func (p *Box) setPosition(x, y float32) {
	p.x = x
	p.y = y
	p.dirty = true
}

func (p *Box) setSize(w, h float32) {
	p.w = w
	p.h = h
	p.dirty = true
}

func (p *Box) intersect(b *Box) bool {
	hw1 := p.w / 2
	hh1 := p.h / 2
	hw2 := b.w / 2
	hh2 := b.h / 2
	if p.x+hw1 < b.x-hw2 || p.y+hh1 < b.y-hh2 || b.x+hw2 < p.x-hw1 || b.y+hh2 < p.y-hh1 {
		return false
	}
	return true
}

func (p *Box) update(d float32) {
	if p.vx != 0 || p.vy != 0 {
		p.x += d * p.vx
		p.y += d * p.vy
		p.dirty = true
	}
}

func (p *Box) draw() {
	if p.dirty {
		hw := p.w / 2
		hh := p.h / 2
		l := gl.Float(p.x - hw)
		t := gl.Float(p.y - hh)
		r := gl.Float(p.x + hw)
		b := gl.Float(p.y + hh)
		p.vertex[0] = l
		p.vertex[1] = t
		p.vertex[2] = l
		p.vertex[3] = b
		p.vertex[4] = r
		p.vertex[5] = b
		p.vertex[6] = r
		p.vertex[7] = t
		p.dirty = false
	}
	gl.Color3f(gl.Float(p.color[0]), gl.Float(p.color[1]), gl.Float(p.color[2]))
	gl.VertexPointer(2, gl.FLOAT, 0, gl.PVoid(&p.vertex[0]))
	gl.DrawArrays(gl.QUADS, 0, 4)
}

const (
	gameWidth  = 320
	gameHeight = 320
	meVx       = 48
	meVy       = 48
)

var (
	over        bool
	score       float32
	genboxSlot  float32
	highest     float32
	highestText string  = "HIGHEST 0.0"
	meScale     float32 = 1.0
)

var (
	boxes []*Box
	me    *Box
)

func onInit() {
	initText()
	gl.Ortho(0.0, gl.Double(gameWidth), 0.0, gl.Double(gameHeight), -1.0, 1.0)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	me = &Box{}
	me.setColor(1, 1, 1)
	me.setSize(6, 6)
}

func onSize(w, h int32) {
	gl.Viewport(0, 0, gl.Sizei(w), gl.Sizei(h))
}

func onDraw() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	drawText("SCORE "+strconv.FormatFloat(float64(score), 'f', 1, 32), 8, gameHeight-13-8)
	drawText(highestText, 8, gameHeight-26-8)
	if !over {
		gl.EnableClientState(gl.VERTEX_ARRAY)
		for _, b := range boxes {
			b.draw()
		}
		me.draw()
		gl.DisableClientState(gl.VERTEX_ARRAY)
	} else {
		drawText("GAME OVER", gameWidth/2-28, gameHeight/2+7)
	}
	context.SwapBuffer()
}

func onKeyDown(kc int) {
	switch kc {
	case context.KEYCODE_LEFT:
		me.vx = -meVx
	case context.KEYCODE_RIGHT:
		me.vx = meVx
	case context.KEYCODE_UP:
		me.vy = meVy
	case context.KEYCODE_DOWN:
		me.vy = -meVy
	case context.KEYCODE_SHIFT:
		meScale = 0.5
		me.setSize(2, 2)
	default:
		if over {
			start()
			return
		}
	}
}

func onKeyUp(kc int) {
	switch kc {
	case context.KEYCODE_LEFT:
		if me.vx < 0 {
			me.vx = 0
		}
	case context.KEYCODE_RIGHT:
		if me.vx > 0 {
			me.vx = 0
		}
	case context.KEYCODE_UP:
		if me.vy > 0 {
			me.vy = 0
		}
	case context.KEYCODE_DOWN:
		if me.vy < 0 {
			me.vy = 0
		}
	case context.KEYCODE_SHIFT:
		meScale = 1.0
		me.setSize(6, 6)
	}
}

func fixBox(b *Box) {
	hw := b.w / 2
	hh := b.h / 2
	if b.x < -hw || b.x > gameWidth+hw || b.y < -hh || b.y > gameHeight+hh {
		b.bad = true
	}
}

func fixMe(b *Box) {
	hw := b.w / 2
	hh := b.h / 2
	if b.x < hw {
		b.x = hw
		b.dirty = true
	} else if b.x+hw > gameWidth {
		b.x = gameWidth - hw
		b.dirty = true
	}
	if b.y < hh {
		b.y = hh
		b.dirty = true
	} else if b.y+hh > gameHeight {
		b.y = gameHeight - hh
		b.dirty = true
	}
}

func newBox(x, y, w, h, vx, vy float32) *Box {
	b := &Box{dirty: true, x: x, y: y, w: w, h: h, vx: vx, vy: vy}
	b.setColor(rand.Float32(), rand.Float32(), rand.Float32())
	return b
}

func start() {
	if len(boxes) > 0 {
		boxes = boxes[:0]
	}
	me.setPosition(gameWidth/2-5, gameHeight/2-5)
	me.setSize(6, 6)
	meScale = 1.0
	score = 0.0
	genboxSlot = 0.0
	over = false
	go game()
}

func game() {
	fps := time.NewTicker(time.Second / 60.0)
	last := time.Now()
	for current := range fps.C {
		delta := float32(current.Sub(last)) / float32(time.Second)
		score += delta
		genboxSlot += delta
		if genboxSlot > 2.4 {
			b := int(math.Log2(float64(score)))
			for k := 0; k < 4; k++ {
				c := b + rand.Intn(4)
				for i := 0; i < c; i++ {
					var x float32
					var y float32
					if k&2 == 0 {
						x = float32((k & 1) * gameWidth)
						y = float32(rand.Intn(gameHeight))
					} else {
						x = float32(rand.Intn(gameWidth))
						y = float32((k & 1) * gameHeight)
					}
					vx := (me.x - x) * (0.4 + float32(rand.Intn(20))/100)
					vy := (me.y - y) * (0.4 + float32(rand.Intn(20))/100)
					boxes = append(boxes, newBox(x, y, 4, 4, vx, vy))
				}
			}
			genboxSlot = 0.0
		}
		for _, b := range boxes {
			b.update(delta)
			fixBox(b)
		}
		me.update(delta * meScale)
		fixMe(me)
		for _, b := range boxes {
			if b.intersect(me) {
				over = true
				break
			}
		}
		removed := 0
		for i := len(boxes) - 1; i >= 0; i-- {
			if boxes[i].bad {
				copy(boxes[i:], boxes[i+1:])
				removed++
			}
		}
		boxes = boxes[:len(boxes)-removed]
		context.RequsetRedraw()
		if over {
			if score > highest {
				highest = score
				highestText = "HIGHEST " + strconv.FormatFloat(float64(score), 'f', 1, 32)
			}
			break
		}
		last = current
	}
	fps.Stop()
}
