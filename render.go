// 1 september 2013
package main

import (
	"image"
	"image/color"
	"image/draw"
)

func Render(m *Mappings, palette color.Palette) (img image.Image) {
	width := uint32(m.Header[0]) * 8
	height := uint32(m.Header[1]) * 8

	img = image.NewNRGBA(image.Rect(0, 0, width, height))

	mp := 0

	for y := 0; y < height; y += 8 {
		for x := 0; x < width; x += 8 {
			var md uint16

			md = uint16(m.Data[mp])
			md |= uint16(m.Data[mp + 1]) << 8
			mp += 2
			tile := RenderTile(md, palette)
			draw.Draw(img,
				image.Rect(x, y, x + 8, y + 8),
				tile, image.ZP, draw.Over)
		}
	}

	return img
}
