package day3

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestSolution(t *testing.T) {
	input := []byte("12345\n12367\n11119\n99999") // 45 + 67 + 19 + 99
	got := Solution(input, 2)
	want := 230
	if got != want {
		t.Errorf("got %d, want, %d", got, want)
	}
}

func ReadFile(t *testing.T, fileName string) []byte {
	t.Helper()

	// Get the directory of this test file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	filePath := filepath.Join(dir, fileName)

	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	return data
}

func TestFileSolution(t *testing.T) {
	t.Run("input.txt", func(t *testing.T) {
		data := ReadFile(t, "input.txt")
		got := Solution(data, 2)
		want := 357
		if got != want {
			t.Errorf("got %d, want, %d", got, want)
		}
	})

	t.Run("input2.txt", func(t *testing.T) {
		data := ReadFile(t, "input2.txt")
		got := Solution(data, 2)
		want := 17330
		if got != want {
			t.Errorf("got %d, want, %d", got, want)
		}
	})
}

func TestSolution2(t *testing.T) {
	input := []byte("511172113\n511172113\n224587191") // 1446
	got := Solution(input, 3)
	want := 2337
	if got != want {
		t.Errorf("got %d, want, %d", got, want)
	}
}

func TestFileSolution2(t *testing.T) {
	t.Run("input.txt", func(t *testing.T) {
		data := ReadFile(t, "input.txt")
		got := Solution(data, 12)
		want := 3121910778619
		if got != want {
			t.Errorf("got %d, want, %d", got, want)
		}
	})

	t.Run("input2.txt", func(t *testing.T) {
		data := ReadFile(t, "input2.txt")
		got := Solution(data, 12)
		want := 171518260283767
		if got != want {
			t.Errorf("got %d, want, %d", got, want)
		}
	})
}
