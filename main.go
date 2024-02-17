package main

import (
	"app/ent"
	"app/ent/todo"
	"context"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var _ctx context.Context
var _ent *ent.Client

func init() {
	_ctx = context.Background()
}
func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed openin sqlite %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(_ctx); err != nil {
		log.Fatalf("failed creating schema resources :%v", err)
	}
	CreateTodo("e", "e", true)

}

func CreateTodo(name, description string, status bool) (todo *ent.Todo, err error) {
	todo, err = _ent.Todo.Create().
		SetName(name).
		SetDescription(description).
		SetStatus(false).
		Save(_ctx)
	if err != nil {
		return nil, errors.New("entity: failed to create the query " + err.Error())

	}
	return
}

func QueryTodo(name string) (res *ent.Todo, err error) {
	res, err = _ent.Todo.Query().
		Where(todo.Name(name)).
		Only(_ctx)

	if err != nil {
		return nil, errors.New("entity failed queryimg query" + err.Error())
	}
	return
}

func UpdateTodo(status bool, id int) (err error) {
	todo, err := _ent.Todo.Get(_ctx, id)
	if err != nil {
		return errors.New("failed to query" + err.Error())
	}

	err = todo.Update().SetStatus(status).Exec(_ctx)

	if err != nil {
		return errors.New("failed to update" + err.Error())
	}
	return
}

func GetTodo(id int) (todo *ent.Todo, err error) {
	todo, err = _ent.Todo.Get(_ctx, id)

	if err != nil {
		return nil, errors.New("failed to query" + err.Error())
	}

	return
}

func RemoveTodo(id int) error {
	err := _ent.Todo.DeleteOneID(id).Exec(_ctx)
	if err != nil {
		err = errors.New("failed to delete query" + err.Error())
	}
	return err
}
