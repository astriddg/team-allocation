package main


type Command interface {
	Action(args []string) error
}

type add struct{}

type del struct{}

type show struct{}

type gen struct{}

type help struct{}
