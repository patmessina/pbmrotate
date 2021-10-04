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

	jCol             int    = 6
	jRow             int    = 10
	letterJByteSlice []byte = []byte("000010000010000010000010000010000010100010011100000000000000")

	letterJData [][]bool = [][]bool{
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

	letterJDataFlip [][]bool = [][]bool{
		{false, false, false, false, false, false},
		{false, false, false, false, false, false},
		{false, false, true, true, true, false},
		{false, true, false, false, false, true},
		{false, true, false, false, false, false},
		{false, true, false, false, false, false},
		{false, true, false, false, false, false},
		{false, true, false, false, false, false},
		{false, true, false, false, false, false},
		{false, true, false, false, false, false}}

	letterJDataCounterClock [][]bool = [][]bool{
		{false, false, false, false, false, false, false, false, false, false},
		{true, true, true, true, true, true, true, false, false, false},
		{false, false, false, false, false, false, false, true, false, false},
		{false, false, false, false, false, false, false, true, false, false},
		{false, false, false, false, false, false, false, true, false, false},
		{false, false, false, false, false, false, true, false, false, false}}

	letterJDataClock [][]bool = [][]bool{
		{false, false, false, true, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false},
		{false, false, false, true, true, true, true, true, true, true},
		{false, false, false, false, false, false, false, false, false, false}}
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
				Data: letterJData,
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
	tcs := []struct {
		Input      string
		ShouldPass bool
	}{
		{
			Input:      "",
			ShouldPass: false,
		},
		{
			Input:      "../../examples/j.pbm",
			ShouldPass: true,
		},
	}

	for _, c := range tcs {
		_, err := NewImageFromFile(c.Input)
		if err != nil && c.ShouldPass {
			t.Error(err)
		}
	}
}

func TestCreateImage(t *testing.T) {
	tcs := []struct {
		Input    []byte
		Col      int
		Row      int
		Expected [][]bool
	}{
		{
			Input:    letterJByteSlice,
			Col:      jCol,
			Row:      jRow,
			Expected: letterJData,
		},
		{
			Input:    []byte("1"),
			Col:      1,
			Row:      1,
			Expected: [][]bool{{true}},
		},
		{
			Input:    []byte(""),
			Col:      0,
			Row:      0,
			Expected: [][]bool{},
		},
	}

	for _, c := range tcs {
		output, err := CreateImage(c.Input, c.Col, c.Row)
		if err != nil {
			t.Error(err)
			break
		}

		if len(output) != len(c.Expected) {
			t.Errorf("output col size %v does not match expected %v",
				len(output), len(c.Expected))
		}
		for i, expectedRow := range c.Expected {
			if len(output[i]) != len(expectedRow) {
				t.Errorf("output row size %v does not match expected %v",
					len(output), len(c.Expected))
			}
			for j, val := range expectedRow {
				// Actual value test
				if output[i][j] != val {
					t.Errorf("Expected %v received %v at index [%v][%v]",
						val, output[i][j], i, j)
				}
			}
		}
	}
}

// Here, passing a pointer to P1Image.Data -- which means tests wont pass without
func TestRotate(t *testing.T) {
	tcs := []struct {
		Input      P1Image
		Degree     int
		Expected   P1Image
		ShouldPass bool
	}{
		{
			Input: P1Image{
				Row:  jRow,
				Col:  jCol,
				Data: letterJData,
			},
			Degree: 90,
			Expected: P1Image{
				// col and row should be switched here
				Row:  jCol,
				Col:  jRow,
				Data: letterJDataClock,
			},
			ShouldPass: true,
		},
		{
			Input: P1Image{
				Row:  jRow,
				Col:  jCol,
				Data: letterJData,
			},
			Degree: -270,
			Expected: P1Image{
				// col and row should be switched here
				Col:  jRow,
				Row:  jCol,
				Data: letterJDataClock,
			},
			ShouldPass: true,
		},
		{
			Input: P1Image{
				Row:  jRow,
				Col:  jCol,
				Data: letterJData,
			},
			Degree: -90,
			Expected: P1Image{
				// col and row should be switched here
				Col:  jRow,
				Row:  jCol,
				Data: letterJDataCounterClock,
			},
			ShouldPass: true,
		},
		{
			Input: P1Image{
				Row:  jRow,
				Col:  jCol,
				Data: letterJData,
			},
			Degree: 180,
			Expected: P1Image{
				Row:  jRow,
				Col:  jCol,
				Data: letterJDataFlip,
			},
			ShouldPass: true,
		},
		{
			Input: P1Image{
				Row:  jRow,
				Col:  jCol,
				Data: letterJData,
			},
			Degree: -180,
			Expected: P1Image{
				Row:  jRow,
				Col:  jCol,
				Data: letterJDataFlip,
			},
			ShouldPass: true,
		},
	}

	for _, c := range tcs {
		err := c.Input.Rotate(c.Degree)
		if err != nil && c.ShouldPass {
			t.Error(err)
			break
		}

		// check col and row
		if c.Input.Col != c.Expected.Col {
			t.Errorf("specified input col size %v does not match expected %v",
				c.Input.Col, c.Expected.Col)
			break
		}
		if c.Input.Row != c.Expected.Row {
			t.Errorf("specified input Row size %v does not match expected %v",
				c.Input.Row, c.Expected.Row)
			break
		}

		// check if values are in the correct place
		output := c.Input.Data
		expected := c.Expected.Data
		if len(output) != len(expected) {
			t.Errorf("output col size %v does not match expected %v",
				len(output), len(expected))
			break
		}
		for i, expectedRow := range expected {
			if len(output[i]) != len(expectedRow) {
				t.Errorf("output row size %v does not match expected %v",
					len(output), len(expected))
				break
			}
			for j, val := range expectedRow {
				// Actual value test
				if output[i][j] != val {
					t.Errorf("Expected %v received %v at index [%v][%v]",
						val, output[i][j], i, j)
					break
				}
			}
		}

		// TODO: refactor so this is not required
		c.Input.Rotate(-c.Degree)

	}

}

func TestGetFormatedData(t *testing.T) {
}

func TestWriteToFile(t *testing.T) {
}
