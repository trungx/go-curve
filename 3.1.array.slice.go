package main

import "fmt"

func xxx() {
	var a [2]string

	// Gán giá trị cho các phần tử trong mảng
	a[0] = "Hello"
	a[1] = "World"

	// In kết quả ra console
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// Khởi tạo một mảng gồm 6 số int và gán luôn giá trị cho nó
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Khởi tạo mảng nhưng không ghi rõ kích thước (thay bằng dấu ba chấm),
	// trình biên dịch sẽ tự hiểu dựa vào số phần tử đã khai báo
	numbers := [...]int{12, 78, 50}
	fmt.Println(numbers)

	a1 := [...]int{1, 2, 3, 4, 5}
	b1 := a1  // b là một array mới có giá trị giống a
	b1[0] = 9 // Thay đổi giá trị một phần tử của b

	fmt.Println("a is ", a1) // In ra 1 2 3 4 5
	fmt.Println("b is ", b1) // In ra 9 2 3 4 5

	// Khởi tạo Array primes
	primes2 := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("a is ", primes2)

	// Khởi tạo Slice s bằng cách cắt từ phần tử ở vị trí 1 (low) đến phần tử ở vị trí 3 (high - 1) của Array primes
	var s []int = primes[1:4]

	// In ra giá trị của Slice s
	fmt.Println(s) // Giá trị của s là [3, 5, 7]

	s2 := []int{2, 3, 5, 7, 11, 13}

	s2 = s2[0:0] // s = [], len(s) = 0, cap(s) = 6
	s2 = s2[0:4] // s = [2, 3, 5, 7], len(s) = 4, cap(s) = 6
	s2 = s2[2:4] // s = [5, 7], len(s) = 2, cap(s) = 4, cap được tính từ vị trí số 2 trở đi
	s2 = s2[0:4] // s = [5, 7, 11, 13], len(s) = 4, cap(s) = 4

	s3 := []int{2, 3, 5, 7, 11, 13}

	s3 = s3[:0] // s = [0:0]
	s3 = s3[:4] // s = [0:4]
	s3 = s3[2:] // s = [2:len(s)] => s = [2:4]
	s3 = s3[:4] // s = [0:4]

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	a3 := make([]int, 5)    // len(a)=5
	b3 := make([]int, 0, 5) // len(b)=0, cap(b)=5
	fmt.Println(a3, b3)

	numbers := [4]int{1, 2, 3, 4}

	a := numbers[0:2] // a = [1, 2]
	b := numbers[1:3] // b = [2, 3]

	b[0] = 5 // Thay đổi giá trị phần tử đầu tiên của Slice b

	fmt.Println(a, b)    // a = [1, 5], b = [5, 3]
	fmt.Println(numbers) // numbers = [1, 5, 3, 4]

	var s []int

	// Append có thể hoạt động với nil slice.
	s = append(s, 0) // s = [0]

	// Append thêm một phần tử vào slice.
	s = append(s, 1) // s = [0, 1]

	// Append thêm nhiều phần tử vào slice.
	s = append(s, 2, 3, 4) // s = [0, 1, 2, 3, 4]

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("i = %d, v = %d \n", i, v)
	}

	var pow2 = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for _, v := range pow2 {
		fmt.Printf("v = %d \n", v)
	}
}
