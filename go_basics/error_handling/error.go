package error_handling

import "fmt"

func Error() {
	fmt.Println("Error Handling in Go")
	ans, _ := divide(10,4)
	fmt.Println("Division is: ", ans)

	// we can ignore the error by using _
	result, err := divide (6,4)
	fmt.Println("Division is: ", result, "error is: ", err)
}

func divide(a, b float64) (float64, error){
	if b == 0 {
		return 0, fmt.Errorf("denominator cannot be zero")
	}
	return a/b, nil
}