package lab0

import (
	"bufio"
	"io"
	"strconv"
	"os"
	//"fmt"
)

// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
// Do NOT modify function signature.
func sumWorker(nums chan int, out chan int) {
	// TODO: implement me
	// HINT: use for loop over `nums`
	sum := 0
	for num := range nums {
		sum += num
	}

	out <- sum
}

func push_to_in_channel(numValues []int, in_chan chan<-int) {
	//fmt.Println("Hello")
	for _, value := range numValues {
		in_chan <- value
	}
	//fmt.Println("Hello there")
	close(in_chan)
}

// Read integers from the file `fileName` and return sum of all values.
// This function must launch `num` go routines running
// `sumWorker` to find the sum of the values concurrently.
// You should use `checkError` to handle potential errors.
// Do NOT modify function signature.
func sum(num int, fileName string) int {
	// TODO: implement me
	// HINT: use `readInts` and `sumWorker`
	// HINT: used buffered channels for splitting numbers between workers

	//1. Read the file / close when done implementing
	file, err := os.Open(fileName)
	checkError(err)
	//if err != nil {
        //      log.Fatal(err)
        //}
	defer file.Close()

	//2. Read integers from the file
	ints, err := readInts(file)
	checkError(err)

	//3. Create channels for inputs and output
	ch_in := make(chan int, len(ints))
	ch_out := make(chan int, num)

	//4. Launch 'num' go routines running 'sumWorker'
	//to find the sum of the values concurrently
	for i := 0; i < num; i++ {
		go sumWorker(ch_in, ch_out)
	}

	//5. Push all of integers into ch_in channel
	go push_to_in_channel(ints, ch_in)

	//6. Collect partial sums from ch_out to complete the total
	sum := 0
	for i := 0; i < num; i++ {
		sum += <- ch_out
	}

	//7. Close out ch_out and return the total sum
	close(ch_out)
	return sum
}

// Read a list of integers separated by whitespace from `r`.
// Return the integers successfully read with no error, or
// an empty slice of integers and the error that occurred.
// Do NOT modify this function.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
		elems = append(elems, val)
	}
	return elems, nil
}
