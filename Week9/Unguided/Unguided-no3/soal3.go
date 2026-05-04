package main

import "fmt"

func cariMin(arr []float64, n int) float64 {
	min := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func cariMax(arr []float64, n int) float64 {
	max := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func cariRata(arr []float64, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	arr := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	min := cariMin(arr, n)
	max := cariMax(arr, n)
	rata := cariRata(arr, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}