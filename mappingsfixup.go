// 2 september 2013
package main

// TODO make these command-line flags
var (
	init_camera_x int16 = 0
	init_camera_y int16 = 0
)

func sub_8010FE8() {
	if r0 >= 0 {
		goto loc_8010FEE
	}
	goto loc_8011026

loc_8010FEE:
	r6 = m.Header[0]		// width
	if r0 < r6 {
		goto loc_8010FF6
	}
	goto loc_8011026

loc_8010FF6:
	if r1 >= 0 {
		goto loc_8010FFC
	}
	goto loc_8011026

loc_8010FFC:
	r6 = m.Header[1]		// height
	if r1 < r6 {
		goto loc_8011004
	}
	goto loc_8011026

loc_8011004:
	r3 = m.Data
	r4 = 0xC
	r7 = m.Header[0] << 1

loc_801100C:
	r2 = r3.ulong[r4]
	r2 = r3 + uint32(uint16(r2 & 0xFFFF))
	r6 = uint16(r0) << 1
	r2 += uint32(r6)
	r6 = r1
	r6 *= r7		// TODO muls
	r2 = uint32(r2.uint16[r6])
	push(r2)
	r4 -= 4
	if r4 > 0 {
		goto loc_801100C
	}
	pop(r0, r1, r2)		// TODO what order?
	return

loc_8011026:
	r0 = 0
	r1 = 0
	goto loc_8011004
}

// sub_80111F8
func fixup_a(m *rawMappings) []byte {
	camera_x := init_camera_x >> 3	// yes arithmetic shift right
	camera_y := init_camera_y >> 3
	for y := 0; y < 0x20; y++ {
		push(camera_x)
		for x := 0; x < 0x20; x++ {
	r0 = camera_x - 0xF
	r1 = camera_y - 0xA
	r6 := uint16(m.Header[0]) >> 1		// width
	r7 := uint16(m.Header[1]) >> 1		// height
	r0 += int16(r6)		// TODO cast in signedness?
	r1 += int16(r7)
	r8, r9 := camera_x, camera_y
	sub_8010FE8()
	camera_x, camera_y = r8, r9
	r6 := 0x1F
	camera_x &= r6	// TODO what do I do about signedness here
	camera_y &= r6	// in asm it doesn't matter (unsigned) but I need to change everything else
	sub_801102C()
	camera_x, camera_y = r8, r9
	camera_x++
		}
		pop(camera_x)
		camera_y++
	}
}
