package getsqlline

import (
  "github.com/gosample/structmanager"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/gosample/errorhandling"
)

func GetTodoList(query string, slice1 []*structmanager.Todo) []*structmanager.Todo{
  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/todo")
  errorhandling.Errorhandle(err)
  defer db.Close()

  rows, err := db.Query(query)
  errorhandling.Errorhandle(err)
  defer rows.Close()

  for rows.Next() {
    var todo structmanager.Todo
    err := rows.Scan(&todo.ID, &todo.Title, &todo.Deadline, &todo.Status, &todo.CreateTime, &todo.UpdateTime, &todo.DeletedAt)
    errorhandling.Errorhandle(err)
    slice1 = append(slice1, structmanager.TodoStruct(todo.ID, todo.Title, todo.Deadline, todo.Status, todo.CreateTime, todo.UpdateTime))
  }
  rowErr := rows.Err()
  errorhandling.Errorhandle(rowErr)
  return slice1
}

func GetSingleTodo(query string) *structmanager.Todo{
  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/todo")
  errorhandling.Errorhandle(err)
  defer db.Close()

  row := db.QueryRow(query)
  var todo structmanager.Todo
  err1 := row.Scan(&todo.ID, &todo.Title, &todo.Deadline, &todo.Status, &todo.CreateTime, &todo.UpdateTime, &todo.DeletedAt)
  errorhandling.Errorhandle(err1)

  return structmanager.TodoStruct(todo.ID, todo.Title, todo.Deadline, todo.Status, todo.CreateTime, todo.UpdateTime)
}
