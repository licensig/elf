package elf_test

import (
	"testing"

	"github.com/tahirmurata/aoc/elf"
)

func TestNew(t *testing.T) {
	tests := []struct {
		token string
		want  *elf.Elf
	}{
		{"token", &elf.Elf{
			Token: "token",
		}},
	}

	for _, tt := range tests {
		testname := tt.token
		t.Run(testname, func(t *testing.T) {
			ans := elf.New(tt.token)
			// if ans != tt.want {
			if ans == tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
