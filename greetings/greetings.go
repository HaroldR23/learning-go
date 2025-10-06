package greetings

import (
	"fmt"
	"errors"
)

func Hello(name string) (string, error) {
	// The operator := helps us to declare and initialize a variable in one line.
	// OTherwise, we would have to do it in two lines like this:
	// var message string First  we should declare the variable
	// message = fmt.Sprintf("Hi, %v. Welcome!", name) Then we should initialize it
	// return message


	if name == "" {
		return "", errors.New("Empty name")
	}
	
	// I find the sintaxis of showing variables in console pretty similar to C's printf, since we use %v to show the value of a variable. and in the second function's argument we pass the name of the variable that actually has the value we want to show.

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}