package updaterecord

import(
  "encoding/json"
  "github.com/gosample/errorhandling"
  "github.com/gosample/structmanager"
  "net/http"
  "github.com/gosample/getsqlline"
  "github.com/gosample/jsonsettings"
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func UpdateTodo(w http.ResponseWriter, r *http.Request){
  todo := updateSelectedTodo(r)
  outputJson, jerr := json.Marshal(&todo)
  errorhandling.Errorhandle(jerr)
  jsonsettings.JsonHandle(w,r, outputJson)
}

func updateSelectedTodo(r *http.Request) *structmanager.Todo{
  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/todo")
  errorhandling.Errorhandle(err)
  defer db.Close()

  upd, err1 := db.Prepare("update todo set title = ?, deadline = ? where id = ? and deleted_at is null")
  errorhandling.Errorhandle(err1)
  id := r.URL.Query().Get("id")
  title := r.URL.Query().Get("title")
  deadline := r.URL.Query().Get("deadline")
  upd.Exec(title, deadline, id)
  return getsqlline.GetSingleTodo(`select * from todo where deleted_at is null and id = ` + id)
}
