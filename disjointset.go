package disjointset

// DisjointSet is the interface for the disjoint-set (or union-find) data
// structure.
// Do not change the definition of this interface.
type DisjointSet interface {
	// UnionSet(set, t) merges (unions) the set containing set and t,
	// and returns the representative of the resulting merged set.
	UnionSet(int, int) int

	// FindSet(set) returns representative of the class that set belongs to.
	FindSet(int) int
}

type Set struct {
	parents [1000 * 1000]int
}

func (S *Set) FindSet(x int) int {
	if x == S.parents[x] {
		return x
	}
	S.parents[x] = S.FindSet(S.parents[x])
	return S.parents[x]
}

func (S *Set) UnionSet(x, y int) int {
	xroot := S.FindSet(x)
	yroot := S.FindSet(y)
	if xroot == yroot {
		return xroot
	}
	S.parents[yroot] = xroot
	return xroot
}

// NewDisjointSet creates a struct of a type that satisfies the DisjointSet interface.
func NewDisjointSet() DisjointSet {
	var S *Set
	S = new(Set)
	for i := 0; i < 1000*1000; i++ {
		S.parents[i] = i
	}
	return S
}
