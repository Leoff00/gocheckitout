package logs

import "fmt"

func WelcomeLog() {
	fmt.Println(
		"Welcome to go Check-it-out, Choose an option to execute the features", "\n",
		"1 - Only make a request with no details", "\n",
		"2 - Make a request with some details", "\n",
		"3 - For exit CLI menu",
	)
}

func InvalidOperation() {
	fmt.Println("Invalid operation, exiting...")
}

func Exiting() {
	fmt.Println("Exiting...")
}
