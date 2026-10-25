package main

import "fmt"

func StudentScores() map[string]int {
	var scores = make(map[string]int)
	scores["Анна"] = 3
	scores["Борис"] = 0
	scores["Вера"] = 5

	scores["Анна"] = scores["Анна"] + 2
	delete(scores, "Вера")
	return scores
}

func main() {
	var studentScores = StudentScores()
	fmt.Println(studentScores["Анна"])
	borisScore, ok := studentScores["Борис"]
	glebScore, ok := studentScores["Глеб"]
	fmt.Println(borisScore, ok)
	fmt.Println(glebScore, ok)
	fmt.Println(len(studentScores))

}
