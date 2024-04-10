package main

import "fmt"

func main() {
	dir := make(map[string][]string)

	dir["home"] = []string{"lib", "log", "Desktop", "Downloads"}

	fmt.Println(dir["home"])
}
