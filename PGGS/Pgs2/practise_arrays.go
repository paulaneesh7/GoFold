package pgs2


import "fmt"

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

type Product struct {
	Id int
	Title string
	Price float64
}


func Practise_Arrays(){
	fmt.Println("Practise Arrays In Go")

	hobbies := [3]string{"Coding", "Watching Anime", "Reading Tech Blogs and Books"}
	fmt.Println("Hobbies: ", hobbies)
	fmt.Println("First Hobby: ", hobbies[0], ",", "Second and Third Hobby: ", hobbies[1], " and ", hobbies[2])

	// Slice based on the 1st Array
	hobbiesDivided := hobbies[0:1]
	hobbiesDivided2 := []string{hobbies[0], hobbies[1]}
	fmt.Println(hobbiesDivided, hobbiesDivided2)


	hobbiesDivided[0] = hobbiesDivided[1]
	hobbiesDivided[0] = hobbies[2]
	fmt.Println("Updated Hobbies: ", hobbiesDivided)

	// Dynamic Array
	course_goals := []string{"Learn Go", "Build Projects with Go", "Later on dive deep into Go Ecosystem"}
	fmt.Println("Course Goals: ", course_goals)
	course_goals[1] = "Build Projects with Go and Learn Kubernetes"
	course_goals = append(course_goals, "Learn to build microservices with Go")
	fmt.Println("Updated Course Goals: ", course_goals)


	// Dynamic Array of Product Struct Type
	products := []Product{{1, "Laptop", 1000.45}, {2, "Mobile", 500.23}}
	fmt.Println("Products: ", products)
}