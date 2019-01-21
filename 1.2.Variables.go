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
    fmt.Println("Lại là Hello World",k,c,python,java,message,i,j,k)
}
func main() {
	test()
}
