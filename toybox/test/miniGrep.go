package main

import(
	"fmt"
	"os"
	"strings"
)

func findAndPrint(term string, lines []string) bool {
	ret := false
	for _,line := range lines {
		if strings.Contains(line, term) {
			fmt.Println(line)
			ret = true
		}
	}
	return ret
}

func main() {
	term := os.Args[1]
	file := os.Args[2]
	data, err := os.ReadFile((file))
	if err != nil {
		fmt.Println("Error opening file")
		return
	} 
	fmt.Println("Search term: " + term)
	lines := strings.Split(string(data), "\n")
	found := findAndPrint(term, lines)
	if !found {
		fmt.Println("No match found")
	}
}
