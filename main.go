package main

import (
	"github.com/MihailShev/algo-1/pkg/tester"
	"github.com/MihailShev/algo-1/tasks"
)

const path1 = "test-data/0.String"
const path2 = "test-data/1.Tickets"

func main() {
	//task1 := tasks.StringLength{}
	tasks.LuckyTickets{}.Run([]string{"4"})
	//test1 := tester.NewTester(task1, path1)
	test := tester.NewTester(tasks.LuckyTickets{}, path2)
	test.RunTest()
}
