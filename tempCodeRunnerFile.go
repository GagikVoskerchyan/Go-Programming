package main

import "fmt"

func main() {
	const name = "Mike"
	const openRate = 30.5

	msg := "Hi %s, your open rate is %.2f percent"

	fmt.Printf(msg,name,openRate)
}
