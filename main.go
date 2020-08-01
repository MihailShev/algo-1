package main

import (
	"github.com/MihailShev/algo-1/pkg/tester"
	"github.com/MihailShev/algo-1/tasks"
)

const path1 = "test-data/0.String"

func main() {
	task1 := tasks.StringLength{}
	test := tester.NewTester(task1, path1)
	test.RunTest()
}
