// go build -o sumgo.exe sum.go

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func somesum(n int64, m int64) int64 {
	sum := 0.0
	var i int64
	for i = 0; i <= n; i++ {
		sum += math.Pow(float64(i), float64(m))
	}
	return int64(sum)
}

func main() {

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("failed to open")

	}

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

	n, err := strconv.ParseInt(text[0], 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	m, err := strconv.ParseInt(text[1], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("n = %d\nm = %d\nn + m = %d", n, m, n+m)

	f, err := os.Create("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	my_sum := strconv.FormatInt(somesum(n, m), 10)

	_, err = f.WriteString(my_sum)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	// fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
