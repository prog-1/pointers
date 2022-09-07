package main

import "fmt"

func UpdateArray(Array [5]int, index int, value int) {
	Array[index] = value
}
func UpdateWithAdress(Array *[5]int, index int, value int) {
	Array[index] = value
}

func main1() {
	numberArray := [5]int{3, 56, 2, 5, 7}
	UpdateArray(numberArray, 1, 6)
	fmt.Println("Array: ", numberArray) //no updates in array

	UpdateWithAdress(&numberArray, 1, 6)
	fmt.Println("Address on Array: ", numberArray) //array is updated
}

//Conclusion: when the function accepts just an array, the array is not being updated. When you paste the adress on the array with the pointer, the array updates normally. Thus I can judge that arrays are values, not references.
