package secao07ponteiros

import "fmt"

func Main() {
	fmt.Println("Seção 07:")
	x := 5
	y := x
	y = 10
	printAll(x, y)
	printAllPointers(&x, &y)

	a := 5
	b := &a
	*b = 10
	printAll(a, *b)
	printAllPointers(&a, b)

	c := 10
	fmt.Println("C before reference and pointer:", c)
	changeValueByReference(c)
	fmt.Println("C after reference:", c)
	changeValueByPointer(&c)
	fmt.Println("C after pointer:", c)
}

func printAll(x int, y int) {
	fmt.Println(x, y)
	fmt.Println(&x, &y)
}

func printAllPointers(x *int, y *int) {
	fmt.Println(*x, *y)
	fmt.Println(x, y)
}

func changeValueByReference(x int) {
	x = 20
}

func changeValueByPointer(x *int) {
	*x = 20
}
