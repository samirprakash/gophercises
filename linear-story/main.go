package main

import (
	"bufio"
	"fmt"
	"os"
)

type storyPage struct {
	text     string
	nextPage *storyPage
}

func playStory(page *storyPage) {
	scanner := bufio.NewScanner(os.Stdin)

	if page == nil {
		fmt.Println("End of Story!")
		return
	}
	fmt.Println(page.text)
	scanner.Scan()
	playStory(page.nextPage)
}

func main() {
	page1 := storyPage{text: "It is a dark and stormy night!", nextPage: nil}
	page2 := storyPage{text: "You move ahead and try to find the magic helmet before the bad guys find it!", nextPage: nil}
	page3 := storyPage{text: "You see a troll ahead in the dark!", nextPage: nil}
	page1.nextPage = &page2
	page2.nextPage = &page3

	playStory(&page1)
}
