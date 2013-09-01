// 1 september 2013
package main

import (
	"io"
	"encoding/binary"
	"image/color"
)

// palette format is simple
// first longword is size of palette data in bytes
// this is immediately followed by the palette itself

type GBAColor uint16

const (
	redShift = 0
	redMask = (0x1F << redShift)
	greenShift = 5
	greenMask = (0x1F << greenShift)
	blueShift = 10
	blueMask = (0x1F << blueShift)
	componentScale = 11
)

func (c GBAColor) RGBA() (r, g, b, a uint32) {
	red := (c & redMask) >> redShift
	green := (c & greenMask) >> greenShift
	blue := (c & blueMask) >> blueShift
	return uint32(red) << componentScale,
		uint32(green) << componentScale,
		uint32(blue) << componentScale,
		0xFFFF
}

func ReadGBAColor(r io.Reader) (c GBAColor, err error) {
	err = binary.Read(r, binary.LittleEndian, &c)
	return c, err			// TODO wrap potential error with more text?
}

func ReadPalette(r io.Reader) (p color.Palette, err error) {
	var n uint32

	err = binary.Read(r, binary.LittleEndian, &n)
	if err != nil {
		return nil, err		// TODO wrap with more text?
	}
	n >>= 1				// bytes -> words for color count
	p = make(color.Palette, n)
	for i := uint32(0); i < n; i++ {
		c, err := ReadGBAColor(r)
		if err != nil {
			return nil, err		// TODO wrap with more text?
		}
		p[i] = c
	}
	return p, nil
}
