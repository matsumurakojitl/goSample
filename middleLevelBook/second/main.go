package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
  e := echo.New()
	//e.GET("/:id", getPathParameter) // path parameterを取れない
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.FormValue("hoge")) // query
	// fmt.Fprintf(w, "Hello World")
// 	multi_line := `Hey!! we
// are going to
// write multiline strings
// in Go.
// `
// 	fmt.Fprintf(w, multi_line)

}

// func getPathParameter(c echo.Context) error {
// 	id := c.Param("id")
//   fmt.Println(id)
// 	return c.String(http.StatusOK, "id = "+id)
// }

// import(
//   "io"
//   "log"
//   "net/http"
//   "os"
//   "path"
//   "path/filepath"
// )
//
// func main(){
//   cwd,err := os.Getwd()
//   if err != nil {
// 		log.Fatal(err)
// 	}
//
//   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
//     if ok,err := path.Match("/data/*.txt", r.URL.Path);
//     err!= nil || !ok {
//       http.Notfound(w,r)
//       return
//     }
//     name := filepath.Join(cwd, "data", filepath.Base(r.URL.Path))
//     f,err := os.Open(name)
//     if err != nil {
//   		http.NotFound(w,r)
//       log.Fatal(err)
//       return
//   	}
//     defer f.Close()
//     io.Copy(w,f)
//   })
//   http.ListenAndServe(":8080", nil)
// }
