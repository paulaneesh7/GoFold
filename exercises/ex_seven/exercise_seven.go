package exseven

func ExerciseSeven(num int) int {

	if num < 0 {
		return 0
	}

	val := 1
	for i := 1; i <= num; i++{
		val *= i
	}

	return val
}