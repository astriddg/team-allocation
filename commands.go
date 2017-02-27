package main

import (
	"github.com/peterh/liner"
)

type Command interface {

	Action(line *liner.State, args []string) error
}

type add struct{}

type del struct{}

type show struct{}

type gen struct{}

type help struct{}

