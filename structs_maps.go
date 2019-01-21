package main

import "fmt"

// Định nghĩa một struct với từ khóa type
type Student struct {
	name string
	age  int
}

func yyy() {
	// Khởi tạo biến s1 có giá trị là struct Student
	s1 := Student{"Robin", 30} // {"Robin", 30}

	// Khởi tạo biến s2 có giá trị là struct Student với 1 field là name
	// Field còn lại sẽ có giá trị mặc định (zero value)
	s2 := Student{name: "Robin"} // {"Robin", 0}

	// Khởi tạo biến s3 có giá trị là struct Student và không khai báo giá trị cho trường nào
	//s3 := Student{} // {"", 0}

	// Thay đổi giá trị field trong struct
	// s3.name = "Robert"
	// s3.age = 25
	s3 := Student{"Job", 30}

	fmt.Println(s3) // s3 = {"Robert", 25}
	if s1 == s2 {
		fmt.Println("s1 = s2") // s1 bằng s2
	} else {
		fmt.Println("s1 != s2")
	}

	if s2 == s3 {
		fmt.Println("s2 = s3")
	} else {
		fmt.Println("s2 != s3") // s2 khác s3
	}

	// Định nghĩa biến demoMap có kiểu dữ liệu map với key kiểu string và value kiểu int
	var demoMap map[string]int

	// Map không so sánh được như struct, nhưng có thể dùng toán tử == để kiểm tra nil
	if demoMap == nil {
		fmt.Println("Map có giá trị nil.")

		// Tạo Map bằng hàm make
		demoMap = make(map[string]int)
	}

}
