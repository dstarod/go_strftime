###C-like strftime on Golang

#### How to use

	package main

	import(
		"fmt"
		"time"
		"github.com/dstarodubtsev/go_strftime"
	)
	
	func main(){
		// http://www.cplusplus.com/reference/ctime/strftime/
		fmt.Println( go_strftime.Strftime("%Y-%m-%d %H:%M:%S", time.Now()) )
	}