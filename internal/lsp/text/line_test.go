package text

import (
	"bytes"
	"testing"
)

const testFile1 = `123
56αβ9

CDE
`

// testFile2 describes a file that does not end with '\n'
const testFile2 = `12345
678`

const testFile3 = `hello`

func TestLineOffsets(t *testing.T) {
	var testCases = []struct {
		file              string
		offset, line, col int
	}{
		{testFile1, 0x0, 0, 0},
		{testFile1, 0x1, 0, 1},
		{testFile1, 0x2, 0, 2},
		{testFile1, 0x3, 0, 3},
		{testFile1, 0x4, 1, 0},
		{testFile1, 0x5, 1, 1},
		{testFile1, 0x6, 1, 2},
		{testFile1, 0x7, 1, 3},
		{testFile1, 0x8, 1, 4},
		{testFile1, 0x9, 1, 5},
		{testFile1, 0xA, 2, 0},
		{testFile1, 0xB, 3, 0},
		{testFile1, 0xC, 3, 1},
		{testFile1, 0xD, 3, 2},
		{testFile1, 0xE, 3, 3},
		{testFile1, 0xF, 4, 0},
		{testFile2, 0x5, 0, 5},
		{testFile2, 0x6, 1, 0},
		{testFile2, 0x7, 1, 1},
		{testFile2, 0x8, 1, 2},
		{testFile2, 0x9, 1, 3},
		{testFile3, 0x0, 0, 0},
		{testFile3, 0x1, 0, 1},
		{testFile3, 0x5, 0, 5},
	}

	for _, tc := range testCases {
		off, err := getNewlineOffsets(bytes.NewBufferString(tc.file))
		if err != nil {
			t.Errorf("failed to compute file offsets: %v", err)
			continue
		}
		if o := off.LineToOffset(tc.line, tc.col); o != tc.offset {
			t.Errorf("LineToOffset(%v, %v) = %v for off=%v; expected %v\n",
				tc.line, tc.col, o, off, tc.offset)
		}
		if line, col := off.OffsetToLine(tc.offset); line != tc.line || col != tc.col {
			t.Errorf("OffsetToLine(%v) = %v, %v for off=%v; expected %v, %v\n",
				tc.offset, line, col, off, tc.line, tc.col)
		}
	}
}

func TestLineOffsetsLeftover(t *testing.T) {
	var testCases = []struct {
		file              string
		offset, line, col int
	}{
		{testFile2, 0x9, 1, 4},
		{testFile2, 0x9, 2, 0},
		{testFile2, 0x9, 2, 1},
		{testFile2, 0x9, 2, 2},
		{testFile2, 0x9, 3, 0},
		{testFile2, 0x9, 3, 1},
		{testFile2, 0x9, 3, 2},
		{testFile2, 0x9, 4, 0},
		{testFile2, 0x9, 4, 1},
		{testFile2, 0x9, 4, 2},
		{testFile3, 0x5, 0, 6},
		{testFile3, 0x5, 0, 7},
		{testFile3, 0x5, 1, 0},
		{testFile3, 0x5, 1, 2},
	}

	for _, tc := range testCases {
		off, err := getNewlineOffsets(bytes.NewBufferString(tc.file))
		if err != nil {
			t.Errorf("failed to compute file offsets: %v", err)
			continue
		}
		if o := off.LineToOffset(tc.line, tc.col); o != tc.offset {
			t.Errorf("LineToOffset(%v, %v) = %v for off=%v; expected %v\n",
				tc.line, tc.col, o, off, tc.offset)
		}
	}
}
