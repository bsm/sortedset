package sortedset

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Set represents a slice of unique items.
type Set[T constraints.Ordered] []T

// New inits a new Set of size.
func New[T constraints.Ordered]() Set[T] {
	return NewCap[T](0)
}

// NewCap inits a new Set with a cap.
func NewCap[T constraints.Ordered](cap int) Set[T] {
	return make([]T, 0, cap)
}

// Len returns the set length.
func (s Set[T]) Len() int {
	return len(s)
}

// Clear returns a truncated set.
func (s Set[T]) Clear() Set[T] {
	return s[:0]
}

// Clone returns a copy of the set.
func (s Set[T]) Clone() Set[T] {
	return slices.Clone(s)
}

// Equal reports whether the sets s and t have the exact same elements.
func (s Set[T]) Equal(t Set[T]) bool {
	return slices.Equal(s, t)
}

// Add adds all v to Set s, and returns the new set.
func (s Set[T]) Add(v ...T) Set[T] {
	for _, w := range v {
		if pos, found := slices.BinarySearch(s, w); !found {
			s = slices.Insert(s, pos, w)
		}
	}
	return s
}

// Delete deletes all v from Set s and returns the new set.
func (s Set[T]) Delete(v ...T) Set[T] {
	for _, w := range v {
		if pos, found := slices.BinarySearch(s, w); found {
			s = slices.Delete(s, pos, pos+1)
		}
	}
	return s
}

// Has reports whether v is an element of Set s.
func (s Set[T]) Has(v T) bool {
	_, found := slices.BinarySearch(s, v)
	return found
}

// Intersection adds the intersection x ∩ y to s and returns the result.
func (s Set[T]) Intersection(x, y Set[T]) Set[T] {
	if len(x) < len(y) {
		x, y = y, x
	}

	var (
		offset int
		found  bool
	)

	for _, v := range x {
		if offset, found = binarySearchWithOffset(y, v, offset); found {
			s = s.Add(v)
		}
	}
	return s
}

// IntersectionWith sets s to the insersection of s ∪ t, and returns the result.
func (s Set[T]) IntersectionWith(t Set[T]) Set[T] {
	u := s.Clear()
	return u.Intersection(s, t)
}

// Intersects reports whether s ∩ t ≠ ∅.
func (s Set[T]) Intersects(t Set[T]) bool {
	sn, tn := len(s), len(t)
	if tn < sn {
		s, t = t, s
		sn, tn = tn, sn
	}
	if sn == 0 || s[0] > t[tn-1] || t[0] > s[sn-1] {
		return false
	}

	offset := 0
	for _, x := range s {
		if pos, found := binarySearchWithOffset(t, x, offset); found {
			return true
		} else if pos >= tn {
			return false
		} else {
			offset = pos
		}
	}
	return false
}

// Union adds the union x ∪ y to s.
func (s Set[T]) Union(x, y Set[T]) Set[T] {
	for _, v := range x {
		s = s.Add(v)
	}
	for _, v := range y {
		s = s.Add(v)
	}
	return s
}

// UnionWith sets s to the union s ∪ t, and returns the result.
func (s Set[T]) UnionWith(t Set[T]) Set[T] {
	return s.Union(nil, t)
}

// Slice returns the set as native slice.
func (s Set[T]) Slice() []T { return s }

func binarySearchWithOffset[T constraints.Ordered](s []T, x T, offset int) (int, bool) {
	pos, found := slices.BinarySearch(s[offset:], x)
	return pos + offset, found
}
