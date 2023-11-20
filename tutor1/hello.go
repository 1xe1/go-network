package main

import "fmt"

func main() {
	// ตัวแปรชนิด int
	var age int
	age = 25
	fmt.Println("Age:", age)

	// ตัวแปรชนิด float64
	var height float64
	height = 175.5
	fmt.Println("Height:", height)

	// ตัวแปรชนิด string
	var name string
	name = "surasak"
	fmt.Println("Name:", name)

	// ใช้ := สำหรับการประกาศและกำหนดค่าในทำเลทีฟ (implicit)
	weight := 68.5
	isStudent := true
	fmt.Println("Weight:", weight)
	fmt.Println("Is Student:", isStudent)

	// ตัวแปรชนิด complex
	var complexNum complex128
	complexNum = 3 + 4i
	fmt.Println("Complex Number:", complexNum)

	// ตัวแปรชนิด array
	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Println("Numbers Array:", numbers)

	// ตัวแปรชนิด slice
	sliceNumbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice of Numbers:", sliceNumbers)

	// ตัวแปรชนิด map
	person := map[string]interface{}{
		"name":   "Alice",
		"age":    25,
		"height": 175.0,
	}
	fmt.Println("Person Map:", person)
}