package jsonsettings

import(
  "net/http"
  "fmt"
  "log"
)

func JsonHandle(w http.ResponseWriter, r *http.Request, outputJson []uint8){
  w.Header().Set("Content-Type", "application/json")
  fmt.Fprint(w, string(outputJson))
  log.Print(r.URL.Path)
}
