package main

func Sum(numbers []int) int {
	if len(numbers) == 0 || numbers == nil {
		return 0
	}
	var sum int = 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func main() {
}
