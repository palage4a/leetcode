package leetcode

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestToSlice(t *testing.T) {
	a := &ListNode{1, &ListNode{3, nil}}
	expected := []int{1,3}
	actual := a.ToSlice()
	if cmp.Equal(actual, expected) != true {
		t.Fatalf("%v != %v", actual, expected)
	}
}

func TestDup(t *testing.T) {
	expected := &ListNode{1, &ListNode{3, &ListNode{5, nil}}}
	actual := expected.Dup()
	if &actual == &expected {
		t.Fatalf("%p == %p", actual, expected)
	}

	expectedSlice := []int{1,3,5}
	actualSlice := actual.ToSlice()
	if cmp.Equal(actualSlice, expectedSlice) != true {
		t.Fatalf("actual: %v != expected: %v", actualSlice, expectedSlice)
	}
}

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct{
		name string
		a *ListNode
		b *ListNode
		expected *ListNode
	}{
		{
			name: "default",
			a: &ListNode{1, &ListNode{3, nil}},
			b: &ListNode{2, &ListNode{4, nil}},
			expected: &ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						&ListNode{
							4,
							nil,
						},
					},
				},
			},
		},
		{
			name: "first_nil",
			a: nil,
			b: &ListNode{1, &ListNode{2, nil}},
			expected: &ListNode{
				1,
				&ListNode{
					2,
					nil,
				},
			},
		},
		{
			name: "second_nil",
			a: &ListNode{1, &ListNode{2, nil}},
			b: nil,
			expected: &ListNode{
				1,
				&ListNode{
					2,
					nil,
				},
			},
		},
		{
			name: "second",
			a: &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			b: &ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			expected: &ListNode{
				1,
				&ListNode{
					1,
					&ListNode{
						2,
						&ListNode{
							3,
							&ListNode{
								4,
								&ListNode{
									4,
									nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "both_nil",
			a: nil,
			b: nil,
			expected: nil,
		},
		{
			name: "third",
			a: &ListNode{-9, &ListNode{3, nil}},
			b: &ListNode{5, &ListNode{7, nil}},
			expected: &ListNode{
				-9,
				&ListNode{
					3,
					&ListNode{
						5,
						&ListNode{
							7,
							nil,
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a := tc.a
			b := tc.b
			expected := tc.expected
			actual := mergeTwoLists(a, b)

			actualSlice := actual.ToSlice()
			expectedSlice := expected.ToSlice()
			if cmp.Equal(actualSlice, expectedSlice) != true {
				t.Fatalf("%v isn't equal to %v", actualSlice, expectedSlice)
			}
		})
	}
}
