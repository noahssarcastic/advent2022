package main

import (
	"bufio"
	"fmt"
	"os"
)

type Canvas struct {
	pixels [][]byte
	w, h   int
}

func NewCanvas(w, h int) *Canvas {
	canv := Canvas{w: w, h: h}
	canv.pixels = make([][]byte, h)
	for j := range canv.pixels {
		canv.pixels[j] = make([]byte, w)
		for i := range canv.pixels[j] {
			canv.pixels[j][i] = '.'
		}
	}
	return &canv
}

func CopyCanvas(canv *Canvas) *Canvas {
	newCanv := Canvas{w: canv.w, h: canv.h}
	newCanv.pixels = make([][]byte, canv.h)
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
	if x >= canv.w || y >= canv.h {
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
