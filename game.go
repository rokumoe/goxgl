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
	if p.x+p.w < b.x || p.y+p.h < b.y || b.x+b.w < p.x || b.y+b.h < p.y {
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
		l := gl.Float(p.x)
		t := gl.Float(p.y)
		r := gl.Float(p.x + p.w)
		b := gl.Float(p.y + p.h)
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
	meVx       = 40
	meVy       = 40
)

var (
	gameOver    bool
	total       float32
	base        int
	genbox      float32
	highest     float32
	highestText string = "HIGHEST 0"
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
	drawText("SCORE "+strconv.FormatFloat(float64(total), 'f', 1, 32), 8, gameHeight-13-8)
	drawText(highestText, 8, gameHeight-26-8)
	if !gameOver {
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
	default:
		if gameOver {
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
	}
}

func fixBox(b *Box) {
	if b.x < -b.w || b.x > gameWidth || b.y < -b.h || b.y > gameHeight {
		b.bad = true
	}
}

func fixMe(b *Box) {
	if b.x < 0 {
		b.x = 0
		b.dirty = true
	} else if b.x+b.w > gameWidth {
		b.x = gameWidth - b.w
		b.dirty = true
	}
	if b.y < 0 {
		b.y = 0
		b.dirty = true
	} else if b.y+b.h > gameHeight {
		b.y = gameHeight - b.w
		b.dirty = true
	}
}

func putBox(x, y, w, h, vx, vy float32) *Box {
	b := &Box{dirty: true, x: x, y: y, w: w, h: h, vx: vx, vy: vy}
	b.setColor(rand.Float32(), rand.Float32(), rand.Float32())
	return b
}

func start() {
	if len(boxes) > 0 {
		boxes = boxes[:0]
	}
	me.setPosition(gameWidth/2-5, gameHeight/2-5)
	total = 0.0
	genbox = 0.0
	base = 0
	gameOver = false
	go game()
}

func game() {
	fps := time.NewTicker(time.Second / 60.0)
	last := time.Now()
	for current := range fps.C {
		delta := float32(current.Sub(last)) / float32(time.Second)
		total += delta
		genbox += delta
		if genbox > 2.4 {
			b := int(math.Log2(float64(total)))
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
					boxes = append(boxes, putBox(x, y, 4, 4, vx, vy))
				}
			}
			genbox = 0.0
		}
		for _, b := range boxes {
			b.update(delta)
			fixBox(b)
		}
		me.update(delta)
		fixMe(me)
		for _, b := range boxes {
			if b.intersect(me) {
				gameOver = true
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
		if gameOver {
			if total > highest {
				highest = total
				highestText = "HIGHEST " + strconv.FormatFloat(float64(total), 'f', 1, 32)
			}
			break
		}
		last = current
	}
	fps.Stop()
}
