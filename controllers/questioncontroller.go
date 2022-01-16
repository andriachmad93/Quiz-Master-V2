/*
Package Controllers implements question functions to make
- create
- list
- detail
- update
- delete
- answer
*/
package controllers

// import needs to use go library for question controller
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"quiz_master_v2/models"
	"quiz_master_v2/utils"
)

var questions []models.Question

// List function shows questions
func List() {
	utils.Block{
		Try: func() {
			// check questions array data more than 0
			if len(questions) > 0 {
				for _, question := range questions {
					fmt.Printf("No. %d\n", question.ID)
					fmt.Printf("Q: %s", question.Q)
					fmt.Printf("A: %s\n", question.Answer)
				}
			} else {
				fmt.Println("No Question Data")
			}
		},
		Catch: func(e utils.Exception) {
			fmt.Printf("Something Wrong: %v\n", e)
		},
	}.Do()
}

// Detail function shows question detail by id
func Detail() {
	utils.Block{
		Try: func() {
			// make user input no
			fmt.Println("No")
			id_reader := bufio.NewReader(os.Stdin)
			id, _ := id_reader.ReadString('\n')
			id = strings.ReplaceAll(id, "\n", "")

			// convert id string to id int with var no
			no, err := strconv.Atoi(id)

			if err != nil {
				utils.Throw(err)
			}

			// get question detail
			detail := models.Get(no, questions)

			if detail.ID == no {
				fmt.Printf("Q: %s", detail.Q)
				fmt.Printf("A: %s\n", detail.Answer)
			} else {
				fmt.Println("No Question found")
			}
		},
		Catch: func(e utils.Exception) {
			fmt.Println("No Question found")
			fmt.Printf("Something Wrong: %v\n", e)
		},
	}.Do()
}

// Create function creates question per item
func Create() {
	utils.Block{
		Try: func() {
			// make user input no, question, answer
			fmt.Println("No")
			id_reader := bufio.NewReader(os.Stdin)
			id, _ := id_reader.ReadString('\n')
			id = strings.ReplaceAll(id, "\n", "")

			// convert id string to id int with var no
			no, err := strconv.Atoi(id)

			if err != nil {
				utils.Throw(err)
			}

			fmt.Println("Question")
			q_reader := bufio.NewReader(os.Stdin)
			q_read, _ := q_reader.ReadString('\n')

			fmt.Println("Answer")
			answer_reader := bufio.NewReader(os.Stdin)
			answer, _ := answer_reader.ReadString('\n')

			q := models.Question{
				ID:     no,
				Q:      q_read,
				Answer: answer,
			}

			// join question object to questions array
			questions = append(questions, q)

			fmt.Printf("\nQuestion No %d created\n", q.ID)
			fmt.Printf("Q: %s", q.Q)
			fmt.Printf("A: %s", q.Answer)
		},
		Catch: func(e utils.Exception) {
			fmt.Println("No Question found")
			fmt.Printf("Something Wrong: %v\n", e)
		},
	}.Do()
}

// Update function updates question by id with deleting question object
func Update() {
	utils.Block{
		Try: func() {
			// make user input no, question, answer
			fmt.Println("No")
			id_reader := bufio.NewReader(os.Stdin)
			id, _ := id_reader.ReadString('\n')
			id = strings.ReplaceAll(id, "\n", "")

			// convert id string to id int with var no
			no, err := strconv.Atoi(id)

			if err != nil {
				utils.Throw(err)
			}

			fmt.Println("Question")
			q_reader := bufio.NewReader(os.Stdin)
			q_read, _ := q_reader.ReadString('\n')

			fmt.Println("Answer")
			answer_reader := bufio.NewReader(os.Stdin)
			answer, _ := answer_reader.ReadString('\n')

			// get question detail
			detail := models.Get(no, questions)

			if detail.ID == no {
				q := models.Question{
					ID:     no,
					Q:      q_read,
					Answer: answer,
				}

				models.Update(no, q, questions)

				fmt.Printf("\nQuestion No %d updated\n", q.ID)
				fmt.Printf("Q: %s", q.Q)
				fmt.Printf("A: %s", q.Answer)
			} else {
				fmt.Println("No Question found")
			}
		},
		Catch: func(e utils.Exception) {
			fmt.Println("No Question found")
			fmt.Printf("Something Wrong: %v\n", e)
		},
	}.Do()
}

// Delete function deletes question by id
func Delete() {
	utils.Block{
		Try: func() {
			// make user input no
			fmt.Println("No")
			id_reader := bufio.NewReader(os.Stdin)
			id, _ := id_reader.ReadString('\n')
			id = strings.ReplaceAll(id, "\n", "")

			// convert id string to id int with var no
			no, err := strconv.Atoi(id)

			if err != nil {
				utils.Throw(err)
			}

			// get question detail
			detail := models.Get(no, questions)

			if detail.ID == no {
				models.Delete(detail, questions)

				fmt.Printf("Question %d was deleted\n", no)
			} else {
				fmt.Println("No Question found")
			}
		},
		Catch: func(e utils.Exception) {
			fmt.Println("No Question found")
			fmt.Printf("Something Wrong: %v\n", e)
		},
	}.Do()
}

// Answer function show questions and answer equals correct or incorrect
func Answer() {
	utils.Block{
		Try: func() {
			// make user input no and answer
			fmt.Println("No")
			id_reader := bufio.NewReader(os.Stdin)
			id, _ := id_reader.ReadString('\n')
			id = strings.ReplaceAll(id, "\n", "")

			fmt.Println("Answer")
			answer_reader := bufio.NewReader(os.Stdin)
			answer, _ := answer_reader.ReadString('\n')

			// convert id string to id int with var no
			no, err := strconv.Atoi(id)

			if err != nil {
				utils.Throw(err)
			}

			// get question detail
			detail := models.Get(no, questions)

			if detail.ID == no {
				// check if answer correct or incorrect
				if detail.Answer == answer {
					fmt.Println("Correct")
				} else {
					fmt.Println("Incorrect")
				}
			} else {
				fmt.Println("No Question found")
			}
		},
		Catch: func(e utils.Exception) {
			fmt.Println("No Question found")
			fmt.Printf("Something Wrong: %v\n", e)
		},
	}.Do()
}
