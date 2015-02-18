package gl

import (
	"unsafe"
)

//#cgo linux	LDFLAGS: -lGL
//#cgo windows  LDFLAGS: -lopengl32
import "C"

/*
#include <GL/gl.h>
*/
import "C"

type (
	PVoid    unsafe.Pointer
	Enum     uint32
	Boolean  bool
	Bitfield uint32
	Byte     int8
	Short    int16
	Int      int32
	Ubyte    uint8
	Ushort   uint16
	Uint     uint32
	Sizei    int32
	Float    float32
	Clampf   float32
	Double   float64
	Clampd   float64
)

const (
	gl_TRUE  = C.GLboolean(1)
	gl_FALSE = C.GLboolean(0)
)

func (t Enum) v() C.GLenum {
	return C.GLenum(t)
}

func (t Boolean) v() C.GLboolean {
	if t {
		return gl_TRUE
	}
	return gl_FALSE
}

func (t Bitfield) v() C.GLbitfield {
	return C.GLbitfield(t)
}

func (t Byte) v() C.GLbyte {
	return C.GLbyte(t)
}

func (t Short) v() C.GLshort {
	return C.GLshort(t)
}

func (t Int) v() C.GLint {
	return C.GLint(t)
}

func (t Ubyte) v() C.GLubyte {
	return C.GLubyte(t)
}

func (t Ushort) v() C.GLushort {
	return C.GLushort(t)
}

func (t Uint) v() C.GLuint {
	return C.GLuint(t)
}

func (t Sizei) v() C.GLsizei {
	return C.GLsizei(t)
}

func (t Float) v() C.GLfloat {
	return C.GLfloat(t)
}

func (t Clampf) v() C.GLclampf {
	return C.GLclampf(t)
}

func (t Double) v() C.GLdouble {
	return C.GLdouble(t)
}

func (t Clampd) v() C.GLclampd {
	return C.GLclampd(t)
}

func Enable(c Enum) {
	C.glEnable(c.v())
}

func DepthFunc(f Enum) {
	C.glDepthFunc(f.v())
}

func ClearDepth(depth Clampd) {
	C.glClearDepth(depth.v())
}

func ClearColor(red, green, blue, alpha Clampf) {
	C.glClearColor(red.v(), green.v(), blue.v(), alpha.v())
}

func MatrixMode(mode Enum) {
	C.glMatrixMode(mode.v())
}

func LoadIdentity() {
	C.glLoadIdentity()
}

func Frustum(left, right, bottom, top, near, far Double) {
	C.glFrustum(left.v(), right.v(), bottom.v(), top.v(), near.v(), far.v())
}

func Ortho(left, right, bottom, top, near, far Double) {
	C.glOrtho(left.v(), right.v(), bottom.v(), top.v(), near.v(), far.v())
}

func Viewport(x, y Int, width, height Sizei) {
	C.glViewport(x.v(), y.v(), width.v(), height.v())
}

func Clear(mask Bitfield) {
	C.glClear(mask.v())
}

func Color3f(red, green, blue Float) {
	C.glColor3f(red.v(), green.v(), blue.v())
}

func Begin(mode Enum) {
	C.glBegin(mode.v())
}

func End() {
	C.glEnd()
}

func Vertex3f(x, y, z Float) {
	C.glVertex3f(x.v(), y.v(), z.v())
}

func EnableClientState(array Enum) {
	C.glEnableClientState(array.v())
}

func DisableClientState(array Enum) {
	C.glDisableClientState(array.v())
}

func VertexPointer(size Int, type_ Enum, stride Sizei, pointer PVoid) {
	C.glVertexPointer(size.v(), type_.v(), stride.v(), unsafe.Pointer(pointer))
}

func DrawArrays(mode Enum, first Int, count Sizei) {
	C.glDrawArrays(mode.v(), first.v(), count.v())
}

func PixelStorei(pname Enum, param Int) {
	C.glPixelStorei(pname.v(), param.v())
}

func GenLists(range_ Sizei) Uint {
	return Uint(C.glGenLists(range_.v()))
}

func NewList(list Uint, mode Enum) {
	C.glNewList(list.v(), mode.v())
}

func EndList() {
	C.glEndList()
}

func Bitmap(w Sizei, h Sizei, xo Float, yo Float, xm Float, ym Float, ptr *Ubyte) {
	C.glBitmap(w.v(), h.v(), xo.v(), yo.v(), xm.v(), ym.v(), (*C.GLubyte)(ptr))
}

func PushAttrib(mask Bitfield) {
	C.glPushAttrib(mask.v())
}

func PopAttrib() {
	C.glPopAttrib()
}

func ListBase(base Uint) {
	C.glListBase(base.v())
}

func CallLists(n Sizei, type_ Enum, lists PVoid) {
	C.glCallLists(n.v(), type_.v(), unsafe.Pointer(lists))
}

func RasterPos2i(x Int, y Int) {
	C.glRasterPos2i(x.v(), y.v())
}
