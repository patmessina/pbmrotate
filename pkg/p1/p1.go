package p1

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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

		// check if file is completely empty
		if len(lines[i]) == i && (len(lines[i]) == 0 || lines[i][0] == '#') {
			return nil, errors.New("file is not the correct format")
		}

		// ignore newlines or comments
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

// Rotates the image
func (img *P1Image) Rotate(deg int) error {

	if deg%90 != 0 {
		msg := "cannot rotate " + strconv.Itoa(deg) + " degrees"
		return errors.New(msg)
	}

	if deg%360 == 0 {
	} else if deg%270 == 0 {
		// if negative rotate right
		if deg < 0 {
			img.rotate(true)
		} else {
			img.rotate(false)
		}
	} else if deg%180 == 0 {
		img.flip()
	} else if deg%90 == 0 {
		// if negative rotate left
		if deg < 0 {
			img.rotate(false)
		} else {
			img.rotate(true)
		}

	} else {
		return errors.New("If you are here, I suck at math.")
	}
	return nil
}

// rotate - rotate Data clockwise or counter clockwise 90 deg
func (img *P1Image) rotate(clockwise bool) {
	rotatedImage := make([][]bool, img.Col)
	for i := 0; i < img.Col; i++ {
		rotatedImage[i] = make([]bool, img.Row)
	}

	imgDataRow := img.Row - 1
	imgDataCol := img.Col - 1
	for i := 0; i < img.Row; i++ {
		for j := 0; j < img.Col; j++ {
			if clockwise {
				rotatedImage[j][imgDataRow-i] = img.Data[i][j]
			} else {
				rotatedImage[imgDataCol-j][i] = img.Data[i][j]
			}
		}
	}
	img.Data = rotatedImage
	img.Row, img.Col = img.Col, img.Row
}

// flip will rotate Data by 180 degrees
func (img *P1Image) flip() {
	i := 0
	j := img.Row - 1
	// swap rows
	for {
		if i >= j {
			break
		}
		img.Data[i], img.Data[j] = img.Data[j], img.Data[i]
		i++
		j--
	}
	// swap col
	for r := 0; r < len(img.Data); r++ {
		i = 0
		j = img.Col - 1
		for {
			if i >= j {
				break
			}
			img.Data[r][i], img.Data[r][j] = img.Data[r][j], img.Data[r][i]
			i++
			j--
		}
	}
}

// getFormatedData formats data into a string that we can write to
// stdout or disk
func (img *P1Image) getFormatedData() strings.Builder {
	builder := strings.Builder{}

	builder.WriteString("P1\n")
	size := strconv.Itoa(img.Col) + " " +
		strconv.Itoa(img.Row) + "\n"
	builder.WriteString(size)
	for i, r := range img.Data {
		for j, v := range r {
			if v {
				builder.WriteString("1")
			} else {
				builder.WriteString("0")
			}
			// append space between elements
			if j < len(r)-1 {
				builder.WriteString(" ")
			}
		}
		// add newline at the end of each row
		if i < len(img.Data)-1 {
			builder.WriteByte('\n')
		}
	}

	return builder
}

// Print prints out P1Image -- TODO: This can be abstracted away from P1Image
func (img *P1Image) Print() {
	builder := img.getFormatedData()
	fmt.Println(builder.String())
}

// WriteToFile will write to a give path
func (img *P1Image) WriteToFile(path string) {
	builder := img.getFormatedData()
	whatever := []byte(builder.String())
	os.WriteFile(path, whatever, 666)
}
