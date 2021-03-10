package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"

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
func getQuery(r *gin.Engine) {

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

		name := c.Query("name")
		age := c.Query("age")

		//输出结果给浏览器
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})

	})

}

//获取form表单提交的参数
func getForm(r *gin.Engine) {
	r.POST("/index/search", func(context *gin.Context) {
		username := context.PostForm("username")
		address := context.PostForm("address")

		context.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
}

func getFormHTML(r *gin.Engine) {
	r.LoadHTMLFiles("./src/login.html", "./src/index.html")

	r.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", nil)
	})

	// login post

	r.POST("/login", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.HTML(http.StatusOK, "index.html", gin.H{
			"Username": username,
			"Password": password,
		})
	})
}

type login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func bindGet(r *gin.Engine) {
	//var login login

	//绑定JSON  ({"user": "q1mi", "password": "123456"})
	r.POST("/loginJSON", func(c *gin.Context) {
		var login login
		err := c.ShouldBind(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		}
	})

	//绑定表单 (user=q1mi&password=123456)
	r.POST("/loginFORM", func(c *gin.Context) {
		var login login
		//绑定
		err := c.ShouldBind(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		}
	})

	//绑定querryString (/loginQuery?user=q1mi&password=123456)
	r.GET("/loginQuery", func(c *gin.Context) {
		var login login
		//绑定
		err := c.ShouldBind(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		}
	})

}

func getURI(r *gin.Engine) {
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"address":  address,
		})
	})

}

func upload(r *gin.Engine) {
	r.LoadHTMLFiles("./src/upload.html")

	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("files")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		log.Println(file.Filename)

		dst := fmt.Sprintf("D:/tmp/%s", file.Filename)

		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s uploaded!", file.Filename),
		})
	})
}

func multiplyUpload(r *gin.Engine) {
	r.LoadHTMLFiles("./src/upload.html")

	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		files := form.File["files"]

		for index, file := range files {
			_ = index
			dst := fmt.Sprintf("D:/tmp/%s_%d", file.Filename, index)

			err = c.SaveUploadedFile(file, dst)
			if err != nil {
				fmt.Println(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("%s uploaded!", file.Filename),
			})
		}
	})
}

func httpRedirect(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")

	})
}

func routerRedirect(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {

		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})

	r.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
}

//匹配所有路由的方法 r.Any("/test", func(c *gin.Context) {...})
func anyRouter(r *gin.Engine) {
	r.LoadHTMLFiles("./src/404.html")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
}

/* 路由组 看起来层次清晰 支持嵌套
func main() {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {...})
		userGroup.GET("/login", func(c *gin.Context) {...})
		userGroup.POST("/login", func(c *gin.Context) {...})

	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {...})
		shopGroup.GET("/cart", func(c *gin.Context) {...})
		shopGroup.POST("/checkout", func(c *gin.Context) {...})
	}
	r.Run()
}

*/

func main() {
	fmt.Printf("hello,gin")
	//创建一个默认的路由引擎
	r := gin.Default()
	//getJSON(r)
	//getQuery(r)

	//getFormHTML(r)
	//getForm(r)

	//getURI(r)

	//bindGet(r)

	//upload(r)
	//multiplyUpload(r)

	//httpRedirect(r)
	//routerRedirect(r)

	//anyRouter(r)

	//GIN中间件

	//启动路由
	_ = r.Run(":9090")
}
