package printer

import "fmt"

const listIndentor = "- "

// List will be printed
func List(iterable []string) {
	for _, item := range iterable {
		fmt.Printf("%s%s\n", listIndentor, item)
	}
}

// Map keys will be printed
func Map(iterable map[string]string) {
	for _, item := range iterable {
		fmt.Printf("%s%s\n", listIndentor, item)
	}
}
