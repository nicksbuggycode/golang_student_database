package lookup

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strings"
)

type Student struct {
	gorm.Model
	Firstname string
	Lastname  string
	Email     string
	Class     int
}

func ConnectToDBFindAll(s string) *gorm.DB {
	db, err := gorm.Open("sqlite3", "/Users/bzzz/go/github.com/nicksbuggycode/orm/studz.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	var students []Student
	db.Find(&students)
	for _, stud := range students {
		fmt.Println(stud)
	}
	return nil
}

func ConnectToDBFindSome(s string) *gorm.DB {
	db, err := gorm.Open("sqlite3", "/Users/bzzz/go/github.com/nicksbuggycode/orm/studz.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	var students []Student
	//Using a string builder to compose a % string %
	builder := strings.Builder{}
	builder.WriteString("%")
	builder.WriteString(s)
	builder.WriteString("%")
	db.Where("firstname LIKE ?", builder.String()).Find(&students)
	for _, stud := range students {
		if stud.Class == 18 {
			fmt.Println(stud.Firstname, stud.Lastname, "is in class ", stud.Class)
		} else {
			fmt.Println(stud.Firstname, stud.Lastname, "is in class 17")
		}
	}
	if len(students) == 0 {
		db.Where("lastname LIKE ?", builder.String()).Find(&students)
		for _, stud := range students {
			if stud.Class == 18 {
				fmt.Println(stud.Firstname, stud.Lastname, "is in class ", stud.Class)
			} else {
				fmt.Println(stud.Firstname, stud.Lastname, "is in class 17")
			}
		}
	}
	return nil
}
