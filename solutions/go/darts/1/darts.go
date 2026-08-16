package darts

func Score(x, y float64) int {
	squaredDistance := x*x + y*y
    switch {
        case squaredDistance > 100:
        	return 0
        case squaredDistance > 25:
        	return 1
        case squaredDistance > 1:
        	return 5
        default:
        	return 10 
    }
}
