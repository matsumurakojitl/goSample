package jsonsettings

import (
  "io/ioutil"
  "github.com/gosample/errorhandling"
  "encoding/json"
  "net/http"
  "github.com/gosample/structmanager"
)

func getTodoJson(r *http.Request) structmanager.Todo {
  body,err1 := ioutil.ReadAll(r.Body)
  errorhandling.Errorhandle(err1)
  var todo structmanager.Todo
  err2 := json.Unmarshal(body, &todo)
  errorhandling.Errorhandle(err2)
  return todo
}

func GetCreatingJson(r *http.Request) structmanager.ForCreate {
  body,err1 := ioutil.ReadAll(r.Body)
  errorhandling.Errorhandle(err1)
  var fc structmanager.ForCreate
  err2 := json.Unmarshal(body, &fc)
  errorhandling.Errorhandle(err2)
  return fc
}
