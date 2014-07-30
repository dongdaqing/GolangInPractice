package main

import (
	"bufio"
	"fmt"
	// "io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("../../conf/httpurl.txt")
	defer f.Close()
	if nil == err {
		// buff := bufio.NewReader(f)
		// for {
		// 	line, err := buff.ReadString('\n')
		// 	if err != nil || io.EOF == err {
		// 		break
		// 	}
		// 	fmt.Println(line)
		// 	// fmt.Printf("%q\n", strings.Split(line, "&"))
		// 	// teststrings := strings.Split(line, "&")
		// 	// fmt.Println(teststrings[0])
		// }
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
			fmt.Printf("%q\n", strings.Split(line, "&"))
		}
	}
}
