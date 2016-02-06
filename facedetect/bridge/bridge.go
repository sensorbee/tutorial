package bridge

/*
#cgo linux pkg-config: opencv
#cgo darwin pkg-config: opencv
#include <stdlib.h>
#include "bridge.h"
*/
import "C"
import (
	"pfi/sensorbee/opencv/bridge"
	"reflect"
	"unsafe"
)

// CascadeClassifier is a bind of `cv::CascadeClassifier`
type CascadeClassifier struct {
	p C.CascadeClassifier
}

// NewCascadeClassifier returns a new CascadeClassifier.
func NewCascadeClassifier() CascadeClassifier {
	return CascadeClassifier{p: C.CascadeClassifier_New()}
}

// Delete CascadeClassifier's pointer.
func (c *CascadeClassifier) Delete() {
	C.CascadeClassifier_Delete(c.p)
	c.p = nil
}

// Load cascade configuration file to classifier.
func (c *CascadeClassifier) Load(name string) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.CascadeClassifier_Load(c.p, cName) != 0
}

// Rect represents rectangle. X and Y is a start point of Width and Height.
type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}

// DetectMultiScale detects something which is decided by loaded file. Returns
// multi results addressed with rectangle.
func (c *CascadeClassifier) DetectMultiScale(img bridge.MatVec3b) []Rect {
	ret := C.CascadeClassifier_DetectMultiScale(c.p, C.MatVec3b(img.GetCPointer()))
	defer C.Rects_Delete(ret)

	cArray := ret.rects
	length := int(ret.length)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cArray)),
		Len:  length,
		Cap:  length,
	}
	goSlice := *(*[]C.Rect)(unsafe.Pointer(&hdr))

	rects := make([]Rect, length)
	for i, r := range goSlice {
		rects[i] = Rect{
			X:      int(r.x),
			Y:      int(r.y),
			Width:  int(r.width),
			Height: int(r.height),
		}
	}
	return rects
}

// MatVec4b is a bind of `cv::Mat_<cv::Vec4b>`
type MatVec4b struct {
	p C.MatVec4b
}

// Delete object.
func (m *MatVec4b) Delete() {
	C.MatVec4b_Delete(m.p)
	m.p = nil
}

// DrawRectsToImage draws rectangle information to target image.
func DrawRectsToImage(img bridge.MatVec3b, rects []Rect) {
	cRectArray := make([]C.struct_Rect, len(rects))
	for i, r := range rects {
		cRect := C.struct_Rect{
			x:      C.int(r.X),
			y:      C.int(r.Y),
			width:  C.int(r.Width),
			height: C.int(r.Height),
		}
		cRectArray[i] = cRect
	}
	cRects := C.struct_Rects{
		rects:  (*C.Rect)(&cRectArray[0]),
		length: C.int(len(rects)),
	}
	C.DrawRectsToImage(C.MatVec3b(img.GetCPointer()), cRects)
}

// LoadAlphaImage loads RGBA type image.
func LoadAlphaImage(name string) MatVec4b {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return MatVec4b{p: C.LoadAlphaImg(cName)}
}

// MountAlphaImage draws img on back leading to rects. img is required RGBA,
// TODO should be check file type.
func MountAlphaImage(img MatVec4b, back bridge.MatVec3b, rects []Rect) {
	cRectArray := make([]C.struct_Rect, len(rects))
	for i, r := range rects {
		cRect := C.struct_Rect{
			x:      C.int(r.X),
			y:      C.int(r.Y),
			width:  C.int(r.Width),
			height: C.int(r.Height),
		}
		cRectArray[i] = cRect
	}
	cRects := C.struct_Rects{
		rects:  (*C.Rect)(&cRectArray[0]),
		length: C.int(len(rects)),
	}
	C.MountAlphaImage(img.p, C.MatVec3b(back.GetCPointer()), cRects)
}
