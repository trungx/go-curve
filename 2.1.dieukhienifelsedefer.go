package main

import (
	"fmt"
	"math"
)

func dieukhienluong() {
	// Tính tổng các số từ 0 - 9
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	sum2 := 0
	for sum2 < 10 {
		sum2 += sum2
	}
	fmt.Println(sum2)
	switch num := 10; {
	case num < 50:
		fmt.Printf("%d < 50\n", num) // In ra 10 < 50
	case num < 100:
		fmt.Printf("%d < 100\n", num) // Lệnh này không chạy mặc dù cũng thỏa mãn điều kiện
	default:
		fmt.Printf("I don't know", num)
	}
	num := 10

	switch { // Tương đương với switch true
	case num >= 0 && num <= 50:
		fmt.Printf("%d < 50 \n", num) // In ra 10 < 50
		fallthrough
	case num < 100:
		fmt.Printf("%d < 100 \n", num) // In ra 10 < 100
	default:
		fmt.Printf("I don't know %d", num)
	}
	defer fmt.Println("World") // Hoãn lệnh in ra chữ "World"
	fmt.Println("Hello")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // In ra giá trị của biến i
	}
}
func pow(x, n, limit float64) float64 {
	// Khai báo biến v trong lệnh điều kiện của if sẽ chỉ sử dụng được trong block if hoặc else
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)
	}

	// Không sử dụng được biến v ở bên ngoài, ví dụ return v sẽ báo lỗi
	return limit
}
