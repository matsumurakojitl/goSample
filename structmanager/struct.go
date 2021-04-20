package structmanager

import (
	"database/sql"
)

type Todo struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Deadline   string `json:"deadline"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
  DeletedAt  sql.NullString `json:"deleted_at"`
}

type TodoList struct{
  List []*Todo `json:"todo"`
}

type Status struct{
  Status string `json:"status"`
}

type ForCreate struct{
	Title      string `json:"title"`
	Deadline   string `json:"deadline"`
}
