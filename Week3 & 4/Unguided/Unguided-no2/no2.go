package main

import "fmt"

// f(x) = x^2
func f(x int) int {
	return x * x
}

// g(x) = x - 2
func g(x int) int {
	return x - 2
}

// h(x) = x + 1
func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai y : ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai z : ")
	fmt.Scan(&c)

	// (f o g o h)(a) = f(g(h(a)))
	fmt.Println("F(G(H(", a, "))) :", f(g(h(a))))

	// (g o h o f)(b) = g(h(f(b)))
	fmt.Println("G(H(F(", b, "))) :", g(h(f(b))))

	// (h o f o g)(c) = h(f(g(c)))
	fmt.Println("H(F(G(", c, "))) :", h(f(g(c))))
}