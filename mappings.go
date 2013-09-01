// 1 september 2013
package main

import (
	"io"
)

// the mappings format is simple
// 0x10 bytes which are copied raw (the first two bytes of whicih appear to be extra special?)
// immediately followed by the mappings themselves, LZ77 compressed

type Mappings struct {
	Header	[0x10]byte
	Data		[]byte
}

func ReadMappings(r io.Reader) (m *Mappings, err error) {
	m = new(Mappings)

	_, err = r.Read(m.Header[:])
	if err != nil {
		return nil, err		// TODO wrap with more text?
	}

	m.Data, err = LZ77Decomp(r)
	if err != nil {
		return nil, err		// TODO wrap with more text?
	}

	return m, nil
}
