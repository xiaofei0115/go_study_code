package main

import (
	"fmt"	"math/cmplx"
)

func variableZeroValue(){
	a,b,c,d:=1,2,true,"s"
	fmt.Println(a,b,c,d)

}

func enums()  {
	const(
		cpp =0
		java =1
		python =2
		golang =3
	)
	fmt.Println(cpp,java,python,golang)
}
func main(){
	fmt.Println("hello,world")
	variableZeroValue()
	enums()
}