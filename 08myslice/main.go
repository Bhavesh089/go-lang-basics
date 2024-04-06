package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("This is the myslice")

	var fruitList = []string{}

	fmt.Printf("Type of the fruit %T\n", fruitList)

	fruitList = append(fruitList, "orange", "banana", "grapes", "mangoes")

	fmt.Println("This is fruitsLis", fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println("This is fruitsLis after slicing", fruitList)

	scores := make([]int, 5)

	scores[0] = 1
	scores[1] = 2
	scores[2] = 3
	scores = append(scores, 3, 4, 5, 6, 7, 8, 9)
	scores[3] = 3

	fmt.Println("This is the scores", scores)

	sort.Ints(scores)
	fmt.Println("This is sorted in scores", scores)

	//remove of index from the slices

	var courses = []string{"python", "swift", "golang", "javascript"}

	fmt.Println("This is the courses", courses)

	removeIndex := 2
	courses = append(courses[:removeIndex], courses[removeIndex+1:]...)

	fmt.Println("This is after removing the index", courses)

}
