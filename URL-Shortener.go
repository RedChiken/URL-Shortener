package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	//"html/template"
	//"github.com/GeertJohan/go.rice"
	//"os"
	//"path/filepath"
)

func main() {
	var result interface{}
	log := readLog("log.txt")

	r := gin.Default()
	//r.SetHTMLTemplate(templates())

	r.Static("/handlers", "./handlers")
	r.LoadHTMLGlob("templates/*")

	r.GET("/main", func(c *gin.Context) {
	/*
	1. js 함수에서 값을 받아오지 못함.
	이를 해결하기 template와 일부 외부 라이브러리를 이용해야
	하나, 이에 대한 이해도가 부족하여 구현하지 못함.
	*/
		url, _ := c.Get("url" +
			"")
		result = shorten(log, fmt.Sprint(url))
		print(result)
		c.HTML(200, "mainpage.tmpl", gin.H{
			"result" : result,
		})
		/*
		2. 연산 결과를 tmpl에 전달하고, 이를 받는 방법을 찾지 못함.
		이 또한 1번과 비슷한 이유로 구현하지 못함.
		*/
	})

	r.GET("/syntax", func(c *gin.Context){
	/*
	3. url을 동적으로 받지 못함.
	검색한 것과 달리 syntax가 충돌 난다면서 작동하지 않음.
	*/
		urlSyntax := c.Param("syntax")
		value, _ := log[urlSyntax]
		c.Redirect(http.StatusMovedPermanently, value)
	})

	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}

//func templates() *template.Template {
//	all := template.New("__main__").Funcs(template.FuncMap{})
//	templateBox := rice.MustFindBox("templates")
//	templateBox.Walk("/", func(path string, info os.FileInfo, err error) error{
//		if info.IsDir(){
//			return nil
//		}
//		if info.Name()[0] == '.'{
//			return nil
//		}
//		slashedPath := filepath.ToSlash(path)
//		template.Must(all.New(slashedPath).Parse(templateBox.MustString(path)))
//		return nil
//	})
//
//	str := fmt.Sprintf("Loaded HTML Template (%d):\n", len(all.Templates()))
//	for _, v := range all.Templates(){
//		str += fmt.Sprintf("\t - %s\n", v.Name())
//	}
//	fmt.Println(str)
//	return all
//}
