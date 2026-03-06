package main
import "fmt"
func main() {
	var berat, tinggi, bmi float64
	fmt.Print("Masukkan berat badan (kg): ")
	fmt.Scanln(&berat)
	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scanln(&tinggi)
	bmi = berat / (tinggi * tinggi)
	fmt.Printf("BMI Anda adalah: %.2f\n", bmi)
}