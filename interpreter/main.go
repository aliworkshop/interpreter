package main

import (
	"grpctest/interpreter/domain"
	"strconv"
	"strings"
	"time"
)

const input = `print "basic: started"

print "basic: wait for 1 sec"
await sleep 10000

print "basic: finished"`

func main() {

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) > 0 {
			p := new(parse)
			p.line = line
			p.Parse().Run()
		}
	}
}

type parse struct {
	line string
}

func (p parse) Parse() domain.Runner {
	if strings.HasPrefix(p.line, "print") {
		_, after, _ := strings.Cut(p.line, "print")
		return &domain.Print{Code: strings.TrimSpace(after)}
	} else if strings.HasPrefix(p.line, "await sleep") {
		_, after, _ := strings.Cut(p.line, "await sleep")
		sleepSec, _ := strconv.Atoi(strings.TrimSpace(after))
		return &domain.Sleep{Time: time.Duration(sleepSec) * time.Millisecond}
	}
	return nil
}
