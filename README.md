###C-like strftime on Golang

#### How to use

	package main

	import(
		"fmt"
		"time"
		"github.com/dstarodubtsev/go_strftime"
	)
	
	func main(){
		// "2015-06-23 17:57:05", as expected
		fmt.Println(go_strftime.Strftime("%Y-%m-%d %H:%M:%S", time.Now()))
	}