// // utils/utils.go

package utils

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"golang.org/x/exp/constraints"
)

// This package contains functions that could be used in multiple days

// This function checks if an element belongs to a slice
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// This function checks the minimum between 2 values
func Min[T constraints.Ordered](e1, e2 T) T {
	if e1 < e2 {
		return e1
	} else {
		return e2
	}
}

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

// Eliminate duplicates in a slice
func UniqueConv[T comparable](s []T) []T {
	inResult := make(map[T]bool)
	var result []T
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func Map[T, F any](vals []T, f func(T) F) []F {
	res := make([]F, len(vals))
	for i, val := range vals {
		res[i] = f(val)
	}
	return res
}

// Queue implementation

type Queue[T any] []T

func (q *Queue[T]) Put(x T) {
	*q = append(*q, x)
}

func (q *Queue[T]) Get() T {
	ret := (*q)[0]
	*q = (*q)[1:]

	return ret
}

func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}

func Make2D[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		endRow := startRow + m
		matrix[i] = rows[startRow:endRow:endRow]
	}
	return matrix
}

// Integer trees, mamma mia
type Tree struct {
	Val  int
	Tree *[]Tree
	Deep int
}

func ParseTree(input *string) Tree {
	pointer := 0
	return ParseTreeAux(input, 0, &pointer)
}

// Parse trees from: "[[1,2],3,[],[[3],2]]"
func ParseTreeAux(input *string, deep int, p *int) Tree {
	//fmt.Println(*p)
	val := string((*input)[*p])
	var treeSol Tree
	treeSol.Deep = deep
	//fmt.Println(val)
	//fmt.Println(deep)

	// If the token is ",", we skip it
	if val == "," {
		*p++
	}

	switch val {
	// If the token is "[" we want to start another treeParse
	case "[":
		treeSol.Val = -1
		treeSol.Tree = &[]Tree{}
		if string((*input)[*p+1]) != "]" {
			for string((*input)[*p]) != "]" {
				*p++
				*treeSol.Tree = append(*treeSol.Tree, ParseTreeAux(input, deep+1, p))
			}
		} else {
			*p++
		}
		// This should be "]", so skip it
		*p++

	// If it goes here, then it is a number
	default:
		//Check if next is a number
		if _, err := strconv.Atoi(string((*input)[*p+1])); err == nil {
			*p++
			val = val + string((*input)[*p])
		}
		//fmt.Println("ok")
		num, _ := strconv.Atoi(val)
		treeSol.Val = num
		treeSol.Tree = nil
		*p++
	}
	return treeSol

}

func (t *Tree) Show() string {
	if t.Tree == nil {
		return fmt.Sprintf("%d", t.Val)
	}

	str := "["
	if len(*t.Tree) != 0 {
		for _, tree := range *t.Tree {
			str += tree.Show() + ","
		}
		str = str[:len(str)-1]
		str += "]"
	} else {
		str = "[]"
	}
	return str
}

func (t *Tree) IsVal() bool {
	return t.Tree == nil
}

func (t *Tree) IsLeaf() bool {
	return t.Tree != nil
}

func (t *Tree) IsDividerPacket() bool {
	aux1 := *t.Tree
	if len(aux1) == 1 {
		aux2 := *aux1[0].Tree
		if len(aux2) == 1 {
			return aux2[0].Val == 2 || aux2[0].Val == 6
		}
	}
	return false
}

// Transpose Matrix
func Transpose[T any](a [][]T) [][]T {
	newArr := make([][]T, len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			newArr[j] = append(newArr[j], a[i][j])
		}
	}
	return newArr
}

//INTervals

type Interval struct {
	L int
	R int
}

func (i Interval) Show() string {
	return fmt.Sprintf("(%d,%d)", i.L, i.R)
}

func SweepLine(intervals []Interval) []Interval {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i].L < intervals[j].L })
	//fmt.Println(intervals)

	cur := intervals[0]
	//fmt.Println(cur, intervals[1:])
	sol := []Interval{}
	for _, x := range intervals[1:] {
		//fmt.Println(cur, x)
		if cur.R < x.L {
			sol = append(sol, cur)
			cur = x
		} else if cur.R < x.R {
			cur.R = x.R
		}
	}
	sol = append(sol, cur)

	return sol
}

func (i Interval) Len() int {
	return i.R - i.L + 1
}
