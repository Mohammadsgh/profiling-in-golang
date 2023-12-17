package profiling_in_golang

import (
	"bytes"
	"strconv"
)

func createStringOne(size int) string {
	var str string
	for i := 0; i < size; i++ {
		str += strconv.Itoa(i) // Convert integer to string and concatenate
	}
	return str
}

func createStringTwo(size int) string {
	var buffer bytes.Buffer
	for i := 0; i < size; i++ {
		buffer.WriteString(strconv.Itoa(i))
	}

	return buffer.String()
}
