package memo

import (
	"memo/app/internal/models"
	"memo/app/internal/services"
	"time"
)

//GetMemo データを取得する
func GetMemo(id string) models.Memo {

	memo := models.Memo{}

	services.DB.
		Select("*").
		Where("id=?", id).
		Find(&memo)

	return memo
}

//GetAllMemo 全データを取得する
func GetAllMemo() []models.Memo {

	memos := []models.Memo{}

	services.DB.
		Select("*").
		Find(&memos)

	return memos
}

//AddMemo データを登録する
func AddMemo(title, body string) int {

	memo := models.Memo{Title: title, Body: body, Created: time.Now()}
	services.DB.Create(&memo)

	return memo.ID
}

//UpdateMemo データを更新する
func UpdateMemo(id, title, body string) {

	memo := models.Memo{Title: title, Body: body}
	services.DB.
		Where("id=?", id).
		Update(&memo)
}

//DeleteMemo データを削除する
func DeleteMemo(id int) {

	memo := models.Memo{ID: id}
	services.DB.Delete(&memo)
}