package main

func main() {
	var s []int
	s = append(s, 3)
	s = append(s, 2)
	s = append(s, 8)
	s = append(s, 8)
	println(len(s))
	println(s[0])
	println(s[1])
}
