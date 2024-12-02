package elf_test

import (
	"testing"

	"github.com/licensig/elf"
)

func TestGetInputFile(t *testing.T) {
	tests := []struct {
		name string
		year int
		day  int
	}{
		{
			name: "2024-1.txt",
			year: 2024,
			day:  1,
		},
		{
			name: "2024-2.txt",
			year: 2024,
			day:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elf := elf.New("token")
			file := elf.GetInputFile(tt.year, tt.day)
			defer file.Close()
			if file == nil {
				t.Errorf("GetInput() = %v; want a file", file)
			}
		})
	}
}
