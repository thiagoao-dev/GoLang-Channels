package goroutines

func Simple(c chan int) {
	c <- 33
}
