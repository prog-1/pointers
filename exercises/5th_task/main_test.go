package main

import (
	"reflect"
	"testing"
)

func TestPrepend(t *testing.T) {
	for _, tc := range []struct {
		input  []int
		result []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1}, []int{1, 3}},
	} {
		t.Run("", func(t *testing.T) {
			n := &Node{tc.input[0], nil}
			for _, v := range tc.input[1:] {
				Prepend(&n, v)
			}
			if !reflect.DeepEqual(n.asSlice(), tc.result) {
				t.Errorf("Input = %v Output = %v Want = %v", tc.input, n.asSlice(), tc.result)
			}
		})

	}
}

func TestAppend(t *testing.T) {
	for _, tc := range []struct {
		input []int
	}{
		{[]int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}},
	} {
		t.Run("", func(t *testing.T) {
			n := &Node{tc.input[0], nil}
			for _, v := range tc.input[1:] {
				Append(&n, v)
			}
			if !reflect.DeepEqual(n.asSlice(), tc.input) {
				t.Errorf("Input = %v Output = %v Want = %v", tc.input, n.asSlice(), tc.input)
			}
		})

	}
}

func TestInsert(t *testing.T) {
	for _, tc := range []struct {
		initialValues []int
		input         int
		output        []int
	}{
		{[]int{1, 2, 4, 5}, 3, []int{1, 2, 3, 4, 5}},
		{[]int{1}, 3, []int{1, 3}},
		{[]int{3}, 1, []int{3, 1}},

		//{[]int{5, 4, 3, 2, 1}},
	} {
		t.Run("", func(t *testing.T) {
			n := &Node{tc.initialValues[0], nil}
			for _, v := range tc.initialValues[1:] {
				Append(&n, v)
			}
			Insert(&n, tc.input)
			if !reflect.DeepEqual(n.asSlice(), tc.output) {
				t.Errorf("Insert(%v, %v) Output = %v Want = %v", tc.initialValues, tc.input, n.asSlice(), tc.output)
			}
		})

	}
}

func TestContains(t *testing.T) {
	for _, tc := range []struct {
		initialValues []int
		input         int
		output        bool
	}{
		{[]int{1, 2, 4, 5}, 3, false},
		{[]int{1, 2, 4, 5}, 4, true},
		{[]int{1}, 3, false},
		{[]int{3}, 3, true},

		//{[]int{5, 4, 3, 2, 1}},
	} {
		t.Run("", func(t *testing.T) {
			n := &Node{tc.initialValues[0], nil}
			for _, v := range tc.initialValues[1:] {
				Append(&n, v)
			}

			if got := Contains(n, tc.input); got != tc.output {
				t.Errorf("Insert(%v, %v) Output = %v Want = %v", tc.initialValues, tc.input, got, tc.output)
			}

		})

	}
}

func TestRemove(t *testing.T) {
	for _, tc := range []struct {
		initialValues []int
		input         int
		output        []int
		bOutput       bool
	}{
		{[]int{1, 2, 4, 5}, 1, []int{2, 4, 5}, true},
		{[]int{1, 2, 4, 5}, 3, []int{1, 2, 4, 5}, false},
		{[]int{1, 2, 4, 5}, 5, []int{1, 2, 4}, true},

		//{[]int{5, 4, 3, 2, 1}},
	} {
		t.Run("", func(t *testing.T) {
			n := &Node{tc.initialValues[0], nil}
			for _, v := range tc.initialValues[1:] {
				Append(&n, v)
			}

			if got := Remove(&n, tc.input); got != tc.bOutput || !reflect.DeepEqual(tc.output, n.asSlice()) {
				t.Errorf("Remove(%v, %v) = %v, %v Want = %v, %v", tc.initialValues, tc.input, got, n.asSlice(), tc.bOutput, tc.output)
			}

		})

	}
}
