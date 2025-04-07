package longestrepeatingcharacterreplacement

import "testing"

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		s        string
		k        int
		expected int
	}{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
	}

	for _, test := range tests {
		result := characterReplacement(test.s, test.k)
		if result != test.expected {
			t.Errorf("characterReplacement(%s, %d) = %d; expected %d", test.s, test.k, result, test.expected)
		}
	}
}
