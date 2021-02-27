package cmd

import "fmt"

func Hello(value string) string {
	message:= fmt.Sprintf("Hi, %v. Welcome!", value)
	return message
}
