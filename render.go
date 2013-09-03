// 1 september 2013
package main

import (
	"image"
	"image/color"
	"image/draw"
)

func Render(offsets [3]uint32, m *Mappings, palette color.Palette) image.Image {
	width := int(uint32(m.Header[0]) * 8)
	height := int(uint32(m.Header[1]) * 8)

	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for i := 0; i < 3; i++ {
		dataoff := uint32(m.Header[(i + 1) * 4]) |
			(uint32(m.Header[((i + 1) * 4) + 1]) << 8) |
			(uint32(m.Header[((i + 1) * 4) + 2]) << 16) |
			(uint32(m.Header[((i + 1) * 4) + 3]) << 24)
		dataoff -= 0x10	// offset is stored relative to top of mappings; make it relative to top of m.Data instead
		mp := uint32(0)
		for y := 0; y < height; y += 8 {
			for x := 0; x < width; x += 8 {
				var md uint16

				md = uint16(m.Data[dataoff + mp])
				md |= uint16(m.Data[dataoff + mp + 1]) << 8
				mp += 2
				tile := RenderTile(offsets[i], md, palette)
				draw.Draw(img,
					image.Rect(x, y, x + 8, y + 8),
					tile, image.ZP, draw.Over)
			}
		}
	}

	return img
}
