package day4

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ReadFile(t *testing.T, fileName string) []byte {
	t.Helper()

	// Get the directory of this test file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	filePath := filepath.Join(dir, fileName)

	data, err := os.ReadFile(filePath)
	assert.NoError(t, err, "failed to read file: %v", err)
	return data
}

const newLn = byte('\n')

func TransformToMatrix(input []byte) [][]byte {
	rowSize := slices.Index(input, newLn)
	out := make([][]byte, rowSize)

	for i, b := range input {
		if b == '\n' {
			continue
		}
		row := i / (rowSize + 1)
		out[row] = append(out[row], b)
	}

	return out
}

func PrintMatrix(input [][]byte) {
	out := ""
	for _, row := range input {
		out += string(row) + "\n"
	}
	fmt.Println(out)
}

func PrintBytes(input []byte) {
	fmt.Println(string(input))
}

func TestFileSolution(t *testing.T) {
	tests := []struct {
		fileName string
		result   int
	}{
		{
			fileName: "input.txt",
			result:   13,
		},
		{
			fileName: "input2.txt",
			result:   1491,
		},
	}

	for _, tt := range tests {
		t.Run(tt.fileName, func(t *testing.T) {
			data := ReadFile(t, tt.fileName)
			transformedInput := TransformToMatrix(data)
			result := Solution(transformedInput)
			assert.Equal(t, tt.result, result)
		})
	}
}

func TestFileSolution2(t *testing.T) {
	tests := []struct {
		fileName string
		result   int
	}{
		{
			fileName: "input.txt",
			result:   43,
		},
		{
			fileName: "input2.txt",
			result:   8722,
		},
	}

	for _, tt := range tests {
		t.Run(tt.fileName, func(t *testing.T) {
			data := ReadFile(t, tt.fileName)
			transformedInput := TransformToMatrix(data)
			result := Solution2(transformedInput)
			assert.Equal(t, tt.result, result)
		})
	}
}
