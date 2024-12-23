package util

type BFSResult struct {
	Distance int
	Point    Point
}

func BFSManyEnds(grid Grid, start Point, endFunc func(BFSResult) bool, getNeighbours func(BFSResult) []Point) []BFSResult {
	q := make([]BFSResult, 0)
	q = append(q, BFSResult{Distance: 0, Point: start})

	ends := []BFSResult{}
	for len(q) > 0 {
		c := q[0]
		q = q[1:]

		if endFunc(c) {
			ends = append(ends, c)
		}

		for _, n := range getNeighbours(c) {
			q = append(q, BFSResult{Distance: c.Distance + 1, Point: n})
		}
	}

	return ends
}
