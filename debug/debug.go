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

func New(w, h int) *Canvas {
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

func Copy(canv *Canvas) *Canvas {
	newCanv := Canvas{}
	newCanv.pixels = make([][]byte, canv.Height())
	copy(newCanv.pixels, canv.pixels)
	return &newCanv
}

// // TODO: fix
// func (canv *Canvas) Resize() {
// 	canv.pixels = append(canv.pixels, make([]byte, canv.h))
// 	for i := 0; i < canv.h; i++ {
// 		canv.pixels = append(canv.pixels, make([]byte, canv.w*2))
// 	}
// 	canv.h *= 2
// 	canv.w *= 2
// }

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
