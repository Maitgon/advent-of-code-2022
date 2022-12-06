// // utils/utils.go

package utils

import (
	"fmt"
	"log"
)

// This package contains functions that could be used in multiple days

// Intersects two generic slices using hash table
func Intersection[T comparable](s1, s2 []T) (inter []T) {
	hash := make(map[T]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		// If elements present in the hashmap then append intersection list.
		if hash[e] {
			inter = append(inter, e)
		}
	}
	return
}

// Intersects 3 generic slices
func Intersection3[T comparable](s1, s2, s3 []T) []T {
	return Intersection(Intersection(s1, s2), s3)
}

// Implementation of Stack
type Stack[T any] []T

func (s *Stack[T]) Push(v T) int {
	*s = append(*s, v)
	return len(*s)
}

func (s *Stack[T]) PushN(v []T) int {
	*s = append(*s, v...)
	return len(*s)
}

func (s *Stack[T]) Last() T {
	l := len(*s)

	// Upto the developer to handle an empty stack
	if l == 0 {
		log.Fatal("Empty Stack")
	}

	last := (*s)[l-1]
	return last
}

func (s *Stack[T]) Pop() T {
	removed := (*s).Last()
	*s = (*s)[:len(*s)-1]

	return removed
}

func (s *Stack[T]) PopN(n int) []T {
	removed := make([]T, n)
	copy(removed, (*s)[len(*s)-n:])
	*s = (*s)[:len(*s)-n]

	return removed
}

// Pointer not needed because read-only operation
func (s Stack[T]) Values() {
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Printf("%v ", s[i])
	}
}

// Checks if a slice is unique using hash tables
func Unique[T comparable](s []T) bool {
	set := map[T]bool{}
	for _, c := range s {
		set[c] = true
	}
	return len(set) == len(s)
}
