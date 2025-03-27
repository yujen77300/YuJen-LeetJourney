package groupanagrams

// import "testing"

// func TestGroupAnagrams(t *testing.T) {
// 	test := []struct {
// 		name     string
// 		in       []string
// 		expected [][]string
// 	}{
// 		{
// 			name:     "Example 1",
// 			in:       []string{"eat", "tea", "tan", "ate", "nat", "bat"},
// 			expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
// 		},
// 		{
// 			name:     "Example 2",
// 			in:       []string{""},
// 			expected: [][]string{{""}},
// 		},
// 		{
// 			name:     "Example 3",
// 			in:       []string{"a"},
// 			expected: [][]string{{"a"}},
// 		},
// 	}

// 	for _, tt := range test {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result := groupAnagrams(tt.in)
// 			if len(result) != len(tt.expected) {
// 				t.Fatalf("got: %v, want: %v", result, tt.expected)
// 			}
// 			for i, arr := range result {
// 				if len(arr) != len(tt.expected[i]) {
// 					t.Fatalf("got: %v, want: %v", result, tt.expected)
// 				}

// 				for j, val := range arr {
// 					if val != tt.expected[i][j] {
// 						t.Fatalf("got: %v, want: %v", result, tt.expected)
// 					}
// 				}
// 			}
// 		})
// 	}
// }
