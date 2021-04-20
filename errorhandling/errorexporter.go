package errorhandling

import(
  "fmt"
  "log"
)

func Errorhandle(err error){
  if err != nil {
    log.Print(err)
    fmt.Print(err)
    panic(err.Error())
  }
}
