package main

//mengirim ke channel output hasil	pangkat 2 nilai dari channel input
func squareWorker(input <-chan int, output chan<- int) {
	for i := 1; i <= 10; i++ {
		// TODO: answer here
		c := <-input
		output <- c * c
	}
}

//mengirim 1-n angka ke squareWorker
func numberGenerator(n int, input chan<- int) {
	for i := 1; i <= n; i++ {
		// TODO: answer here
		input <- i
	}
}
