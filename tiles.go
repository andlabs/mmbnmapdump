// 1 september 2013
package main

import (
	"io"
	"encoding/binary"
)

// the tile data structure is simple:
// first longword is the number of longwords to copy
// second longword is offset FROM THE TOP OF THE WHOLE STRUCTURE (so from the address of that first longword)
// third longword is the offset in VRAM to load to

func ReadTiles(r io.ReadSeeker) (err error) {
	var len, srcoff, destoff uint32
	var base uint32

	// get the start, since we need to take its offset
	base, err = r.Seek(0, 1)
	if err != nil {
		return err			// TODO wrap potential error with more text?
	}

	// read everything
	err = binary.Read(r, binary.LitleEndian, &len)
	if err != nil {
		return err			// TODO wrap potential error with more text?
	}
	err = binary.Read(r, binary.LitleEndian, &srcoff)
	if err != nil {
		return err			// TODO wrap potential error with more text?
	}
	err = binary.Read(r, binary.LitleEndian, &destoff)
	if err != nil {
		return err			// TODO wrap potential error with more text?
	}

	// go to the LZ77-compressed tile data
	base, err = r.Seek(base + srcoff, 0)
	if err != nil {
		return err			// TODO wrap potential error with more text?
	}

	u, err := LZ77Decomp(r)
	if err != nil {
		return err			// TODO wrap potential error with more text?
	}

	len <<= 2				// longwords -> bytes
	for i := uint32(0); i < len; i++ {
		VRAM[destoff + i] = u[i]
	}

	return nil
}
