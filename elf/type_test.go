package elf_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/tahirmurata/santa/elf"
)

func TestStringToInts(t *testing.T) {
	tests := []struct {
		str, sep string
		want     []int
	}{
		{"12 345 678 910", " ", []int{12, 345, 678, 910}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.str, tt.sep)
		t.Run(testname, func(t *testing.T) {
			ans, err := elf.StringToInts(tt.str, tt.sep)
			// if ans != tt.want {
			if err != nil {
				t.Error(err)
			}
			if slices.Compare(ans, tt.want) != 0 {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkStringToInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		elf.StringToInts("12 345 678 910", " ")
	}
}
