package structmanager

import (
	"database/sql"
)

func TodoStruct(id int, title string, deadline string, status int, createTime string, updateTime string) (r *Todo) {
	r = new(Todo)
	r.ID = id
	r.Title = title
  r.Deadline = deadline
  r.Status = status
  r.CreateTime = createTime
  r.UpdateTime = updateTime
  r.DeletedAt = sql.NullString{String: "", Valid: false}
	return r
}

func StatusStruct(status string) (r *Status) {
	r = new(Status)
	r.Status = status
	return r
}
