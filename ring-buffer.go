package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// Ring Buffer Structure
type RingBuffer struct {
	data       []*Data
	size       int
	lastInsert int
	nextRead   int
	emitTime   time.Time
}

// Data interface fed into the Ring Buffer
type Data struct {
	Stamp time.Time
	Value string
}

// Function to create a new RingBuffer
// takes in the buffer size
// returns a pointer to the newly created ringbuffer
func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		// make an array of pointers to the Data scruct
		// the size of the array will be determined by the passed in size parementer
		data:       make([]*Data, size),
		size:       size,
		lastInsert: -1,
	}
}

// This method will insert new data into the RingBuffer
// This method will create a pointer to the ringbuffer struct called 'r'
// This method will take in Data struct called input
func (r *RingBuffer) Insert(input Data) {
	// mod will always return the result of the iteration
	// array len = 5
	// (3 + 1) % 5 = 4
	// (4 + 1) % 5 = 0
	r.lastInsert = (r.lastInsert + 1) % r.size
	// The data the value of the lastInsert will used select the index of the data array
	// The data array at this value will be assigned a pointer to the input value
	// The data object takes in pointers to values
	r.data[r.lastInsert] = &input

	// If the read value catches up to the next write
	if r.nextRead == r.lastInsert {
		// iterate the read pointer
		r.nextRead = (r.nextRead + 1) % r.size
	}
}

// This method will read data out of the RingBuffer
// This method will create a pointer to the ringbuffer struct called 'r'
// This method will return an array of pointers to Data objects
func (r *RingBuffer) Emit() []*Data {
	// create a local output slice
	output := []*Data{}
	// iterate infinitly
	for {
		// as long as the data is not nil - add it to the output slice
		if r.data[r.nextRead] != nil {
			output = append(output, r.data[r.nextRead])
			// clean up
			r.data[r.nextRead] = nil
		}
		// If reads catch up with writes or nothing has been written - exit
		if r.nextRead == r.lastInsert || r.lastInsert == -1 {
			break
		}
		// iterate to the next index of the RingBuffer
		r.nextRead = (r.nextRead + 1) % r.size
	}
	return output
}

func main() {
	// create a new ring buffer of size 5
	rb := NewRingBuffer(5)
	fmt.Println("EMPTY TEST:")
	// use spew to Dump the memory
	spew.Dump(rb.Emit())
	// use a Rune to represent data
	currentRune := 'a' - 1
	// create iterator
	for i := 0; i < 10; i++ {
		// iterate rune by one
		currentRune++
		// create data
		rb.Insert(Data{
			Stamp: time.Now(),
			Value: string(currentRune),
		})
	}
	fmt.Println("FULL TEST:")
	spew.Dump(rb.Emit())
}
