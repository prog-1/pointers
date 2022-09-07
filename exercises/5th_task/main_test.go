package main

import (
	"reflect"
	"testing"
)

func TestPrepend(t *testing.T) {
	for _, tc := range []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1}, []int{1, 3}},
		{nil, []int{0}},
	} {
		t.Run("", func(t *testing.T) {
			var n *Node
			if tc.input != nil {
				n = &Node{tc.input[0], nil}
				for _, v := range tc.input[1:] {
					Prepend(&n, v)
				}
			} else {
				Prepend(&n, 0)
			}
			if !reflect.DeepEqual(n.asSlice(), tc.want) {
				t.Errorf("Input = %v Output = %v Want = %v", tc.input, n.asSlice(), tc.want)
			}
		})

	}
}

func TestAppend(t *testing.T) {
	for _, tc := range []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{5, 4, 3, 2, 1}},
		{nil, []int{0}},
	} {
		t.Run("", func(t *testing.T) {
			var n *Node
			if tc.input != nil {
				n = &Node{tc.input[0], nil}
				for _, v := range tc.input[1:] {
					Append(&n, v)
				}
			} else {
				Append(&n, 0)
			}
			if !reflect.DeepEqual(n.asSlice(), tc.want) {
				t.Errorf("Input = %v Output = %v Want = %v", tc.input, n.asSlice(), tc.want)
			}
		})

	}
}

func TestInsert(t *testing.T) {
	for _, tc := range []struct {
		initialValues []int
		input         int
		want          []int
	}{
		{[]int{1, 2, 4, 5}, 3, []int{1, 2, 3, 4, 5}},
		{[]int{1}, 3, []int{1, 3}},
		{[]int{3}, 1, []int{1, 3}},
		{nil, 0, []int{0}},

		//{[]int{5, 4, 3, 2, 1}},
	} {
		t.Run("", func(t *testing.T) {
			var n *Node
			if tc.initialValues != nil {
				n := &Node{tc.initialValues[0], nil}
				for _, v := range tc.initialValues[1:] {
					Append(&n, v)
				}
			}
			Insert(&n, tc.input)
			if !reflect.DeepEqual(n.asSlice(), tc.want) {
				t.Errorf("Insert(%v, %v) Output = %v Want = %v", tc.initialValues, tc.input, n.asSlice(), tc.want)
			}
		})

	}
}

func TestContains(t *testing.T) {
	for _, tc := range []struct {
		initialValues []int
		input         int
		want          bool
	}{
		{[]int{1, 2, 4, 5}, 3, false},
		{[]int{1, 2, 4, 5}, 4, true},
		{[]int{1}, 3, false},
		{[]int{3}, 3, true},
		{nil, 0, false},

		//{[]int{5, 4, 3, 2, 1}},
	} {
		t.Run("", func(t *testing.T) {
			var n *Node
			if tc.initialValues != nil {
				n = &Node{tc.initialValues[0], nil}
				for _, v := range tc.initialValues[1:] {
					Append(&n, v)
				}
			}

			if got := Contains(n, tc.input); got != tc.want {
				t.Errorf("Insert(%v, %v) Output = %v Want = %v", tc.initialValues, tc.input, got, tc.want)
			}

		})

	}
}

func TestRemove(t *testing.T) {
	for _, tc := range []struct {
		initialValues []int
		input         int
		want          []int
		bWant         bool
	}{
		{[]int{1, 2, 4, 5}, 1, []int{2, 4, 5}, true},
		{[]int{1, 2, 4, 5}, 3, []int{1, 2, 4, 5}, false},
		{[]int{1, 2, 4, 5}, 5, []int{1, 2, 4}, true},
		{nil, 1, nil, false},
		//{[]int{5, 4, 3, 2, 1}},
	} {
		t.Run("", func(t *testing.T) {
			var n *Node
			if tc.initialValues != nil {
				n = &Node{tc.initialValues[0], nil}
				for _, v := range tc.initialValues[1:] {
					Append(&n, v)
				}
			}
			if got := Remove(&n, tc.input); got != tc.bWant || !reflect.DeepEqual(tc.want, n.asSlice()) {
				t.Errorf("Remove(%v, %v) = %v, %v Want = %v, %v", tc.initialValues, tc.input, got, n.asSlice(), tc.bWant, tc.want)
			}

		})

	}
}
