package homework012

import (
	"fmt"
)

// 問題5
// User構造体を実装してください。

// 問題5 を実装してください。
// 中身を適宜変更してください。引数も変更して大丈夫です。
func insertUsers() {
}

// 問題5 を実装してください。
// 中身を適宜変更してください。引数も変更して大丈夫です。
func ascUsers() {
}

// 問題5 を実装してください。
// 中身を適宜変更してください。引数も変更して大丈夫です。
func descUsers() {
}

// func deleteUser(engine xorm.Engine) {
// 	user := User{}
// 	result, err := engine.Where("class=?", 1).Delete(&user)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	if result == 0 {
// 		log.Fatal("Not Found")
// 	}
// 	fmt.Printf("< deleteUser > ID:%d 名前:%s, Class:%d, 年齢:%d\n", user.ID, user.Name, user.Class, user.Age)
// 	fmt.Println()
// }

func Task1() {
	fmt.Println("homework012_1")

	// engine, err := xorm.NewEngine("mysql", "root:root@tcp([127.0.0.1]:3306)/sample_db?charset=utf8mb4&parseTime=true")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = engine.Sync2(new(User))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 各関数を実行します
	// insertUsers関数を実装してください。

	// ascUsers関数を実装してください。

	// descUsers関数を実装してください。

	// deleteUser(*engine)
}
