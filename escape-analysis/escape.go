package main

import (
	"fmt"
	"github.com/pkg/errors"
	"regexp"
	"strings"

	"github.com/jedib0t/go-pretty/text"

	"github.com/jedib0t/go-pretty/table"
)

type Variable struct {
	Stacks []string
	Heaps  []string
}

type EscapeAnalysis struct {
	PathRegex  *regexp.Regexp
	StackRegex []*regexp.Regexp
	HeapRegex  []*regexp.Regexp
	CodePath   []string
	Result     map[string]*Variable
}

func NewEscapeAnalysis() (*EscapeAnalysis, error) {
	pathR, err1 := regexp.Compile(`^(.*):\d+:\d+`)
	heapR1, err2 := regexp.Compile(`:[[:space:]](.*) escapes to heap$`)
	heapR2, err3 := regexp.Compile(`moved to heap:[[:space:]](.*)$`)
	stackR, err4 := regexp.Compile(`:[[:space:]](.*) does not escape$`)

	for _, err := range []error{err1, err2, err3, err4} {
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse regex")
		}
	}

	return &EscapeAnalysis{
		PathRegex:  pathR,
		HeapRegex:  []*regexp.Regexp{heapR1, heapR2},
		StackRegex: []*regexp.Regexp{stackR},
		Result:     make(map[string]*Variable),
	}, nil
}

func (ea *EscapeAnalysis) Run(data string) {
	s := strings.Split(data, "\n")
	for _, val := range s {
		if ea.PathRegex.MatchString(val) {
			path := ea.PathRegex.FindStringSubmatch(val)[0]
			for _, stack := range ea.StackRegex {
				if stack.MatchString(val) {
					v, ok := ea.Result[path]
					if !ok {
						v = &Variable{
							Stacks: nil,
							Heaps:  nil,
						}
						ea.CodePath = append(ea.CodePath, path)
					}
					v.Stacks = append(v.Stacks, stack.FindStringSubmatch(val)[1])
					ea.Result[path] = v
				}
			}

			for _, heap := range ea.HeapRegex {
				if heap.MatchString(val) {
					v, ok := ea.Result[path]
					if !ok {
						v = &Variable{
							Stacks: nil,
							Heaps:  nil,
						}
						ea.CodePath = append(ea.CodePath, path)
					}
					v.Heaps = append(v.Heaps, heap.FindStringSubmatch(val)[1])
					ea.Result[path] = v
				}
			}
		}
	}
}

func (ea *EscapeAnalysis) String() string {
	tw := table.NewWriter()
	tw.SetTitle("ESCAPE ANALYSIS RESULTS")
	tw.AppendHeader(table.Row{"CODE", "STACK", "HEAP"})
	var ts []table.Row
	for _, code := range ea.CodePath {
		var stackResult string
		stacks := ea.Result[code].Stacks
		for _, s := range stacks {
			stackResult += fmt.Sprintf("%s,", s)
		}
		var heapResult string
		heaps := ea.Result[code].Heaps
		for _, h := range heaps {
			heapResult += fmt.Sprintf("%s,", h)
		}
		ts = append(ts, table.Row{
			code, stackResult, heapResult,
		})
	}

	tw.AppendRows(ts)

	tw.SetStyle(table.StyleBold)
	tw.Style().Title.Align = text.AlignCenter
	tw.Style().Options.SeparateRows = true

	return tw.Render()
}
