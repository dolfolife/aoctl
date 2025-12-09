# Day 08 Algorithms

This solution relies on two key data structures to efficiently solve the problem of connecting junction boxes: **EdgeHeap** and **UnionFind**.

## EdgeHeap (Min Heap)

**Why:** We need to process potential connections starting from the shortest distance. Sorting all edges ($O(E \log E)$) is acceptable, but a Min Heap allows us to efficiently pop the smallest element ($O(\log E)$) on demand.

**Example:**
Imagine edges with lengths `[10, 5, 8, 3]`.
A Min Heap organizes them effectively so `Pop()` always gives the smallest:
1. `Pop()` -> `3`
2. `Pop()` -> `5`
...

## UnionFind (Disjoint Set Union)

**Why:** We need to keep track of "circuits" (connected components) and know:
1. Are two points already connected? (To avoid redundant cycles in Part 1/2)
2. How many separate circuits remain? (To know when we have a single component in Part 2)

**Example:**
Points: `A`, `B`, `C`, `D`
- Start: `{A}`, `{B}`, `{C}`, `{D}` (4 components)
- Connect `A-B`: `{A,B}`, `{C}`, `{D}` (3 components)
- Connect `B-C`: Since `B` is in `{A,B}`, this merges `{A,B}` and `{C}` -> `{A,B,C}`, `{D}` (2 components)
