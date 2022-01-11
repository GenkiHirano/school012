package homework012

import (
	"fmt"
	"log"

	"xorm.io/xorm"
)

// 問題4
// Student構造体を実装してください。
type Student struct {
	ID    int    `xorm:"id"`
	Name  string `xorm:"name"`
	Class int    `xorm:"class"`
	Age   int    `xorm:"age"`
}

// 問題4 を実装してください。
// 中身を適宜変更してください。引数も変更して大丈夫です。
func insertStudent(engine xorm.Engine, students []Student) {
	// ここにinsertの処理を書く
	_, err := engine.Table("student").Insert(students)
	if err != nil {
		log.Fatal(err)
	}

	for _, student := range students {
		fmt.Printf("< insertStudent > ID:%d 名前:%s, Class:%d, 年齢:%d\n", student.ID, student.Name, student.Class, student.Age)
	}

	// fmt.Println("レコードの追加が完了しました")
	fmt.Print("------------------------------------------------------------\n")
}

// 問題4 を実装してください。
// 中身を適宜変更してください。引数も変更して大丈夫です。
func whereAndStudent(engine xorm.Engine) {
	student := []Student{}
	err := engine.Where("ID > ? && Age > ?", 1, 25).Find(&student)
	if err != nil {
		log.Println(err)
	}
	for _, student := range student {
		fmt.Printf("< whereAndStudent > ID:%d 名前:%s, Class:%d, 年齢:%d\n", student.ID, student.Name, student.Class, student.Age)
	}
	fmt.Print("------------------------------------------------------------\n")
}

// 問題4 を実装してください。
// 中身を適宜変更してください。引数も変更して大丈夫です。
func whereOrStudent(engine xorm.Engine) {
	student := []Student{}
	err := engine.Where("ID < ? || Age = ?", 2, 30).Find(&student)
	if err != nil {
		log.Println(err)
	}
	for _, student := range student {
		fmt.Printf("< whereOrStudent > ID:%d 名前:%s, Class:%d, 年齢:%d\n", student.ID, student.Name, student.Class, student.Age)
	}
	fmt.Print("------------------------------------------------------------\n")
}

func deleteStudent(engine xorm.Engine) {
	student := Student{}
	result, err := engine.Where("class=?", 1).Delete(&student)
	if err != nil {
		log.Println(err)
	}
	if result == 0 {
		log.Fatal("Not Found")
	}
	fmt.Printf("< deleteStudent > ID:%d 名前:%s, Class:%d, 年齢:%d\n", student.ID, student.Name, student.Class, student.Age)
	fmt.Println()
}

func Task2() {
	fmt.Println("homework012_2")

	engine, err := xorm.NewEngine("mysql", "root:root@tcp([127.0.0.1]:3306)/sample_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	err = engine.Sync2(new(Student))
	if err != nil {
		log.Fatal(err)
	}

	studentA := Student{
		ID:    1,
		Name:  "田中太郎",
		Class: 1,
		Age:   20,
	}

	studentB := Student{
		ID:    2,
		Name:  "鈴木次郎",
		Class: 1,
		Age:   25,
	}

	studentC := Student{
		ID:    3,
		Name:  "伊藤桃子",
		Class: 1,
		Age:   30,
	}

	students := []Student{studentA, studentB, studentC}

	// 各関数を実行します
	// InsertSum関数を実行してください。
	insertStudent(*engine, students)

	// whereAndStudent関数を実行してください。
	whereAndStudent(*engine)

	// whereOrStudent関数を実行してください。
	whereOrStudent(*engine)

	deleteStudent(*engine)
}
