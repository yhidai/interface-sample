package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yhidai/interface-sample/models"
	"log"
	"os"
)

/*
 * usersテーブルへデータを保存する。
 * 10件以上データが存在する場合、エラーメッセージを表示する。
 */
func main() {

	// DBへ接続
	db, err := sql.Open("mysql", "root:password@tcp(localhost:33316)/sample")
	if err != nil {
		log.Fatal(err.Error())
	}

	// 引数からユーザー名を取得
	name := os.Args[0]


	// 処理のセットアップ
	userTable := &UserTable{db}
	ctrl := UserController{userTable}

	// ユーザーの登録
	message := ctrl.Add(name)
	fmt.Println(message)
}

type UserController struct {
	userTable UserTableInterface
}

func (c *UserController) Add(name string) string {
	// ユーザー件数をカウント
	count := c.userTable.Count()
	if count >= 100 {
		return "100件以上登録できません。"
	}
	// 存在しなければ登録する

	// DB登録処理 ・・・・

	return "ユーザーを登録しました。"
}

type UserTableInterface interface {
	Count() int
}

type UserTable struct {
	db *sql.DB
}

func (u UserTable) Count() int {
	count, _ := models.Users().Count(context.Background(), u.db)
	return int(count)
}