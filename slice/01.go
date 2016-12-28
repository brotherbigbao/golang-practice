package main

import "fmt"

func main() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Println(months)
	fmt.Println(len(months))
	fmt.Println(cap(months))

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)

	for k, v := range summer {
		fmt.Println(k)
		fmt.Println(v)
	}

	summer[0] = "Wrong"
	fmt.Println(summer)
	fmt.Println(months)

	newSummer := summer[:5]
	fmt.Println(newSummer)

	//out of range
	//nnewSummer := summer[:20]
	//fmt.Println(nnewSummer)

	for k, v := range newSummer {
		fmt.Printf("%d=>%s\n", k, v)
	}
	//fmt.Println(newSummer[5])
}