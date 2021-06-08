package main

import (
	"fmt"
	"io/ioutil"
)

func bounded(v int) int {
	if v > 100 {
		return 100
	} else if v < 0 {
		return 0
	} else {
		return v
	}
}
func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "c"
	case score < 90:
		g = "b"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong scroe %d", score)) //panic终止运行，并打印出错误
	}
	return g
}

func convertToBin(n int) string{
	result:=""
	for ;n>0;n/2={
		lsb:=n%2
		result=strconv.Itoa(lsb)+result
	}
	return result
}


func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Print(grade(200))
	fmt.Println(
		convertToBin(5)
		convertToBin(13)
	)
}
