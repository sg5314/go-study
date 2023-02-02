package main

import (
	"fmt"
	"math"
	"runtime"
)

type Inken struct{
	id   int
	name string
	tag  string
}

func add (x, y string) (ans string, ans1 string){

	ans = y
	ans1 = x
	return
}

func check(x, n float64) bool{

	// float にする
	if v := math.Pow(x,2); v<n{
		return true
	}else{
		return false
	}
}

func OS_check(){
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func main()  {
	// fmt.Println(add("こんにちは","すが"))
	// var i int = 64
	// var f float64 = float64(i)
	// var u uint = uint(i)

	// fmt.Println(i,f,u)

	// sun:=0
	// for sun < 1000{
	// 	sun+=1
	// }
	// fmt.Println(sun)

	// fmt.Println(check(2, 5))
	// OS_check()

	// efer文はリソースの解放を行うような処理を書く際に使うことが多い
	defer fmt.Println("world")

	fmt.Println("hello")

}