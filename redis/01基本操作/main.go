package main



func main() {
	defer client01.Close()
	wg.Add(1)
	go 	doSet()

	wg.Wait()
}
