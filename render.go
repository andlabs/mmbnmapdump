// 1 september 2013
package main

import (
	"image"
	"image/color"
	"image/draw"
)

func Render(offsets [3]uint32, m *Mappings, palette color.Palette) image.Image {
	width := int(uint32(m.Width) * 8)
	height := int(uint32(m.Height) * 8)

	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for i := 0; i < 3; i++ {
		dataoff := m.BGOff[i]
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
