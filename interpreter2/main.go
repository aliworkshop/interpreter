package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"grpctest/interpreter2/domain"
	"strconv"
	"strings"
	"time"
)

type Input struct {
	lines  []string
	start  chan bool
	paused bool
}

var files = []string{
	`print "file1: started"

print "file1: wait for 1 sec"
await sleep 1000

print "file1: wait for 2 sec"
await sleep 2000

print "file1: wait for 3 sec"
await sleep 3000

print "file1: wait for 4 sec"
await sleep 4000

print "file1: finished"`,

	`print "file2: started"

print "file2: wait for 4 sec"
await sleep 4000

print "file2: wait for 3 sec"
await sleep 3000

print "file2: wait for 2 sec"
await sleep 2000

print "file2: wait for 1 sec"
await sleep 1000

print "file2: finished"`,

	`print "file3: started"

print "file3: wait for 5 sec"
await sleep 5000

print "file3: wait for another 5 sec"
await sleep 5000

print "file3: finished"`}

func main() {

	inputs := make([]Input, 0)
	for _, file := range files {
		channel := make(chan bool, 1)
		inputs = append(inputs, Input{lines: strings.Split(file, "\n"), start: channel})
	}
	eg, _ := errgroup.WithContext(context.Background())

	for _, input := range inputs {
		input.start <- true
		eg.Go(func(in Input) func() error {
			return func() error {
				in.Work()
				return nil
			}
		}(input))
	}
	if err := eg.Wait(); err != nil {
		panic(err)
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

func (in *Input) Work() {
	if !in.paused {
		<-in.start
		for i, line := range in.lines {
			if len(line) > 0 {
				p := new(parse)
				p.line = line
				if p.Parse().ShoutPause() {
					go p.Parse().Run()
					in.paused = true
					in.lines = in.lines[i:]
					break
				}
				p.Parse().Run()
			}
		}
	}
}
