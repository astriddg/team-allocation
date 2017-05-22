package main

import "github.com/nlopes/slack"

type Command interface {
	Action(rtm *slack.RTM, args []string) error
}

type add struct{}

type del struct{}

type show struct{}

type gen struct{}

type help struct{}
