package main

import (
	"fmt"
	"github.com/MihailShev/algo-1/pkg/tester"
	"github.com/MihailShev/algo-1/tasks"
)

const path1 = "test-data/0.String"
const path2 = "test-data/1.Tickets"

func main() {
	fmt.Printf("\n\n*** Test string length ***\n\n")

	testLength := tester.NewTester(tasks.StringLength{}, path1)
	testLength.RunTest()

	fmt.Printf("\n\n*** Test lucky tickets ***\n\n")

	testLuckyTickets := tester.NewTester(tasks.LuckyTickets{}, path2)
	testLuckyTickets.RunTest()

	_, _ = fmt.Scanf(" ")
}
