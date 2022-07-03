package main

import "fmt"

/*
generic hampir sama dengan interface, sama2 tidak memperdulikan tipe data
perbedaan nya adalah, interface akan membungkus data aslinya atau disebut underlying value,
sedangkan Generic kita perlu mendefinisikan cakupan tipe data yang kompatibel untuk dipakai saat pemanggilan code.
atau kita bisa menggunakan keyword (comparable)
*/

// contoh generic
func Sum(numbers []int) int {
	var total int

	for _, e := range numbers {
		total += e
	}

	return total
}

// tipe data V artinya adalah kompatibel
func SumGeneric[V int](numbersGeneric []V) V {
	var total V

	for _, e := range numbersGeneric {
		total += e
	}
	return total
}

func main() {

	total1 := Sum([]int{1, 2, 3, 4, 5})
	total2 := SumGeneric([]int{1, 2, 3, 4, 5})
	fmt.Println(total1)
	fmt.Println(total2)

}
