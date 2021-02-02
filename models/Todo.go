package models

import (
	"Todogo/dao"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

type Todo struct {
	ID     string `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Status bool   `json:"status" db:"status"`
}

func Create(t *Todo) (err error) {
	u1 := uuid.NewV4().String()
	sql := "INSERT INTO todo(id, title,status) VALUES(?,?,?)"
	_, err1 := dao.Db.Exec(sql, u1, t.Title, t.Status)
	return err1
}

func FindList() (todoList []*Todo, err error) {
	sql := "SELECT id,title,status from todo"
	err1 := dao.Db.Select(&todoList, sql)
	return todoList, err1
}
func Update(id string) (err error) {
	var todo Todo
	sql := "SELECT id,title,status from todo where id = ?"
	if err := dao.Db.Get(&todo, sql, id); err != nil {
		return err
	}

	todo.Status = !todo.Status

	updateSql := "UPDATE todo set title = ?,status = ? where id = ?"

	result, err1 := dao.Db.Exec(updateSql, todo.Title, todo.Status, todo.ID)
	if err1 != nil {
		return err1
	}
	_, err2 := result.RowsAffected()

	if err2 != nil {
		return err2
	} else {
		return
	}

}
func Delete(id string) (err error) {
	deleteSql := "DELETE from todo where id = ?"

	result, err1 := dao.Db.Exec(deleteSql, id)
	if err1 != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return err1
	}
	_, err2 := result.RowsAffected()
	if err2 != nil {
		return err2
	} else {
		return
	}

}
