package main

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)
// エラー用確認用のモックを定義する。
type MockUserTableForAddSuccess struct {
	db *sql.DB
}

func (u MockUserTableForAddSuccess) Count() int {
	return 99
}

func TestUserController_Add_Success(t *testing.T) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:33316)/sample")
	if err != nil {
		log.Fatal(err.Error())
	}

	// 名前
	name := "bob"

	// ユーザーの登録
	userTable := &MockUserTableForAddSuccess{db}
	ctrl := UserController{userTable}
	message := ctrl.Add(name)
	assert.Equal(t, message, "ユーザーを登録しました。")
}

// エラー用確認用のモックを定義する。
type MockUserTableForAddFail struct {
	db *sql.DB
}

func (u MockUserTableForAddFail) Count() int {
	return 100
}


func TestUserController_Add_Fail(t *testing.T) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:33316)/sample")
	if err != nil {
		log.Fatal(err.Error())
	}

	// 名前
	name := "bob"

	// ユーザーの登録
	userTable := &MockUserTableForAddFail{db}
	ctrl := UserController{userTable}
	message := ctrl.Add(name)
	assert.Equal(t, message, "100件以上登録できません。")

}
