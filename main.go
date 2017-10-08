package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	leader := new(GroupLeader)
	leader.score = 100
	leader.name = "Egor"

	student := &Student{name: "Olga", score: 77}

	fmt.Println(totalScores(leader, student))
}

func calculate(heightMeters float32, weight float32) float32 {
	return weight / (heightMeters * heightMeters)
}

func toCelcium(farengheigt float32) float32 {
	celcium := (farengheigt - 32) * 5.0/9.0
	return celcium
}

func counToN(number int) {
	iter := 1
	for iter <= number {
		fmt.Println(iter)
		iter += 1
	}
}

func anotherCountToN(number int) {
	for i := 1; i <= number; i += 1 {
		fmt.Println(i)
	}
}

func printEvenToN(number int) {
	for i := 1; i <= number; i += 1 {
		var typeDelimited string

		if i % 2 == 0 {
			typeDelimited = "Even"
		} else {
			typeDelimited = "Odd"
		}

		fmt.Println(i, typeDelimited)
	}
}

func switchCase(number int) {
	switch number {
	case 0: fmt.Println("Zero")
	case 1: fmt.Println("One")
	case 2: fmt.Println("Two")
	default:
		fmt.Println("More then 2 or less then 0")
	}
}

func userArray(numberOfElems int) {
	array := make([]int, numberOfElems)
	arrLen := len(array)
	for i := 0; i < arrLen; i += 1 {
		array[i] = rand.Int() + i
	}

	for i := 0; i < arrLen; i += 1 {
		fmt.Println(i)
	}
}

func anotherFormOfFor() (float32){
	ages := []int{24, 33, 17}
	total := 0
	for _, value := range ages {
		total += value
	}

	fmt.Println("Total is", total)

	return float32(total) / float32(len(ages))
}

func useSlices() {
	x := []int{1,2,3,4,5}
	fmt.Println(x[0:2])
	appendedSlice := append(x, 6, 7)

	fmt.Println(appendedSlice)

	dest := make([]int, 3)
	copy(dest, x[1:])

	fmt.Println(dest)
}

func useMaps() {
	studentScores := make(map[string]int)

	studentScores["Egor"] = 100
	studentScores["Oleg"] = 87
	studentScores["Inna"] = 93

	score, ok := studentScores["dsf"]
	fmt.Println(score, ok)

	for name, score := range studentScores {
		fmt.Println("Studen", name, "has score", score)
	}
}

func maxmin(xs []int) (int, int) {
	smallest := xs[0]
	biggest := xs[0]
	for _, v := range xs {
		if v > biggest {
			biggest = v
		}

		if v < smallest {
			smallest = v
		}
	}

	return biggest, smallest
}

func addVariadic(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	return total
}

func generateAddFunc(number int) func(int)int {
	sumFunc := func(n int) int {
		return number + n
	}

	return sumFunc
}

func handleErrors() {
	defer func() {
		str := recover()
		fmt.Println("Runtime error occurs", str)
	}()

	panic("Hello world")
}

func intPtr(xPtr *int) {
	*xPtr++
}

//Structs and interfaces
type Student struct {
	name string
	score uint
}

func (c *Student) doubleScores() {
	c.score = c.score * 2
}

func (c *Student) getScore() uint {
	return c.score
}

//is a
type GroupLeader struct {
	Student
	website string
}

type Scorable interface {
	getScore() uint
}

func totalScores(scorable ...Scorable) uint {
	total := uint(0)
	for _, sc := range scorable {
		total += sc.getScore()
	}

	return total
}


