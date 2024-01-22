package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("can't divide a number by zero")
		/* or */
		// return 0, fmt.Errorf("can't divide %d by zero", a)
	}

	return a / b, nil
}

/* custom error */
type CustomError struct {
	data string
}

func (err *CustomError) Error() string {
	return fmt.Sprintf("error occurred due to: %v", err.data)
}

func main() {
	fmt.Println("Errors in Go")

	// result, err := divide(4, 2)
	result, err := divide(4, 0)

	fmt.Println(result)
	fmt.Println(err)

	result1, err1 := divide(3, 0)

	if err1 != nil {
		var custErr CustomError
		custErr.data = "zero se divide kar rha hai"
		fmt.Println(custErr.Error())
		return
	}

	fmt.Println(result1)
}
