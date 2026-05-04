package main

import "fmt"

func main() {
	var x, y int
	var ikan [1000]float64

 	fmt.Scan(&x, &y)
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	jumlahWadah := (x + y - 1) / y

	var wadah [1000]float64

	index := 0
	for i := 0; i < jumlahWadah; i++ {
		sum := 0.0
		for j := 0; j < y && index < x; j++ {
			sum += ikan[index]
			index++
		}
		wadah[i] = sum
	}

	for i := 0; i < jumlahWadah; i++ {
		fmt.Print(wadah[i])
		if i != jumlahWadah-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	total := 0.0
	for i := 0; i < jumlahWadah; i++ {
		total += wadah[i]
	}
	rata := total / float64(jumlahWadah)

	fmt.Println(rata)
}