package gobyexample

import (
	"fmt"
	"strconv"
)

func main() {
	b, _ := strconv.ParseBool("True")
	fmt.Println(b)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	u, _ := strconv.ParseUint("1234", 0, 64)
	fmt.Println(u)
}
