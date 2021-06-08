package main

import(
	"io/ioutil"
	"fmt"
)

func bounded(v int) int{
	if v>100{
		return 100
	}else if v<0{
		return 0
	}else{
		return v
	}
}

func main(){
	const filename="abc.txt"
	contents,err:=ioutil.ReadFile(filename)
	if err !=nil{
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n",contents)
	}
}