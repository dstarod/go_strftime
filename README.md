###C-like strftime on Golang

Implemented most [useful patterns](http://www.cplusplus.com/reference/ctime/strftime/).
Skipped `%U`, `%W` (not ISO week number) and modifiers `E` and `O`.

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
