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
	vflipMask = 1 << vflipShift
	paletteLineShift = 12
	paletteLineMask = 0xF << paletteLineShift
)

func RenderTile(tileoffset uint32, mapping uint16, palette color.Palette) image.Image {
	vram := VRAM[tileoffset:]

	tileno := (mapping & tilenoMask) >> tilenoShift
	hflip := ((mapping & hflipMask) >> hflipShift) != 0
	vflip := ((mapping & vflipMask) >> vflipShift) != 0
	paletteLine := (mapping & paletteLineMask) >> paletteLineShift

	// TODO paletted?
	tile := image.NewNRGBA(image.Rect(0, 0, 8, 8))

	vp := uint32(tileno) * tileSize
	paletteLine *= 16		// 16 colors per line

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x += 2 {
			b := vram[vp]
			vp++
			left := b & 0xF
			right := (b & 0xF0) >> 4
			// TODO access data directly to save time?
			if left != 0 {			// transparent
				tile.Set(x, y, palette[paletteLine + uint16(left)])
			}
			if right != 0 {
				tile.Set(x + 1, y, palette[paletteLine + uint16(right)])
			}
		}
	}

	// TODO there should be some optimization to make all this unnecessary...
	if hflip {
		for x := 0; x < 4; x++ {
			for y := 0; y < 8; y++ {
				c1 := tile.At(x, y)
				c2 := tile.At(7 - x, y)
				tile.Set(x, y, c2)
				tile.Set(7 - x, y, c1)
			}
		}
	}
	if vflip {
		for y := 0; y < 4; y++ {
			for x := 0; x < 8; x++ {
				c1 := tile.At(x, y)
				c2 := tile.At(x, 7 - y)
				tile.Set(x, y, c2)
				tile.Set(x, 7 - y, c1)
			}
		}
	}

	return tile
}
