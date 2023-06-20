package main

import "fmt"

func main() {
	salaries := map[string]int{
		"test":  15000,
		"test1": 16000,
	}

	salaries["test3"] = 25000
	for key, value := range salaries {
		fmt.Println(key, value)
	}
	delete(salaries, "test1")
	fmt.Println(salaries)
}
