package main

import(
	"fmt"
	"log"
	"os"
)

func main(){
	f, err := os.Create("example.txt")

	if err != nil{
		log.Println(err)
		return
	}

	defer func(){
		_ = f.Close()
	}()

	if _ , err := f.Write([]byte("content")); err!= nil{
		log.Println(err)
		return
	}

	fmt.Println("ok")
}


// package main

// import (
//     "fmt"
//     "log"
//     "os"
// )

// func main() {
//     f, err := os.Create("example.txt")
//     if err != nil {
//         log.Println(err)
//         return
//     }

//     if _, err := f.Write([]byte("content")); err != nil {
//         _ = f.Close()
//         log.Println(err)
//         return
//     }
//     _ = f.Close()
//     fmt.Println("ok")
// }