package main

import "fmt"

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	// Calculate the number of pages needed to display all items
	pageSize := 4
	numPages := len(items) / pageSize
	if len(items)%pageSize != 0 {
		numPages++
	}

	// Loop through all pages
	for pageNum := 0; pageNum < numPages; pageNum++ {
		// Calculate the start and end index of the current page
		startIndex := pageNum * pageSize
		endIndex := startIndex + pageSize
		if endIndex > len(items) {
			endIndex = len(items)
		}

		// Get the items for the current page
		currentPageItems := items[startIndex:endIndex]

		// Print the page header and items
		pageHeader := fmt.Sprintf("Page %d", pageNum+1)
		fmt.Println(pageHeader, currentPageItems)
		// fmt.Print(startIndex)
		// fmt.Print(endIndex)
	}
}
