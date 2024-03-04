This project represents a disjoint set including merge and find features. It works off the assumption of $S= \\{\\dots\\{-2\\},\\{-1\\},\\{0\\},\\{1\\},\\{2\\},\\dots \\}$ . Each dynamic set in $S$, $S_i$, has a representative that remains the same no matter if $S_i$ has been modified. `FndSet(int)` returns the representative of the set containing `int`. `UnionSet(u,v)` merges the sets containing `u` and `v` resulting in $S_u \\cup S_v$, which gets added to set $S$ and sets $S_u$ and $S_v$ are removed. 

To run execute `go test`
