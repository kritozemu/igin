package main

import (
	"imitate_gin/igin"
	"net/http"
)

//type student struct {
//	Name string
//	Age  int8
//}
//
//func FormatAsDate(t time.Time) string {
//	year, month, day := t.Date()
//	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
//}
//
//func main() {
//	r := igin.NewEngine()
//	r.Use(igin.Logger())
//	r.SetFuncMap(template.FuncMap{
//		"FormatAsDate": FormatAsDate,
//	})
//	r.LoadHTMLGlob("templates/*")
//	r.Static("/assets", "./static")
//
//	stu1 := &student{Name: "Dora", Age: 20}
//	stu2 := &student{Name: "Jack", Age: 22}
//	r.GET("/", func(c *igin.Context) {
//		c.HTML(http.StatusOK, "css.tmpl", nil)
//	})
//	r.GET("/students", func(c *igin.Context) {
//		c.HTML(http.StatusOK, "arr.tmpl", igin.H{
//			"title":  "igin",
//			"stuArr": [2]*student{stu1, stu2},
//		})
//	})
//
//	r.GET("/date", func(c *igin.Context) {
//		c.HTML(http.StatusOK, "custom_func.tmpl", igin.H{
//			"title": "igin",
//			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
//		})
//	})
//
//	r.Run(":9999")
//}

func main() {
	r := igin.Default()
	r.GET("/", func(c *igin.Context) {
		c.String(http.StatusOK, "Hello IGIN\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *igin.Context) {
		names := []string{"IGIN"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":8080")
}
