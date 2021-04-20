package deleterecord

import(
  "encoding/json"
  "github.com/gosample/errorhandling"
  "github.com/gosample/structmanager"
  "net/http"
  "github.com/gosample/jsonsettings"
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DeleteTodo(w http.ResponseWriter, r *http.Request){
  deleteSelectedTodo(r)
  status := structmanager.StatusStruct("success")
  outputJson, jerr := json.Marshal(&status)
  errorhandling.Errorhandle(jerr)
  jsonsettings.JsonHandle(w,r, outputJson)
}

func deleteSelectedTodo(r *http.Request){
  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/todo")
  errorhandling.Errorhandle(err)
  defer db.Close()

  upd, err1 := db.Prepare("update todo set deleted_at = current_timestamp() where id = ?")
  errorhandling.Errorhandle(err1)
  id := r.URL.Query().Get("id")
  upd.Exec(id)
}
