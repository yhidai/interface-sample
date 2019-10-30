package main

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestUserController_Add_Success(t *testing.T) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:33316)/sample")
	if err != nil {
		log.Fatal(err.Error())
	}

	// 名前の取得
	name := "bob"

	// ユーザーの登録
	ctrl := UserController{db}
	message := ctrl.Add(name)
	assert.Equal(t, message, "ユーザーを登録しました。")

}

//func TestUserController_Add_Fail(t *testing.T) {
//	db, err := sql.Open("mysql", "root:password@tcp(localhost:33316)/sample")
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//
//	// 名前の取得
//	name := os.Args[0]
//
//	// ユーザーの登録
//	ctrl := UserController{db}
//	message := ctrl.Add(name)
//	assert.Equal(t, message, "100件以上登録できません。")
//
//}
