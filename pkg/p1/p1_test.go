package p1

import (
	"testing"
)

var (
	letterJ string = `P1
# This is an example bitmap of the letter "J"
6 10
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
1 0 0 0 1 0
0 1 1 1 0 0
0 0 0 0 0 0
0 0 0 0 0 0`

	letterJExpected [][]bool = [][]bool{
		{false, false, false, false, true, false},
		{false, false, false, false, true, false},
		{false, false, false, false, true, false},
		{false, false, false, false, true, false},
		{false, false, false, false, true, false},
		{false, false, false, false, true, false},
		{true, false, false, false, true, false},
		{false, true, true, true, false, false},
		{false, false, false, false, false, false},
		{false, false, false, false, false, false}}
)

func TestNewImagePass(t *testing.T) {
	// Should pass test cases
	tcs := []struct {
		Input    []byte
		Expected *P1Image
	}{
		// Check single value
		{
			Input: []byte("P1\n1 1\n1"),
			Expected: &P1Image{
				Row:  1,
				Col:  1,
				Data: [][]bool{{true}},
			},
		},

		// check example J image
		{
			Input: []byte(letterJ),
			Expected: &P1Image{
				Row:  1,
				Col:  1,
				Data: letterJExpected,
			},
		},
		// check empty value
		{
			Input: []byte("P1\n1 0\n0\n"),
			Expected: &P1Image{
				Row:  1,
				Col:  1,
				Data: [][]bool{},
			},
		},
	}

	for _, c := range tcs {
		image, err := NewImage(c.Input)
		if err != nil {
			t.Errorf("failed to get image: %v", err)
			break
		}
		if len(image.Data) != len(c.Expected.Data) {
			t.Errorf("data size %v does not match %v",
				len(image.Data), len(c.Expected.Data))
			break
		}

		for i, expectedRow := range c.Expected.Data {
			if len(image.Data[i]) != len(expectedRow) {
				t.Errorf("data size %v does not match %v",
					len(image.Data), len(c.Expected.Data))
				break
			}
			for j, val := range expectedRow {
				// actually checking that the values are the same
				if val != image.Data[i][j] {
					t.Errorf("Expected %v received %v at index [%v][%v]",
						val, image.Data[i][j], i, j)
					break
				}
			}
		}
	}
}

func TestNewImageFromFile(t *testing.T) {
}

func TestCreateImage(t *testing.T) {
}

func TestRotate(t *testing.T) {
}

func TestFlip(t *testing.T) {
}

func TestGetFormatedData(t *testing.T) {
}

func TestWriteToFile(t *testing.T) {
}
