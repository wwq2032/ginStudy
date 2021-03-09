package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
)

type msg struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Message string `json:"message"`
}
//返回JSON
func getJSON(r *gin.Engine) {
	r.GET("/json", func(c *gin.Context) {
		var data msg
		data.Name = "小王总"
		data.Age = 18
		data.Message = "hello world"
		c.JSON(http.StatusOK, data)
	})
}

//查询
func getQuery(r *gin.Engine){

	////localhost:9090/user/search?name=%E5%B0%8F%E7%8E%8B%E6%80%BB&age=18
	r.GET("/user/search", func(c *gin.Context) {
	//	temp:=msg{
	//		Name: "",
	//		Age: 0,
	//		Message: "hello world",
	//	}
	//
	//	temp.Name=c.Query("name")
	//	temp.Age,_=strconv.Atoi(c.Query("age"))
	//
	//
	//	//http://localhost:9090/user/search?msg[name]=wwq&msg[age]=18
	//
	//	//temp:=c.QueryMap("msg")
	//
	//
	//	//http://localhost:9090/user/search?msg=aaa&msg=bbb
	//	//temp:=c.QueryArray("msg")
	//
	//	//fmt.Printf("%v\n",temp)
	//
	//	c.JSON(http.StatusOK,temp)

	//直接拼接字符串

			name:=c.Query("name")
			age:=c.Query("age")

			//输出结果给浏览器
			c.JSON(http.StatusOK,gin.H{
				"name":name,
				"age":age,
			})

	})

}

func main() {
	fmt.Printf("hello,gin")
	//创建一个默认的路由引擎
	r := gin.Default()
	//getJSON(r)
	getQuery(r)

	//启动路由
	_ = r.Run(":9090")
}
