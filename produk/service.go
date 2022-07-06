package produk

func LowPrice(price []int) int {
	var min int
	for i, e := range price {
		switch {
		case i == 0:
			min = e
		case e < min:
			min = e
		}
	}
	return min
}

func HighRate(n []float32) float32 {
	var rate float32
	for i, v := range n {
		switch {
		case i == 0:
			rate = v
		case v > rate:
			rate = v
		}
	}
	return rate
}

func Likes(like []int) int {
	var max int
	for i, e := range like {
		switch {
		case i == 0:
			max = e
		case e > max:
			max = e
		}
	}
	return max
}
