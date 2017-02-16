package main

import (
	"github.com/peterh/liner"
)

type Command interface {
	// Description of what your command does
	Help() string

	// The actual function that is called
	// Probably should be changed to list of strings, or list of interfaces
	F(l *liner.State, args []string) error

	// // Autocompleter function
	// Autocomplete(string, int) (string, []string, string)
}

type add struct{}

func (a add) Help() string { return "add a department or person" }

type del struct{}

func (d del) Help() string { return "delete a department or person" }

type show struct{}

func (s show) Help() string { return "Show a list of departments or people" }

type gen struct{}

func (g gen) Help() string { return "Generate the very scary matches" }
