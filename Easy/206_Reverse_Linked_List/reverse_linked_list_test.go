package reverselinkedlist

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		in       *ListNode
		expected *ListNode
	}{
		{
			name: "Example 1",
			in: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			expected: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
		},
		{
			name: "Example 2",
			in: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			expected: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
		{
			name:     "Example 3",
			in:       nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseList(tt.in)
			if !linkedListsEqual(result, tt.expected) {
				t.Errorf("reverseList() = %v, want %v", linkedListToSlice(result), linkedListToSlice(tt.expected))
			}
		})
	}
}

// Helper function to compare two linked lists
func linkedListsEqual(l1, l2 *ListNode) bool {
	return reflect.DeepEqual(linkedListToSlice(l1), linkedListToSlice(l2))
}

// Helper function to convert a linked list to a slice for easier comparison
func linkedListToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
