package main

import (
	"bufio"
	"fmt"
	"os"
)

type storyNode struct {
	text    string
	yesPath *storyNode
	noPath  *storyNode
}

func (s *storyNode) play() {
	fmt.Println(s.text)

	scanner := bufio.NewScanner(os.Stdin)

	if s.yesPath != nil && s.noPath != nil {
		for {
			fmt.Print("Choice (yes/no) => ")
			scanner.Scan()
			c := scanner.Text()

			if c == "yes" {
				s.yesPath.play()
				break
			} else if c == "no" {
				s.noPath.play()
				break
			} else {
				fmt.Println("This is not a valid response. Say 'yes' or 'no'")
			}
		}
	}
}

func (s *storyNode) print() {
	fmt.Println(s.text)
	if s.yesPath != nil {
		s.yesPath.print()
	}
	if s.noPath != nil {
		s.noPath.print()
	}
}

func main() {

	r := storyNode{text: "You are at the entrance of a cave. Do you want to go in the cave?", yesPath: nil, noPath: nil}
	w := storyNode{text: "You have won!", yesPath: nil, noPath: nil}
	l := storyNode{text: "You have lost!", yesPath: nil, noPath: nil}
	r.yesPath = &l
	r.noPath = &w
	r.play()
	r.print()
}
