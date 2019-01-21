package main

import "fmt" //https://golang.org/pkg/fmt/

func test() {
	// Khai báo biến message có kiểu dữ liệu string
	var message string

	// Khai báo 3 biến c, python, java đều có kiểu dữ liệu là bool
	var c, python, java bool

	// Khai báo 2 biến i, j có kiểu dữ liệu là int và khởi tạo luôn giá trị cho chúng
	var i, j int = 1, 2

	// Khai báo ngắn gọn biến k và khởi tạo giá trị luôn cho nó.
	// Không dùng từ khóa var mà dùng dấu hai chấm, lúc này kiểu dữ liệu sẽ được ngầm định tùy theo giá trị của biến.
	k := 3
	t := 55           // kiểu int
	w := 67.8         // kiểu float64
	sum := i + int(j) // Để tính tổng cần phải ép kiểu j về int (sum = 55 + 67)
	fmt.Println(sum)  // Kết quả trả về là 122
	const Pi = 3.14

	a, b := swap("hello", "world")
	fmt.Println("Lại là Hello World", k, c, python, java, message, i, j, k, t, w, a, b)
}

// Hàm tính tổng, có 2 tham số truyền vào với kiểu dữ liệu int và kết quả trả về cũng là int
func add(x int, y int) int {
	return x + y
}

// Hàm swap trả về kết quả là 2 chuỗi
func swap(x, y string) (string, string) {
	return y, x
}

// Hàm tính tổng, có 2 tham số truyền vào với kiểu dữ liệu int và kết quả trả về cũng là int
// func add(x, y int) int {
// 	return x + y
// }
// Hàm split khai báo 2 kết quả trả về là x và y có kiểu dữ liệu là int
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// return có thể để trống, function sẽ tự động trả về x và y (không khuyến khích)
	return
}
