package debug

import (
	"bufio"
	"fmt"
	"os"
)

type Canvas struct {
	pixels [][]byte
}

func (canv *Canvas) Width() int {
	return len(canv.pixels[0])
}

func (canv *Canvas) Height() int {
	return len(canv.pixels)
}

func NewCanvas(w, h int) *Canvas {
	canv := Canvas{}
	canv.pixels = make([][]byte, h)
	for j := range canv.pixels {
		canv.pixels[j] = make([]byte, w)
		for i := range canv.pixels[j] {
			canv.pixels[j][i] = '.'
		}
	}
	return &canv
}

func (canv *Canvas) Copy() *Canvas {
	newCanv := Canvas{make([][]byte, canv.Height())}
	for i := range canv.pixels {
		newCanv.pixels[i] = make([]byte, canv.Width())
		copy(newCanv.pixels[i], canv.pixels[i])
	}
	return &newCanv
}

func (canv *Canvas) Resize(newWidth, newHeight int) {
	oldHeight := canv.Height()
	oldWidth := canv.Width()
	canv.pixels = append(canv.pixels, make([][]byte, newHeight-oldHeight)...)
	for j := range canv.pixels {
		canv.pixels[j] = append(canv.pixels[j], make([]byte, newWidth-oldWidth)...)
		for i := range canv.pixels[j][oldWidth:] {
			canv.pixels[j][oldWidth+i] = '.'
		}
	}
}

func (canv *Canvas) Draw(x, y int, p byte) {
	if x >= canv.Width() || y >= canv.Height() {
		// canv.Resize()
		// canv.Draw(x, y, p)
		panic("must resize!")
	} else {
		canv.pixels[y][x] = p
	}
}

type BBox struct {
	x0, y0 int
	x1, y1 int
}

func Bounds(x0, y0, x1, y1 int) *BBox {
	return &BBox{x0, y0, x1, y1}
}

func (bb *BBox) XMin() int {
	return bb.x0
}
func (bb *BBox) XMax() int {
	return bb.x1
}
func (bb *BBox) YMin() int {
	return bb.y0
}
func (bb *BBox) YMax() int {
	return bb.y1
}

func (bb *BBox) Expand(x, y int) {
	if x < bb.x0 {
		bb.x0 = x
	}
	if x > bb.x1 {
		bb.x1 = x
	}
	if y < bb.y0 {
		bb.y0 = y
	}
	if y > bb.y1 {
		bb.y1 = y
	}
}

func (canv *Canvas) Print(bb BBox) {
	for y := bb.y0; y <= bb.y1; y++ {
		row := canv.pixels[y][bb.x0 : bb.x1+1]
		fmt.Println(string(row))
	}
}

func Pause() {
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
