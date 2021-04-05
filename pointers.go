package main

func getPointers(n *int) int {
	return *n = *n * *n
}

func returnPointers(n int) *int {
	return
}
