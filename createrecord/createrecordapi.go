package createrecord

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

func CreateTodo(w http.ResponseWriter, r *http.Request){
  //todo := getTodoJson(r)
  todo := createNewTodo(r)
  outputJson, jerr := json.Marshal(&todo)
  errorhandling.Errorhandle(jerr)
  jsonsettings.JsonHandle(w,r, outputJson)
}

func createNewTodo(r *http.Request) *structmanager.Todo {
  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/todo")
  errorhandling.Errorhandle(err)
  defer db.Close()

  // check, err2 := db.Prepare("select exists(select * from todo where title = ?)")
  // errorhandling.Errorhandle(err2)
  // log.Print(check.Exec(title))
  // if getsqlline.GetSingleTodo(query) != nil{
  //   panic("このtodoはすでに登録しています")
  // }

  ins, err1 := db.Prepare("insert into todo(title, deadline) values(?,?)")
  errorhandling.Errorhandle(err1)
  fc := jsonsettings.GetCreatingJson(r)
  title := fc.Title
  deadline := fc.Deadline
  ins.Exec(title, deadline)

  return getsqlline.GetSingleTodo(`select * from todo where title = "` + title + `" and deadline = "` + deadline + `"and deleted_at is null`)
}
