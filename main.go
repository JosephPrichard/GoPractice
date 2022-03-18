package main

import (
	"fmt"
	"main/queue"
	"main/sort"
	"main/str"
)

type Person struct {
	name string
	age  int
}

func printQ(q *queue.Queue) {
	curr := q.Head
	for curr != nil {
		fmt.Printf("%+v\n", curr.Val)
		curr = curr.Next
	}
}

func main() {
	fmt.Println("Hello World")

	array := []int{9, 6, 3, 1, 10, 5, 3, 2, 4, 7, 8}

	array1 := make([]int, len(array))
	copy(array1, array)

	sort.BubbleSort(array)
	sort.MergeSort(array1)

	fmt.Println(array)
	fmt.Println(array1)

	fmt.Printf("%c\n", str.HighestFreq("hello world"))
	fmt.Printf("%c\n", str.HighestFreq("aa-bb_cc.ddd,ee`f"))

	q := queue.MakeEmpty()

	q.Push(Person{
		name: "Joseph",
		age:  20,
	})

	q.Push(Person{
		name: "Ritesh",
		age:  22,
	})

	q.Push(Person{
		name: "Vincent",
		age:  20,
	})

	q.Push(Person{
		name: "Abhishek",
		age:  20,
	})

	for q.Head != nil {
		poll, err := q.Poll()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(poll)
		}
	}
}
