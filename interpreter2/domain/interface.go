package domain

import (
	"fmt"
	"time"
)

type Sleep struct {
	Time time.Duration
}

type Print struct {
	Code string
}

type Parser interface {
	Parse() Runner
}

type Runner interface {
	ShouldPause() bool
	Run()
}

func (s Sleep) Run() {
	time.Sleep(s.Time)
}

func (s Sleep) ShouldPause() bool {
	return true
}

func (p Print) Run() {
	fmt.Println(p.Code)
}

func (p Print) ShouldPause() bool {
	return false
}
