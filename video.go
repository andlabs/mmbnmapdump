// 1 september 2013
package main

import (
	"image"
	"image/color"
)

var (
	VRAM	[0x18000]byte
)

const (
	tileSize = 32		// this is 4bpp mode

	// mapping bits
	tilenoShift = 0
	tilenoMask = 0x3F << tilenoShift
	hflipShift = 10
	hflipMask = 1 << hflipShift
	vflipShift = 11
	vflipMask = 1 << vflipMask
	paletteLineShift = 12
	paletteLineMask = 0xF << paletteMask
)

func RenderTile(mapping uint16, palette color.Palette) (tile image.Image) {
	tileno := (mapping & tilenoMask) >> tilenoShift
	hflip := ((mapping & hflipMask) >> hflipShift) != 0
	vflip := ((mapping & vflipMask) >> vflipShift) != 0
	paletteLine := (mapping & paletteLineMask) >> paletteLineShift

	// TODO paletted?
	tile = image.NewNRGBA(image.Rect(0, 0, 8, 8)

	yoff := 0
	if yflip {
		yoff = 7
	}
	xoff := 0
	if xflip {
		xoff = 7
	}

	vp := uint32(tileno) * tileSize
	paletteLine *= 16		// 16 colors per line

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x += 2 {
			b := VRAM[vp]
			left := b & 0xF
			right := (b & 0xF0) >> 4
			// TODO access data directly to save time?
			tile.Set(xoff - x, yoff - y, palette[paletteLine + uint16(left)])
			tile.Set(xoff - (x + 1), yoff - y, palette[paletteLine + uint16(right)])
		}
	}

	return tile
}
