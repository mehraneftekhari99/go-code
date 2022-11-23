// useful functions for slices

// [x] Index/IndexOf    | index of element by direct comparison / by predicate, -1 if not found
// [x] Find             | first element that satisfies predicate, nil if not found
// [x] Contains         | true if found by direct comparison
// [x] ForAny           | true if found by predicate
// [x] ForAll           | true if all satisfy predicate
// [x] Filter/FilterNot | a copy with only elements that satisfy/not predicate
// [ ] Unique           | a copy with only unique elements
// [ ] Join             | a string with elements joined with a separator between them
// [ ] Map              | a copy with f applied to elements
// [ ] FoldL/FoldR      | combine all elements sequentially starting with a z, starting with z and going left/right
// [ ] ScanL/ScanR      | combine adjacent elements into a new sequence, starting with z and going left/right

package slices

func Index[T comparable](L []T, I T) int {
	for i := range L {
		if L[i] == I {
			return i
		}
	}
	return -1
}

func IndexOf[T comparable](L []T, f func(T) bool) int {
	for i := range L {
		if f(L[i]) {
			return i
		}
	}
	return -1
}

func Find[T comparable](L []T, I T) T {
	return L[Index(L, I)]
}

func Contains[T comparable](L []T, I T) bool {
	return Index(L, I) > -1
}

func ForAny[T comparable](L []T, f func(T) bool) bool {
	return IndexOf(L, f) > -1
}

func ForAll[T comparable](L []T, f func(T) bool) bool {
	for i := range L {
		if !f(L[i]) {
			return false
		}
	}
	return true
}

func Filter[T comparable](L1 []T, f func(T) bool) []T {
	var L2 []T
	for i := range L1 {
		if f(L1[i]) {
			L2 = append(L2, L1[i])
		}
	}
	return L2
}
