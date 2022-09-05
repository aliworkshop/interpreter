package main

import (
	"interpreter2/domain"
	"strings"
)

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
	list := domain.List{}
	for _, file := range files {
		list.Insert(strings.Split(file, "\n"))
	}
	domain.ConvertSinglyToCircular(&list)

	t := domain.Task{}
	t.Add(len(files))
	input := list.Head()
	for t.HasTask() {
		input.Work()
		input = input.Next()
		if input.Finished() {
			t.Done()
		}
	}
}
