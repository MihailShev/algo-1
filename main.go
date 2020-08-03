package main

import (
	"fmt"
	"github.com/MihailShev/algo-1/pkg/tester"
	"github.com/MihailShev/algo-1/tasks"
)

const path1 = "test-data/0.String"
const path2 = "test-data/1.Tickets"

func main() {
	fmt.Println("Test string length")

	testLength := tester.NewTester(tasks.StringLength{}, path1)
	testLength.RunTest()

	fmt.Printf("=========================\n=========================\n")
	fmt.Println("Test lucky tickets")

	testLuckyTickets := tester.NewTester(tasks.LuckyTickets{}, path2)
	testLuckyTickets.RunTest()
}
