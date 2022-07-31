package pipe4

func main() {
	ch1 := make(chan string)
	ch2 := make(chan chan string)
	ch2 <- ch1
}
