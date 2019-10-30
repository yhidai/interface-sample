package main

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserController_Add_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "bob"

	// ユーザーの登録
	userTable := NewMockUserTableInterface(ctrl)
	userTable.EXPECT().Count().Return(99)

	controller := UserController{userTable}
	message := controller.Add(name)
	assert.Equal(t, message, "ユーザーを登録しました。")
}

func TestUserController_Add_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "bob"

	// ユーザーの登録
	userTable := NewMockUserTableInterface(ctrl)
	userTable.EXPECT().Count().Return(99)

	// ユーザーの登録
	controller := UserController{userTable}
	message := controller.Add(name)
	assert.Equal(t, message, "100件以上登録できません。")

}
