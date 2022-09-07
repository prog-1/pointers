package main

import "fmt"

func UpdateMap(Map map[string]int, number string, value int) {
	Map[number] = value
}

func main2() {
	justMap := make(map[string]int)
	justMap["n1"] = 42
	justMap["n2"] = 6
	justMap["n3"] = 12

	UpdateMap(justMap, "n2", 4)
	fmt.Println(justMap)
}

//Conclusion: this Update works without making any pointers, which means that map is the storage of references by itself.
