package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input a string, Coltxi: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)

	fmt.Println("Please input number, Coltxi: ")
	str, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	n, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println("Who is against the queen, will die! ", err)
	} else {
		fmt.Println("Your number is: ", n)
	}
}
