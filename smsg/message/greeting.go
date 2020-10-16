package message

import "fmt"

//Greeting is a function returns a greeting string
func Greeting(name, message string) (salutation string) {
	salutation = fmt.Sprintf("%s , %s", message, name)
	return
}
