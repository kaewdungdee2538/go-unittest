package services_test

import (
	"gotest/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckGrade(t *testing.T) {

	type testCasse struct {
		name     string
		score    int
		expected string
	}

	cases := []testCasse{
		{name: "A", score: 80, expected: "A"},
		{name: "B", score: 70, expected: "B"},
		{name: "C", score: 60, expected: "C"},
		{name: "D", score: 50, expected: "D"},
		{name: "F", score: 5, expected: "F"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			grade := services.CheckGrade(c.score)
			expected := c.expected
			// function for validate value
			assert.Equal(t, expected, grade)

			// if grade != expected {
			// 	t.Errorf("got %v expected %v", grade, expected)
			// }
		})
	}

}

func BenchmarkCheckGrade(b *testing.B) {
	for i := 0; i < b.N; i++ {
		services.CheckGrade(80)
	}
}