package domain

import (
	"strconv"
	"strings"
	"time"
)

type Input struct {
	name        string
	lines       []string
	currentLine chan parse
	start       chan bool
	next        *Input
	finished    bool
}

type parse struct {
	line string
}

func (p parse) Parse() Runner {
	if strings.HasPrefix(p.line, "print") {
		_, after, _ := strings.Cut(p.line, "print")
		return &Print{Code: strings.TrimSpace(after)}
	} else if strings.HasPrefix(p.line, "await sleep") {
		_, after, _ := strings.Cut(p.line, "await sleep")
		sleepSec, _ := strconv.Atoi(strings.TrimSpace(after))
		return &Sleep{Time: time.Duration(sleepSec) * time.Millisecond}
	}
	return nil
}

func (in *Input) Work() {
	for i, line := range in.lines {
		if len(line) > 0 {
			p := parse{line: line}
			if len(in.lines)-1 == i {
				in.finished = true
				in.lines = nil
			}
			if p.Parse().ShouldPause() {
				in.lines = in.lines[i+1:]
				go p.Parse().Run()
				return
			}
			p.Parse().Run()
		}
	}
}

func (in *Input) Next() *Input {
	return in.next
}

func (in *Input) Finished() bool {
	return in.finished
}
