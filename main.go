package main

import (
	"fmt"

	"github.com/aldisaputra17/list_data/produk"
)

func main() {
	var harga = []int{3000, 4000000, 3000, 4000000, 4000000, 4000000}
	var min = produk.LowPrice(harga)
	var rate = []float32{5, 4.5, 4, 4.5, 3.5, 3.5}
	var highRate = produk.HighRate(rate)
	var likes = []int{150, 123, 250, 42, 90, 87}
	var maxLike = produk.Likes(likes)

	fmt.Printf("data : %v\nmin  : %v\n", harga, min)
	fmt.Printf("data : %v\nhighrate  : %v\n", rate, highRate)
	fmt.Printf("data : %v\nmaxlike  : %v\n", likes, maxLike)
}
