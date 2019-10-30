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

	// 名前
	name := os.Args[0]

	// ユーザーの登録
	ctrl := UserController{db}
	message := ctrl.Add(name)

	// 完了メッセージを表示
	fmt.Println(message)
}

type UserController struct {
	db *sql.DB
}

func (c *UserController) Add(name string) string {
	// ユーザー名が存在しているかチェックする
	count, _ := models.Users().Count(context.Background(), c.db)
	if count >= 100 {
		return "100件以上登録できません。"
	}
	// 存在しなければ登録する

	// DB登録処理 ・・・・

	return "ユーザーを登録しました。"
}

