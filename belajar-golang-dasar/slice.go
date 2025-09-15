package main

import "fmt"

func main() {

	var data = [6]string{"rida", "adila", "tes", "surabaya", "oke"}

	var slice1 = data[1:4]
	fmt.Println(slice1)

	var slice2 []string = data[:]
	fmt.Println(slice2)

	var slice3 []string = data[3:]
	fmt.Println(slice3)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	fmt.Println(days)
	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "libur baru")
	fmt.Println(daySlice2)
	fmt.Println(daySlice1)
	fmt.Println(days)
	daySlice2[0] = "sabtu lama"
	fmt.Println(daySlice2)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "januari"
	newSlice[1] = "februari"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "maret")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))
	newSlice2[0] = "juni"
	fmt.Println(newSlice)
	fmt.Println(newSlice2)

	fromSlice := days[:]
	copyDays := make([]string, len(fromSlice), cap(fromSlice))
	copy(copyDays, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(copyDays)
}
