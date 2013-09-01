// 1 september 2013
package main

import (
	"fmt"
	"os"
	"encoding/binary"
	"image/png"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: %s ROM tripletloc-hex\n", os.Args[0])
		os.Exit(1)
	}

	var tilespos, palettepos, mappingpos uint32

	tripletpos, err := strconv.ParseUInt(os.Args[2], 16, 32)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid headless hex number %q for triplet pos\n", os.Args[2])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.Seek(tripletpos, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error seeking to triplet pos 0x%X: %v\n", tripletpos, err)
		os.Exit(1)
	}
	err = binary.Read(f, binary.LittleEndian, &tilespos)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading tiles pos: %v\n", err)
		os.Exit(1)
	}
	err = binary.Read(f, binary.LittleEndian, &palettepos)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading palette pos: %v\n", err)
		os.Exit(1)
	}
	err = binary.Read(f, binary.LittleEndian, &mappingspos)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading mappings pos: %v\n", err)
		os.Exit(1)
	}

	_, err = f.Seek(tilespos, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error seeking to tiles pos 0x%X: %v\n", tilespos, err)
		os.Exit(1)
	}
	err = ReadTiles(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading tiles: %v\n", err)
		os.Exit(1)
	}

	_, err = f.Seek(palettepos, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error seeking to palette pos 0x%X: %v\n", palettepos, err)
		os.Exit(1)
	}
	palette, err = ReadPalette(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading palette: %v\n", err)
		os.Exit(1)
	}

	_, err = f.Seek(mappingspos, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error seeking to mappings pos 0x%X: %v\n", mappingspos, err)
		os.Exit(1)
	}
	mappings, err = ReadMappings(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading mappings: %v\n", err)
		os.Exit(1)
	}

	image := Render(mappings, palette)
	err = png.Encode(os.Stdout, image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error saving png: %v\n", err)
		os.Exit(1)
	}

	// otherwise phew, all clear!
}
