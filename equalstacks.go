package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func heigh(h []int32) (sum int32) {
	for _, v := range h {
		sum += v
	}
	return sum
}

func removeTopElt(h []int32) []int32 {
	return h[1:]
}

func allEqual(heighs ...int32) bool {
	h := heighs[0]
	for _, v := range heighs {
		if v != h {
			return false
		}
	}
	return true
}

func biggestHeigh(heighs ...int32) (pos int) {
	var max int32
	pos = -1
	for i, v := range heighs {
		if v > max {
			max = v
			pos = i
		}
	}
	return pos
}

/*
 * Complete the equalStacks function below.
 */
func equalStacks(h1 []int32, h2 []int32, h3 []int32, w io.Writer) int32 {
	stacks := [][]int32{h1, h2, h3}

	var heighs []int32
	for i := 0; i < len(stacks); i++ {
		heighs = append(heighs, heigh(stacks[i]))
	}

	for {
		//for _, s := range stacks {
		//    fmt.Fprint(w, heigh(s), " ")
		//}
		//fmt.Fprintln(w)

		if allEqual(heighs...) {
			break
		}

		pos := biggestHeigh(heighs...)
		heighs[pos] -= stacks[pos][0]
		stacks[pos] = removeTopElt(stacks[pos])
	}
	return heighs[0]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	n1N2N3 := strings.Split(readLine(reader), " ")

	n1Temp, err := strconv.ParseInt(n1N2N3[0], 10, 64)
	checkError(err)
	n1 := int32(n1Temp)

	n2Temp, err := strconv.ParseInt(n1N2N3[1], 10, 64)
	checkError(err)
	n2 := int32(n2Temp)

	n3Temp, err := strconv.ParseInt(n1N2N3[2], 10, 64)
	checkError(err)
	n3 := int32(n3Temp)

	h1Temp := strings.Split(readLine(reader), " ")

	var h1 []int32

	for h1Itr := 0; h1Itr < int(n1); h1Itr++ {
		h1ItemTemp, err := strconv.ParseInt(h1Temp[h1Itr], 10, 64)
		checkError(err)
		h1Item := int32(h1ItemTemp)
		h1 = append(h1, h1Item)
	}

	h2Temp := strings.Split(readLine(reader), " ")

	var h2 []int32

	for h2Itr := 0; h2Itr < int(n2); h2Itr++ {
		h2ItemTemp, err := strconv.ParseInt(h2Temp[h2Itr], 10, 64)
		checkError(err)
		h2Item := int32(h2ItemTemp)
		h2 = append(h2, h2Item)
	}

	h3Temp := strings.Split(readLine(reader), " ")

	var h3 []int32

	for h3Itr := 0; h3Itr < int(n3); h3Itr++ {
		h3ItemTemp, err := strconv.ParseInt(h3Temp[h3Itr], 10, 64)
		checkError(err)
		h3Item := int32(h3ItemTemp)
		h3 = append(h3, h3Item)
	}

	result := equalStacks(h1, h2, h3, writer)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
