package main

import "fmt"

func main(){

	//事前に扱う要素素が見当がつく場合は、最初に make 関数で指定しておくほうがいい
	//要素数が容量を超えたとき、アドレスが新しく割り振られる
	pow := make([]int, 10) //事前に要素数がわかっている場合
	fmt.Println(pow)

	for index:= range pow{
		fmt.Println(1 << uint(index))
		// pow[index] = 1 << uint(index)
	}
}