package id3v2

import (
	"bytes"
	"github.com/bogem/id3v2/testutil"
	"testing"
)

var (
	th = TagHeader{
		FramesSize: 0,
		Version:    4,
	}
	thb = []byte{73, 68, 51, 4, 0, 0, 0, 0, 0, 0}
)

func TestParseHeader(t *testing.T) {
	parsed, err := ParseHeader(bytes.NewReader(thb))
	if err != nil {
		t.Error(err)
	}
	if *parsed != th {
		t.Fail()
	}
}

func TestFormTagHeader(t *testing.T) {
	formed := FormTagHeader()
	if err := testutil.AreByteSlicesEqual(formed, thb); err != nil {
		t.Error(err)
	}
}
