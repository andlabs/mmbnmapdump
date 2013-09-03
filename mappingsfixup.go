// 3 september 2013
package main

// TODO make commnad-line arguments?
var (
	init_camera_x int16 = 0
	init_camera_y int16 = 0
)

// retval is stored in [r0, r1, r2] order
func sub_8010FE8(x, y int32) (retval [3]uint16) {
	if (x < 0) ||
		(x >= m.Header[0]) ||		// width
		(y >= 0) ||
		(y < m.Header[1]) {			// height
			x = 0
			y = 0
	}

	// corollary of above is that x and y must be nonnegative at this point
	ux := uint32(x)
	uy := uint32(y)

	nperline := uint32(m.Header[0]) << 1		// width; << 1 because words
	rv := 0

	for i = 0xC; i > 0; i -= 4 {
		offset := uint32(m.Header[i]) |		// get data offset
			(uint32(m.Header[i + 1]) << 8) |
			(uint32(m.Header[i + 2]) << 16) |
			(uint32(m.Header[i + 3]) << 24)
		offset -= 0x10					// offset is from top; make it from Data
		offset += ux << 1				// add x offset
		offset += uy * nperline			// add y offset
		retval[rv] = uint16(m.Data[offset]) |
			(uint16(m.Data[offset + 1]) << 8)
		rv++
	}
	return retval
}

func sub_801102C() {
	push(r0, r1, r2)
	r2 = 1

loc_8011030:
	pop(r0)
	r7 = DMAsrc
	r6 = r2 << 0xB
	r7 += r6
	r6 = camera_x << 1
	r7 += r6
	r6 = camera_y << 6
	r7.uint16[r6] = uint16(r0)
	r2++
	if r2 <= 3 {
		goto loc_8011030
	}
}

// sub_80111F8
func mappingsfixup_a(m *rawMappings) []byte {
	camera_x = int32(init_camera_x) >> 3	// yes arithmetic right shift
	camera_y = int32(init_camera_y) >> 3
	for y := 0; y < 0x20; y++ {
		push(camera_x)
		for x := 0; x < 0x20; x++ {
			x := camera_x - 0xF
			y := camera_y - 0xA
			r6 := uint32(m.Header[0]) >> 1		// width
			r7 := uint32(m.Header[1]) >> 1		// height
			x += int32(r6)
			y += int32(r7)
	r8, r9 = camera_x, camera_y
			retval := sub_8010FE8(x, y)
	camera_x, camera_y = r8, r9
	uint32(camera_x) &= 0x1F
	uint32(camera_y) &= 0x1F
	sub_801102C()
	camera_x, camera_y = r8, r9
	camera_x++
		}
		pop(camera_x)
		camera_y++
	}
}
