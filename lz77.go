// 1 september 2013
package main

import (
	"fmt"
	"io"
	"encoding/binary"
)

const (
	// header
	lzFormatShift = 4
	lzFormatMask = 0xF << lzFormatShift
	lzSizeShift = 8
	lzSizeMask = 0xFFFFFF << lzSizeShift

	// copy block spec
	lzDispLeftShift = 0
	lzDispLeftMask = 0xF << lzDispLeftShift
	lznCopyShift = 4
	lznCopyMask = 0xF << lznCopyShift
	lzDispRightShift = 8
	lzDispRightMask = 0xFF << lzDispRightShift

	// combines the two halves of the disp above
	lzDispFinalShift = 8
)

func LZ77Decomp(r io.Reader) (data []byte, err error) {
	readbyte := func() (b byte, err error) {
		err = binary.Read(r, binary.LittleEndian, &b)
		return b, err
	}
	readword := func() (w uint16, err error) {
		err = binary.Read(r, binary.LittleEndian, &w)
		return w, err
	}
	readlong := func() (l uint32, err error) {
		err = binary.Read(r, binary.LittleEndian, &l)
		return l, err
	}

	header, err := readlong()
	if err != nil {
		return nil, err		// TODO wrap with more text?
	}

	format := (header & lzFormatMask) >> lzFormatShift
	size := (header & lzSizeMask) >> lzSizeShift

	if format != 1 {
		return nil, fmt.Errorf("error: invalid LZ77 data compression format (expected 1, got 0x%X)", format)
	}

	data = make([]byte, 0, size)

	for uint32(len(data)) < size {
		var b byte

		b, err = readbyte()
		if err != nil {
			return nil, err		// TODO wrap with more text?
		}
		for bits := 0; bits < 8; bits++ {
			if (b & 0x80) == 0 {	// copy next byte
				c, err := readbyte()
				if err != nil {
					return nil, err		// TODO wrap with more text?
				}
				data = append(data, c)
			} else {			// opy n+3 bytes from data
				copyspec, err := readword()
				if err != nil {
					return nil, err		// TODO wrap with more text?
				}
				dispRight := (copyspec & lzDispRightMask) >> lzDispRightShift
				dispLeft := (copyspec & lzDispLeftMask) >> lzDispLeftShift
				disp := (dispLeft << lzDispFinalShift) | dispRight
				nCopy := (copyspec & lznCopyMask) >> lznCopyShift
				for i := uint16(0); i < nCopy + 3; i++ {
					// TODO convert len(data) and disp to uint32?
					data = append(data, data[len(data) - int(disp) - 1])
				}
			}
			b <<= 1
		}
	}

//	if uint32(len(data)) != size {
//		return nil, fmt.Errorf("error: somehow we decompressed a larger blob than expected (expected size 0x%X, got 0x%X)", size, len(data))
//	}

	return data, nil
}
