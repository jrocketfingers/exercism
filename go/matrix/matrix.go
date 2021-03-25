package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	matrix [][]int
}

func (m Matrix) Rows() [][]int {
	result := make([][]int, len(m.matrix))
	for r, row := range m.matrix {
		result[r] = make([]int, len(row))
		copy(result[r], row)
	}

	return result
}

func (m Matrix) Cols() [][]int {
	nRows := len(m.matrix)
	if nRows == 0 {
		return make([][]int, 0)
	}

	transposed := make([][]int, len(m.matrix[0]))
	for c, _ := range m.matrix[0] {
		transposed[c] = make([]int, len(m.matrix))
	}

	for r, row := range m.matrix {
		for c, value := range row {
			transposed[c][r] = value
		}
	}

	return transposed
}

func (m Matrix) Set(r int, c int, value int) bool {
	if r < 0 || c < 0 || r >= len(m.matrix) || c >= len(m.matrix[r]) {
		return false
	}

	m.matrix[r][c] = value
	return true
}

func New(in string) (*Matrix, error) {
	rows := strings.Split(in, "\n")
	m := Matrix{}
	m.matrix = make([][]int, len(rows))

	for r, row := range rows {
		if len(row) == 0 {
			return &m, errors.New("Rows can't be empty")
		}

		fields := strings.Split(row, " ")

		width := len(fields)
		m.matrix[r] = make([]int, width)

		if r > 0 && len(m.matrix[r]) != len(m.matrix[r-1]) {
			return &m, errors.New("The rows are uneven!")
		}

		for c, field := range fields {
			value, err := strconv.ParseInt(field, 10, 64)
			if errors.Is(err, strconv.ErrRange) || errors.Is(err, strconv.ErrSyntax) {
				return &m, err
			}
			if err != nil {
				panic(err)
			}
			m.matrix[r][c] = int(value)
		}
	}
	return &m, nil
}
