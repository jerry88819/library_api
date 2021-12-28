package model

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	//"github.com/uptrace/bun/dialect/pgdialect"
)

type Book struct {
	bun.BaseModel 								`bun:"table:book"`
	Uuid     		 	 uuid.UUID 				`bun:"type:uuid,default:uuid_generate_v4()"`
	ID          		 uint 					`bun:",pk"` // primary key
	Name         		 string
	Tag          		 []string 				`bun:"column:tag,array"` // `bun:"column:tag",",array"`  *[]string
} // Book()

func GetOneBookByName(bookname string) (Book, error) {

	var book Book
	err := Db.NewSelect().Model(&book).Where("name = ?", bookname).Scan(Ctx)
	return book, err

} // GetAllBook()

func GetBookByTag(tagY []string, tagN []string) ([]Book, error) {
	var book []Book
	include := pgdialect.Array(tagY) // 這句太屌了
	exclude := pgdialect.Array(tagN)

	if len(tagY) == 0 && len(tagN) == 0 {
		err := Db.NewSelect().Model(&book).Scan(Ctx)
		fmt.Println("1")
		return book, err
	} else if len(tagY) == 0 {
		err := Db.NewSelect().Model(&book).
			Where("not tag @> ?", exclude ).
			Scan(Ctx)

		fmt.Println("2")
		return book, err
	} else if len(tagN) == 0 {

		err := Db.NewSelect().Model(&book).
			Where("tag @> ?", include ).
			Scan(Ctx)

		fmt.Println("3")
		return book, err
	} else {
		// temp := "tag @> array ['1']::VARCHAR[]"
		err := Db.NewSelect().Model(&book).
			Where("tag @> ?", include ).
			Where("not tag @> ?", exclude ).
			Scan(Ctx)

		fmt.Println("4")
		return book, err
	} // else()
} // GetBookByTag()
