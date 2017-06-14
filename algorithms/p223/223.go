package p223

/**
Find the total area covered by two rectilinear rectangles in a 2D plane.

Each rectangle is defined by its bottom left corner and top right corner as shown in the figure.

Rectangle Area
Assume that the total area is never beyond the maximum possible value of int.


*/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	ares := (C-A)*(D-B) + (G-E)*(H-F)
	if C <= E || A >= G || B >= H || D <= F {
		return ares
	}
	return ares - (min(C, G)-max(A, E))*(min(D, H)-max(B, F))
}
