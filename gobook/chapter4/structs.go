package main

import (
	"fmt"
	"time"
)

// Employee base class to manage employees
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

// Point represent point in decart plane
type Point struct {
	X, Y float32
}

// Circle represent a circle :)
type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	NumberOfSpokes int
}

func main() {
	// Cannot be used in such way, because no fields
	// wheel := &Wheel{X: 10, Y: 20, Radius: 11, NumberOfSpokes: 6}

	wheel := &Wheel{
		Circle: Circle{
			Point:  Point{X: 10, Y: 20},
			Radius: 11,
		},
		NumberOfSpokes: 6,
	}

	fmt.Println(wheel.X)
	fmt.Println(wheel.Y)
}

func EmployeeById(id int) *Employee {
	return &Employee{ID: id}
}
