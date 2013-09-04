// 1 september 2013
package main

import (
	"io"
)

// the mappings format is simple
// a 0x10-byte header of the form
// 	struct Header {
// 		Width	byte
// 		Height	byte
// 		Unknown	[2]byte
// 		BG0Off	uint32	// little endian
// 		BG1Off	uint32	// little endian
// 		BG2Off	uint32	// little endian
// 	}
// immediately followed by the mappings themselves, LZ77 compressed
// When loaded, the header is first copied over untouched, then the compressed data is decompressed immediately afterward. The three offset fields are offset from the first byte of this block (so from the Width byte in the copied header).

type Mappings struct {
	Width	byte
	Height	byte
	BGOff	[3]uint32
	Data		[]byte
}

func ReadMappings(r io.Reader) (m *Mappings, err error) {
	m = new(Mappings)

	var header [0x10]byte

	_, err = r.Read(header[:])
	if err != nil {
		return nil, err		// TODO wrap with more text?
	}

	m.Width = header[0]
	m.Height = header[1]
	for i := 0; i < 3; i++ {
		j := (i + 1) * 4
		m.BGOff[i] = uint32(header[j]) |
			(uint32(header[j + 1]) << 8) |
			(uint32(header[j + 2]) << 16) |
			(uint32(header[j + 3]) << 24)
		m.BGOff[i] -= 0x10	// value in header is relative to top of header; make it relative to top of Data instead
	}

	m.Data, err = LZ77Decomp(r)
	if err != nil {
		return nil, err		// TODO wrap with more text?
	}

	return m, nil
}
