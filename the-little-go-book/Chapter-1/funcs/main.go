package main

func log(message string) {
	println(message)
}

func math(a int, b int) (int, int) {
	return a + b, a * b
}

func main() {
	one, two := math(3, 2)
	println(one, two)
}

