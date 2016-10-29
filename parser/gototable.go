// generated by gocc; DO NOT EDIT.

package parser

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
)

const numNTSymbols = 17

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
)

var gotoTab = gotoTable{}

func init() {
	tab := [][]int{}
	data := []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xe2, 0xfd, 0xdf, 0xcc, 0xc4, 0xc8,
		0xf4, 0xbf, 0x85, 0x81, 0xf1, 0x7f, 0x13, 0x03, 0x03, 0xcf, 0xff, 0x46, 0x10, 0xaf, 0x89, 0x81,
		0x91, 0x85, 0x81, 0xe1, 0x1f, 0xc7, 0xfd, 0xff, 0x2d, 0x0c, 0x35, 0x82, 0x8c, 0x4c, 0x8c, 0xa8,
		0x40, 0x90, 0x11, 0x1d, 0x60, 0x88, 0xf0, 0x10, 0xa1, 0x46, 0x4c, 0x90, 0x91, 0x51, 0x41, 0x49,
		0x45, 0x8b, 0x91, 0x91, 0x51, 0x83, 0x51, 0xcd, 0x88, 0x51, 0x87, 0x51, 0x8e, 0x08, 0x5d, 0x98,
		0x22, 0x36, 0x18, 0x22, 0x0e, 0x82, 0x8c, 0x8c, 0x2e, 0x44, 0x9a, 0xec, 0x81, 0x22, 0xe2, 0x43,
		0x9a, 0x7b, 0x02, 0xa0, 0x22, 0x61, 0x94, 0xfb, 0x82, 0x54, 0x91, 0x28, 0x10, 0x11, 0x03, 0x13,
		0x49, 0xc2, 0x50, 0x93, 0x82, 0x21, 0x92, 0xc6, 0xc8, 0xc8, 0x98, 0x81, 0xa2, 0x0b, 0x01, 0x72,
		0xb0, 0xda, 0x55, 0x80, 0x11, 0x1a, 0x25, 0x44, 0x84, 0x4f, 0x15, 0x79, 0x71, 0x8a, 0x11, 0x86,
		0xff, 0x9b, 0x88, 0x35, 0xe8, 0x7f, 0xd7, 0xff, 0x1e, 0x18, 0xb3, 0x8d, 0x88, 0xf8, 0xf8, 0x3f,
		0x89, 0x08, 0x47, 0xfd, 0x9f, 0x46, 0x51, 0x1c, 0xfd, 0x9f, 0x05, 0x37, 0x68, 0x0e, 0xaa, 0xaa,
		0xff, 0x4b, 0x18, 0xff, 0x2f, 0x62, 0xfc, 0xbf, 0x80, 0xe6, 0xc9, 0xe4, 0xff, 0x0a, 0x54, 0x27,
		0xfc, 0x5f, 0x45, 0xa5, 0x98, 0x21, 0x46, 0xcd, 0xff, 0x2d, 0xc4, 0x58, 0x86, 0xa9, 0x6d, 0x17,
		0x35, 0xdd, 0xf8, 0xff, 0x00, 0x5a, 0xc2, 0x20, 0x32, 0xe4, 0x18, 0x19, 0xff, 0x1f, 0x23, 0x4f,
		0x1f, 0xbd, 0x45, 0xfe, 0x9f, 0x63, 0xfc, 0x7f, 0x06, 0x92, 0x98, 0xfe, 0x5f, 0xc0, 0x17, 0x74,
		0x01, 0xe4, 0xda, 0xf7, 0xff, 0x0a, 0xb1, 0x51, 0x77, 0x83, 0x0a, 0x85, 0x39, 0x05, 0xc5, 0x29,
		0x05, 0x85, 0x30, 0xbe, 0xd8, 0xfe, 0xff, 0x80, 0x34, 0xa3, 0x9e, 0xe0, 0x31, 0xea, 0x19, 0xf9,
		0x85, 0xe1, 0x1b, 0xda, 0xe5, 0xdd, 0xc1, 0x26, 0x82, 0xe9, 0xf9, 0x2f, 0xb4, 0xf1, 0x3c, 0x20,
		0x00, 0x00, 0xff, 0xff, 0x0d, 0x67, 0xeb, 0x5a, 0xfd, 0x08, 0x00, 0x00,
	}
	buf, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(&tab); err != nil {
		panic(err)
	}
	for i := 0; i < numStates; i++ {
		for j := 0; j < numNTSymbols; j++ {
			gotoTab[i][j] = tab[i][j]
		}
	}
}
