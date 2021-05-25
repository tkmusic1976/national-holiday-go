package main

import (
	"fmt"

	holiday "github.com/tkmusic1976/national-holiday-go"
)

func main() {
	entries, err := holiday.AllEntries()
	if err != nil {
		panic(err)
	}
	for _, e := range entries {
		fmt.Printf("%s = %s\n", e.YMD, e.Name)
	}
}
