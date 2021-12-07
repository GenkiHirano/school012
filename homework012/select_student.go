package homework012

import (
	"fmt"
)

// 問題4
// Student構造体を実装してください。

// 問題4 を実装してください。
// 中身を適宜変更してください。引数も変更して大丈夫です。
func insertStudent() {
}

// 問題4 を実装してください。
// 中身を適宜変更してください。引数も変更して大丈夫です。
func whereAndStudent() {
}

// 問題4 を実装してください。
// 中身を適宜変更してください。引数も変更して大丈夫です。
func whereOrStudent() {
}

// func deleteStudent(engine xorm.Engine) {
// 	student := Student{}
// 	result, err := engine.Where("class=?", 1).Delete(&student)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	if result == 0 {
// 		log.Fatal("Not Found")
// 	}
// 	fmt.Printf("< deleteStudent > ID:%d 名前:%s, Class:%d, 年齢:%d\n", student.ID, student.Name, student.Class, student.Age)
// 	fmt.Println()
// }

func Task2() {
	fmt.Println("homework012_2")

	// engine, err := xorm.NewEngine("mysql", "root:root@tcp([127.0.0.1]:3306)/sample_db?charset=utf8mb4&parseTime=true")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = engine.Sync2(new(Student))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// studentA := Student{
	// 	ID:    1,
	// 	Name:  "田中太郎",
	// 	Class: 1,
	// 	Age:   20,
	// }

	// studentB := Student{
	// 	ID:    2,
	// 	Name:  "鈴木次郎",
	// 	Class: 1,
	// 	Age:   25,
	// }

	// studentC := Student{
	// 	ID:    3,
	// 	Name:  "伊藤桃子",
	// 	Class: 1,
	// 	Age:   30,
	// }

	// students := []Student{studentA, studentB, studentC}

	// 各関数を実行します
	// InsertSum関数を実行してください。

	// whereAndStudent関数を実行してください。

	// whereOrStudent関数を実行してください。

	// deleteStudent(*engine)
}
