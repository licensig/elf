// Package elf implements helper functions for solving the Advent of Code problems
package elf

type Elf struct {
	token string
}

// New creates a new Elf with the given token
func New(token string) *Elf {
	return &Elf{token: token}
}
