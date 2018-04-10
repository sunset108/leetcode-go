func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	return check(p1, p2, p3, p4) || check(p1, p3, p2, p4) || check(p1, p2, p4, p3)
}

func dist(p1 []int, p2 []int) int {
	return (p1[0] - p2[0]) * (p1[0] - p2[0]) + (p1[1] - p2[1]) * (p1[1] - p2[1])
}

func check(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	return dist(p1, p2) > 0 && dist(p1, p2) == dist(p2, p3) && dist(p2, p3) == dist(p3, p4) && dist(p3, p4) == dist(p4, p1) && dist(p1, p3) == dist(p2, p4)
}
