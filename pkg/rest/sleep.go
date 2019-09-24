package rest

import "fmt"

type SleepFramework struct{}

func NewSleeper() SleepFramework {
	return SleepFramework{}
}

func (SleepFramework) Rest() {
	fmt.Println("This rest is sleep")
}
