package main

func Increment(value *int) {
	if value == nil {
		return
	}
	*value++
}

func main() {
}
