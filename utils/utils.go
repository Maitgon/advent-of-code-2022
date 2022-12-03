// // utils/utils.go

package utils

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
