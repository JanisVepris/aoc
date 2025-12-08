// Package dsu implements a Disjoint Set Union (Union-Find) data structure.
package dsu

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)

	for i := range n {
		parent[i] = i
		size[i] = 1
	}

	return &DSU{
		parent: parent,
		size:   size,
	}
}

// Find returns the representative (root) of the set that contains element x.
func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x]) // path compression
	}
	return d.parent[x]
}

func (d *DSU) Union(a, b int) {
	ra := d.Find(a)
	rb := d.Find(b)
	if ra == rb {
		return
	}

	// attach smaller to larger
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}

	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
}

func (d *DSU) IsRoot(a int) bool {
	return d.parent[a] == a
}

func (d *DSU) GetSize(a int) int {
	return d.size[a]
}

func (d *DSU) GetRootSize(a int) int {
	return d.size[d.Find(a)]
}

func (d *DSU) GetComponentCount() int {
	counts := map[int]bool{}
	for i := range d.parent {
		root := d.Find(i)
		counts[root] = true
	}
	return len(counts)
}

func (d *DSU) GetRoots() []int {
	roots := []int{}
	for i := range d.parent {
		if d.IsRoot(i) {
			roots = append(roots, i)
		}
	}
	return roots
}
