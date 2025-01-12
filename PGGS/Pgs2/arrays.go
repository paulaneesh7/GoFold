package pgs2


import "fmt"


func Array_basics() {
	fmt.Println("Arrays in Go")

	// Creating Array
	prices := [4]float32{10.99, 5.45, 67.89, 1.23}
	fmt.Println(prices)

	var productNames = [4]string{"Pen", "Pencil", "Eraser", "Sharpner"}
	fmt.Println(productNames)


	// Accessing values from Array
	fmt.Println("First product name: ", productNames[0])
	fmt.Println("Thrid value in price array: ", prices[2]) // 67.89 as arrays starts with index 0


	// Updating values in Array
	productNames[0] = "Ball Pen"
	fmt.Println("Updated product name: ", productNames[0], "," , "Array after updating: ", productNames)


	// Length of Array
	fmt.Println("Length of productNames array: ", len(productNames))


	// Slice of Array
	fmt.Println("Slice of productNames array: ", productNames[1:3]) // [Pencil Eraser]
	fmt.Println("We can also do this with slices:", prices[:4])
	fmt.Println("And also this with slices:", prices[0:])
	fmt.Println("And this:", prices[:])


}

func Dynamic_Arrays() {
	fmt.Println("Dynamic Arrays in Go")


	// Creating Dynamic Array (not mentioning size)
	prices := []float64{10.99, 5.45, 67.89, 1.23}
	fmt.Println(prices[1])
}