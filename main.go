package main

import (
	"fmt"

	"quiz_master_v2/controllers"
)

func init() {
	fmt.Println("Welcome to Quiz Master")
}

func main() {
	var command string

	fmt.Println("Please Input Your Command : ")
	fmt.Scanf("%s", &command)

	switch command {
	case "create_question":
		controllers.Create()
		main()
	case "questions":
		controllers.List()
		main()
	case "question":
		controllers.Detail()
		main()
	case "update_question":
		controllers.Update()
		main()
	case "delete_question":
		controllers.Delete()
		main()
	case "answer_question":
		controllers.Answer()
		main()
	case "help":
		fmt.Println("Command | Description")
		fmt.Println("create_question | Creates a question")
		fmt.Println("update_question | Updates a question")
		fmt.Println("delete_question | Deletes a question")
		fmt.Println("question | Shows a question")
		fmt.Println("questions | Shows question list")
		fmt.Println("answer_question | Answers a question")

		main()
	case "exit":
		fmt.Println("You Quit from Quiz Master")
	default:
		fmt.Println("Invalid Command")
		main()
	}
}
