// 1 september 2013
package main

import (
	"io"
	"encoding/binary"
)

// the tile data structure is simple: it's three consecutive sets of three longwords
// first longword is the number of longwords to copy
// second longword is offset FROM THE TOP OF THE WHOLE STRUCTURE (so from the address of that first longword of the first set)
// third longword is the offset in VRAM to load to

func ReadTiles(r io.ReadSeeker) (offsets [3]uint32, err error) {
	var tiles [3]struct {
		Len		uint32
		Srcoff	uint32
		Destoff	uint32
	}
	var base uint32

	// get the start, since we need to take its offset
	xbase, err := r.Seek(0, 1)
	if err != nil {
		return offsets, err		// TODO wrap potential error with more text?
	}
	base = uint32(xbase)

	// read everything
	err = binary.Read(r, binary.LittleEndian, &tiles)
	if err != nil {
		return offsets, err		// TODO wrap potential error with more text?
	}

	for i := 0; i < 3; i++ {
		// go to the LZ77-compressed tile data
		_, err = r.Seek(int64(base + tiles[i].Srcoff), 0)
		if err != nil {
			return offsets, err	// TODO wrap potential error with more text?
		}

		u, err := LZ77Decomp(r)
		if err != nil {
			return offsets,err	// TODO wrap potential error with more text?
		}

		tiles[i].Len <<= 2		// longwords -> bytes
		for j := uint32(0); j < tiles[i].Len; j++ {
			VRAM[tiles[i].Destoff + j] = u[j]
		}
		offsets[i] = tiles[i].Destoff
	}

	return offsets, nil
}
