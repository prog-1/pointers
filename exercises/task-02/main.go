package main

import "fmt"

func change(m map[string]string) { m["Checkers"] = "Strategy" }

func main() {
	game := map[string]string{
		"Checkers": "Shooter",
		"Chess":    "Strategy",
	}
	fmt.Println(game) //Output: map[Checkers:Shooter Chess:Strategy]
	change(game)
	fmt.Println(game) //Output: map[Checkers:Strategy Chess:Strategy]
}

//maps are references
