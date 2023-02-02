package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "time"
) 

type Todo struct{
	Name string
	Content string
	Status  int
	CreatedAt string//日時
}

func main() {
  r := gin.Default()
	//   r.GET("/ping", func(c *gin.Context) {
	//     c.J SON(http.StatusOK, gin.H{
	//       "message": "pong",
	//     }) 
	//   })

	// todo := make([]Todo, 0, 5)
	// t := Todo{
	// 	Name: "suga",
	// 	Content: "友達と遊ぶ",
	// 	Status: 0,
	// 	CreatedAt: time.Now(),
	// }
	// todo = append(todo, t)


	todo := []Todo{
		Todo{
			Name: "suga",
			Content: "友達と遊ぶ",
			Status: 0,
			CreatedAt: time.Now().Format("2006-01-02"),
		},
		Todo{
			Name: "taro",
			Content: "早く起きる",
			Status: 0,
			CreatedAt: time.Now().Format("2006-01-02"),
		},
		Todo{
			Name: "jiro",
			Content: "買い物へ行く",
			Status: 0,
			CreatedAt: time.Now().Format("2006-01-02"),
		},
		Todo{
			Name: "saburo",
			Content: "旅行にいく",
			Status: 0,
			CreatedAt: time.Now().Format("2006-01-02"),
		},
	}
	

	r.LoadHTMLFiles("./template/index.html")

	r.GET("/todo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"now": time.Now().Format("2006-01-02 15:04:05"),
			"todo": todo,
		})
	})

  r.Run(":3000") // listen and serve on 0.0.0.0:3000 (for windows "localhost:3000")
}