/*
Package Controllers implements question struct to add objects
- ID
- Q
- Answer
*/
package models

import "sort"

type Question struct {
	ID     int
	Q      string // Q = Question because name equals name struct
	Answer string
}

// Get function shows question detail by id
func Get(id int, questions []Question) Question {
	sort.Slice(questions, func(i, j int) bool {
		return questions[i].ID <= questions[j].ID
	})

	idx := sort.Search(len(questions), func(i int) bool {
		return questions[i].ID >= id
	})

	return questions[idx]
}

// Update function to update question detail by id
func Update(id int, question Question, questions []Question) Question {
	for i := 0; i < len(questions); i++ {
		attr := questions[i]

		if attr.ID == id {
			attr.Q = question.Q
			attr.Answer = question.Answer
		}
	}

	return question
}

func Delete(question Question, questions []Question) []Question {
	for idx, q := range questions {
		if q.ID == question.ID {
			return append(questions[0:idx], questions[idx+1:]...)
		}
	}

	return questions
}
