package main

import (
	"fmt"
)

type storyPage struct {
	text     string
	nextPage *storyPage
}

func (page *storyPage) playStory() {
	for page != nil {
		fmt.Println(page.text)
		page = page.nextPage
	}
}

// Add nodes to the end of the linked list
func (page *storyPage) addToEnd(text string) {
	newPage := storyPage{text: text, nextPage: nil}
	for page.nextPage != nil {
		page = page.nextPage
	}
	page.nextPage = &newPage
}

// Add nodes in  between the linked list
func (page *storyPage) addAfter(text string) {
	newPage := storyPage{text: text, nextPage: page.nextPage}
	page.nextPage = &newPage
}

// Deleted nodes from the linked list
func (page *storyPage) delete() {

}

func main() {
	page := storyPage{text: "It is a dark and stormy night!", nextPage: nil}
	page.addToEnd("You move ahead and try to find the magic helmet before the bad guys find it!")
	page.addToEnd("You see a troll ahead in the dark!")

	page.addAfter("Testing add after")

	page.playStory()
}
