package client

import "testing"

type TestCleanInputTestCase struct {
	input    string
	expected []string
}

func TestCleanInput(t *testing.T) {
	cases := []TestCleanInputTestCase{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "    helloworld  ",
			expected: []string{"helloworld"},
		},
		{
			input:    "my name is meoth",
			expected: []string{"my", "name", "is", "meoth"},
		},
		{
			input:    "MY NaMe is meoth",
			expected: []string{"my", "name", "is", "meoth"},
		},
	}

	for _, test := range cases {
		actual := cleanInput(test.input)
		expected := test.expected

		if len(actual) != len(expected) {
			t.Errorf("expected length: %d == actual length: %d", len(expected), len(actual))
		}

		notMatched := false
		for i := range actual {
			if actual[i] != expected[i] {
				notMatched = true
				break
			}
		}
		if notMatched {
			t.Errorf("Expected : %v \n Got: %v", expected, actual)
		}
	}

}
