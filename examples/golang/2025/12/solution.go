package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Orientation struct {
	Cells []Point
	W     int
	H     int
}

type Shape struct {
	ID           int
	Area         int
	Orientations []Orientation
}

type Region struct {
	W      int
	H      int
	Counts []int
}

type pkey struct {
	shape int
	w     int
	h     int
}

func Part1(input string) (string, error) {
	shapes, regions, err := parseInput(input)
	if err != nil {
		return "", err
	}

	placementsCache := map[pkey][][]int{}

	fittable := 0
	for _, r := range regions {
		ok := canFitRegion(r.W, r.H, shapes, r.Counts, placementsCache)
		if ok {
			fittable++
		}
	}
	return strconv.Itoa(fittable), nil
}

func canFitRegion(
	W int,
	H int,
	shapes []Shape,
	counts []int,
	cache map[pkey][][]int,
) bool {
	if W <= 0 || H <= 0 {
		return false
	}

	pieceShape := make([]int, 0)
	totalArea := 0

	n := min(len(shapes), len(counts))
	for sid := 0; sid < n; sid++ {
		c := counts[sid]
		if c <= 0 {
			continue
		}
		if shapes[sid].Area == 0 {
			return false
		}
		for k := 0; k < c; k++ {
			pieceShape = append(pieceShape, sid)
		}
		totalArea += c * shapes[sid].Area
	}

	if len(pieceShape) == 0 {
		return true
	}
	if totalArea > W*H {
		return false
	}

	numPieces := len(pieceShape)
	numCells := W * H
	totalCols := numPieces + numCells

	cols := make([]*column, 0, totalCols)
	root := &column{name: "root"}
	root.L = root
	root.R = root

	primaryRemaining := 0
	for i := 0; i < totalCols; i++ {
		primary := i < numPieces
		c := newColumn(fmt.Sprintf("c%d", i), primary)
		cols = append(cols, c)
		insertColumn(root, c)
		if primary {
			primaryRemaining++
		}
	}

	d := &dlx{
		root:             root,
		primaryRemaining: primaryRemaining,
	}

	for pid := 0; pid < numPieces; pid++ {
		sid := pieceShape[pid]
		key := pkey{shape: sid, w: W, h: H}

		placements, ok := cache[key]
		if !ok {
			placements = generatePlacements(W, H, shapes[sid])
			cache[key] = placements
		}

		if len(placements) == 0 {
			return false
		}

		for _, cellIdxs := range placements {
			rowCols := make([]*column, 0, 1+len(cellIdxs))
			rowCols = append(rowCols, cols[pid])
			for _, cell := range cellIdxs {
				rowCols = append(rowCols, cols[numPieces+cell])
			}
			d.addRow(rowCols)
		}
	}

	return d.search()
}

type node struct {
	L   *node
	R   *node
	U   *node
	D   *node
	C   *column
	row int
}

type column struct {
	name    string
	primary bool
	S       int

	L *column
	R *column

	head node
}

func newColumn(name string, primary bool) *column {
	c := &column{name: name, primary: primary}
	c.head.C = c
	c.head.U = &c.head
	c.head.D = &c.head
	c.head.L = &c.head
	c.head.R = &c.head
	return c
}

func insertColumn(root, c *column) {
	c.R = root
	c.L = root.L
	root.L.R = c
	root.L = c
}

type dlx struct {
	root             *column
	primaryRemaining int
}

func (d *dlx) addRow(cols []*column) {
	var first *node
	var prev *node

	for _, c := range cols {
		n := &node{C: c}

		n.D = &c.head
		n.U = c.head.U
		c.head.U.D = n
		c.head.U = n
		c.S++

		if first == nil {
			first = n
			n.L = n
			n.R = n
		} else {
			n.L = prev
			n.R = first
			prev.R = n
			first.L = n
		}
		prev = n
	}
}

func (d *dlx) cover(c *column) {
	c.R.L = c.L
	c.L.R = c.R
	if c.primary {
		d.primaryRemaining--
	}

	for i := c.head.D; i != &c.head; i = i.D {
		for j := i.R; j != i; j = j.R {
			j.D.U = j.U
			j.U.D = j.D
			j.C.S--
		}
	}
}

func (d *dlx) uncover(c *column) {
	for i := c.head.U; i != &c.head; i = i.U {
		for j := i.L; j != i; j = j.L {
			j.C.S++
			j.D.U = j
			j.U.D = j
		}
	}

	// Restore column header in the horizontal list.
	c.R.L = c
	c.L.R = c
	if c.primary {
		d.primaryRemaining++
	}
}

func (d *dlx) choosePrimaryColumn() *column {
	var best *column
	bestSize := int(^uint(0) >> 1) // max int
	for c := d.root.R; c != d.root; c = c.R {
		if !c.primary {
			continue
		}
		if c.S < bestSize {
			bestSize = c.S
			best = c
			if bestSize == 0 {
				break
			}
		}
	}
	return best
}

func (d *dlx) search() bool {
	if d.primaryRemaining == 0 {
		return true
	}

	c := d.choosePrimaryColumn()
	if c == nil || c.S == 0 {
		return false
	}

	d.cover(c)
	for r := c.head.D; r != &c.head; r = r.D {
		for j := r.R; j != r; j = j.R {
			d.cover(j.C)
		}

		if d.search() {
			return true
		}

		for j := r.L; j != r; j = j.L {
			d.uncover(j.C)
		}
	}
	d.uncover(c)
	return false
}

func generatePlacements(W, H int, s Shape) [][]int {
	out := make([][]int, 0)

	for _, o := range s.Orientations {
		if o.W > W || o.H > H {
			continue
		}
		for y := 0; y <= H-o.H; y++ {
			for x := 0; x <= W-o.W; x++ {
				cells := make([]int, 0, len(o.Cells))
				for _, p := range o.Cells {
					idx := (y+p.Y)*W + (x + p.X)
					cells = append(cells, idx)
				}
				sort.Ints(cells)
				out = append(out, cells)
			}
		}
	}

	return out
}

func buildShape(id int, grid []string) (Shape, error) {
	if len(grid) == 0 {
		return Shape{}, fmt.Errorf("shape %d: empty grid", id)
	}
	w := len(grid[0])
	for _, row := range grid {
		if len(row) != w {
			return Shape{}, fmt.Errorf("shape %d: ragged rows", id)
		}
	}

	pts := make([]Point, 0)
	for y, row := range grid {
		for x := 0; x < len(row); x++ {
			if row[x] == '#' {
				pts = append(pts, Point{X: x, Y: y})
			}
		}
	}
	if len(pts) == 0 {
		return Shape{}, fmt.Errorf("shape %d: no occupied cells", id)
	}

	oris := uniqueOrientations(pts, w, len(grid))
	return Shape{
		ID:           id,
		Area:         len(pts),
		Orientations: oris,
	}, nil
}

func uniqueOrientations(points []Point, w, h int) []Orientation {
	seen := map[string]bool{}
	out := make([]Orientation, 0, 8)

	for t := 0; t < 8; t++ {
		cells, tw, th := transform(points, w, h, t)
		cells, tw, th = normalize(cells)

		key := encodeCells(cells)
		if seen[key] {
			continue
		}
		seen[key] = true

		out = append(out, Orientation{
			Cells: cells,
			W:     tw,
			H:     th,
		})
	}
	return out
}

func transform(points []Point, w, h, t int) ([]Point, int, int) {
	curW := w
	curH := h

	tmp := make([]Point, len(points))
	copy(tmp, points)

	if (t & 4) != 0 {
		for i := range tmp {
			tmp[i].X = curW - 1 - tmp[i].X
		}
	}

	rot := t & 3
	for r := 0; r < rot; r++ {
		for i := range tmp {
			x := tmp[i].X
			y := tmp[i].Y
			tmp[i].X = y
			tmp[i].Y = curW - 1 - x
		}
		curW, curH = curH, curW
	}

	return tmp, curW, curH
}

func normalize(points []Point) ([]Point, int, int) {
	minX := int(^uint(0) >> 1)
	minY := int(^uint(0) >> 1)
	maxX := -1
	maxY := -1

	for _, p := range points {
		if p.X < minX {
			minX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
	}

	out := make([]Point, len(points))
	for i, p := range points {
		x := p.X - minX
		y := p.Y - minY
		out[i] = Point{X: x, Y: y}
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	sort.Slice(out, func(i, j int) bool {
		if out[i].Y != out[j].Y {
			return out[i].Y < out[j].Y
		}
		return out[i].X < out[j].X
	})

	return out, maxX + 1, maxY + 1
}

func encodeCells(cells []Point) string {
	var b strings.Builder
	for i, p := range cells {
		if i > 0 {
			b.WriteByte(';')
		}
		b.WriteString(strconv.Itoa(p.X))
		b.WriteByte(',')
		b.WriteString(strconv.Itoa(p.Y))
	}
	return b.String()
}

func parseInput(input string) ([]Shape, []Region, error) {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	type rawShape struct {
		id   int
		grid []string
	}
	raw := make([]rawShape, 0)

	i := 0
	maxID := -1
	for i < len(lines) {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			i++
			continue
		}
		if isRegionLine(line) {
			break
		}
		if !strings.HasSuffix(line, ":") {
			return nil, nil, fmt.Errorf("expected shape header at line %d: %q", i+1, lines[i])
		}
		idStr := strings.TrimSuffix(line, ":")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, nil, fmt.Errorf("bad shape id at line %d: %q", i+1, lines[i])
		}
		i++

		grid := make([]string, 0)
		for i < len(lines) {
			row := strings.TrimSpace(lines[i])
			if row == "" || isRegionLine(row) {
				break
			}
			grid = append(grid, row)
			i++
		}
		raw = append(raw, rawShape{id: id, grid: grid})
		if id > maxID {
			maxID = id
		}
	}

	if maxID < 0 {
		return nil, nil, fmt.Errorf("no shapes found")
	}

	shapes := make([]Shape, maxID+1)
	seen := make([]bool, maxID+1)
	for _, rs := range raw {
		s, err := buildShape(rs.id, rs.grid)
		if err != nil {
			return nil, nil, err
		}
		shapes[rs.id] = s
		seen[rs.id] = true
	}
	for id := 0; id <= maxID; id++ {
		if !seen[id] {
			return nil, nil, fmt.Errorf("missing shape %d", id)
		}
	}

	regions := make([]Region, 0)
	for i < len(lines) {
		line := strings.TrimSpace(lines[i])
		i++
		if line == "" {
			continue
		}
		if !isRegionLine(line) {
			return nil, nil, fmt.Errorf("expected region line, got: %q", line)
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("bad region line: %q", line)
		}
		dims := strings.TrimSpace(parts[0])
		req := strings.TrimSpace(parts[1])

		xi := strings.IndexByte(dims, 'x')
		if xi < 0 {
			return nil, nil, fmt.Errorf("bad dims: %q", dims)
		}
		w, err := strconv.Atoi(dims[:xi])
		if err != nil {
			return nil, nil, fmt.Errorf("bad width in %q", dims)
		}
		h, err := strconv.Atoi(dims[xi+1:])
		if err != nil {
			return nil, nil, fmt.Errorf("bad height in %q", dims)
		}

		fields := strings.Fields(req)
		if len(fields) != len(shapes) {
			return nil, nil, fmt.Errorf(
				"region %dx%d: expected %d counts, got %d",
				w,
				h,
				len(shapes),
				len(fields),
			)
		}
		counts := make([]int, len(shapes))
		for k := range fields {
			v, err := strconv.Atoi(fields[k])
			if err != nil {
				return nil, nil, fmt.Errorf("bad count %q in line %q", fields[k], line)
			}
			counts[k] = v
		}

		regions = append(regions, Region{W: w, H: h, Counts: counts})
	}

	return shapes, regions, nil
}

func isRegionLine(s string) bool {
	return strings.Contains(s, "x") && strings.Contains(s, ":") && strings.IndexByte(s, 'x') < strings.IndexByte(s, ':')
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
