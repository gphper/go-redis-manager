/*
 * @Description:部署拷贝静态文件时忽略该文件
 * @Author: gphper
 * @Date: 2021-07-31 10:28:43
 */

package web

import (
	"embed"
	"html/template"
	"time"

	"io/fs"
	"net/http"
	"strings"

	"github.com/gphper/multitemplate"
)

//go:embed statics
var StaticPath embed.FS

//go:embed views
var viewPath embed.FS

var globalTemplateFun template.FuncMap

var StaticsFs http.FileSystem

func init() {
	static, err := fs.Sub(StaticPath, "statics")
	if err != nil {
		panic(err.Error())
	}
	StaticsFs = http.FS(static)

	globalTemplateFun = template.FuncMap{
		"formatAsDate": func(t time.Time, format string) string {
			return t.Format(format)
		},
		"judegInMap": func(find string, items map[string]struct{}) bool {
			_, ok := items[find]
			return ok
		},
		"equal": func(first string, second string) bool {
			return first == second
		},
		"equalInt": func(first int, second int) bool {
			return first == second
		},
		"add": func(key int, key1 int) int {
			return key + key1
		},
		"compareInt": func(key1 int, key2 int) bool {
			return key1 > key2
		},
	}
}

func LoadTemplates() multitemplate.Renderer {

	r := multitemplate.NewRenderer()

	layouts, err := fs.Glob(viewPath, "views/layout/*.html")

	if err != nil {
		panic(err.Error())
	}
	includes, err := fs.Glob(viewPath, "views/template/*/*.html")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))

		copy(layoutCopy, layouts)

		files := append(layoutCopy, include)

		dirSlice := strings.Split(include, "/")

		fileName := strings.Join(dirSlice[len(dirSlice)-2:], "/")

		r.AddFromFsFuncs(fileName, globalTemplateFun, viewPath, files...)
	}

	return r
}
