package main

import "fmt"

type Suga struct{
	X int
	Y int
}

func main(){
	s := Suga{1,2}

	s.X = 22

	qq := []int{1,2,3,4,5,6}
	qq = qq[:4]
	qq = append(qq,99)
	fmt.Println(qq)
	// printSlice(qq)　

	var num []int
	num = append(num, 2)
	
	fmt.Println(num)

	//配列要素を取り出す
	for index,value := range qq{
		fmt.Println(index,value)
	}
}