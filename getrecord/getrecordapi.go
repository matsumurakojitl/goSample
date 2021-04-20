package getrecord

import (
  "encoding/json"
  "github.com/gosample/errorhandling"
  "github.com/gosample/structmanager"
  "net/http"
  "github.com/gosample/getsqlline"
  "github.com/gosample/jsonsettings"
)

func FetchAllItems(w http.ResponseWriter, r *http.Request){

  var todoList = structmanager.TodoList{}
  var slice1 []*structmanager.Todo

  todoList.List = getsqlline.GetTodoList("select * from todo where deleted_at is null", slice1)
  outputJson, jerr := json.Marshal(&todoList)
  errorhandling.Errorhandle(jerr)
  jsonsettings.JsonHandle(w,r, outputJson)
}
