package p1

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// P1Image - PBM P1 formatted image
type P1Image struct {
	Row  int
	Col  int
	Data [][]bool
}

// NewImage Given a byte slice of a p1 image parse and return a
// types.P1Image
func NewImage(content []byte) (*P1Image, error) {
	image := &P1Image{}
	var err error

	// split and trim content
	lines := bytes.Split(content, []byte{'\n'})
	for i, l := range lines {
		lines[i] = bytes.Trim(l, "\n")
	}

	// we want to keep our place later
	i := 0
	// we want to make sure to skip any comments and blank lines
	for ; i < len(lines); i++ {
		if len(lines[i]) == 0 || lines[i][0] == '#' {
			continue
		}
		// If the first line after any comments is the correct
		// format continue on with the program
		if string(lines[i]) == "P1" {
			i++
			break
		}
		return nil, errors.New("file is not the correct format")

	}

	// similarly we want to make sure to ignore comments and
	// newlines
	for ; i < len(lines); i++ {
		if len(lines[i]) == 0 || lines[i][0] == '#' {
			continue
		}
		numbers := bytes.Split(lines[i], []byte{' '})
		if len(numbers) != 2 {
			return nil, errors.New("missing picture size")
		}

		image.Col, err = strconv.Atoi(string(numbers[0]))
		if err != nil {
			return nil, err
		}
		image.Row, err = strconv.Atoi(string(numbers[1]))
		if err != nil {
			return nil, err
		}
		i++
		break
	}

	data := []byte{}
	for ; i < len(lines); i++ {
		// skip blank lines and comments
		if len(lines[i]) == 0 || lines[i][0] == '#' {
			continue
		}
		data = append(data, lines[i]...)
	}

	data = bytes.ReplaceAll(data, []byte(" "), []byte(""))
	image.Data, err = CreateImage(data, image.Col, image.Row)
	if err != nil {
		return nil, err
	}

	return image, nil

}

// NewImageFromFile takes a PBM file path string and returns an
// PBMImage datastructure
func NewImageFromFile(path string) (*P1Image, error) {

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return NewImage(content)
}

// CreateImage takes a trimmed byte slice and converts int to a 2
// dimensional boolean array
func CreateImage(data []byte, col, row int) ([][]bool, error) {
	image := make([][]bool, row)
	for i := 0; i < row; i++ {
		image[i] = make([]bool, col)
		for j := 0; j < col; j++ {
			digit, err := strconv.Atoi(string(data[col*i+j]))
			if err != nil {
				return nil, err
			}
			image[i][j] = digit == 1
		}
	}

	return image, nil
}

// getFormatedData formats data into a string that we can write to
// stdout or disk
func (img *P1Image) getFormatedData() strings.Builder {
	builder := strings.Builder{}

	builder.WriteString("P1\n")
	size := strconv.Itoa(len(img.Data[0])) + " " +
		strconv.Itoa(len(img.Data)) + " \n"
	builder.WriteString(size)
	for _, r := range img.Data {
		for _, v := range r {
			if v {
				builder.WriteString("1 ")
			} else {
				builder.WriteString("0 ")
			}
		}
		builder.WriteByte('\n')
	}

	return builder
}

// Print prints out P1Image -- TODO: This can be abstracted away from P1Image
func (img *P1Image) Print() {
	builder := img.getFormatedData()
	fmt.Println(builder.String())
}
