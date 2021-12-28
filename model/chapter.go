package model

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Chapter struct {
	bun.BaseModel 					`bun:"table:chapter"`
	Uuid       			uuid.UUID 	`bun:"type:uuid,default:uuid_generate_v4()"`
	ID  	          	uint 		`bun:",pk"`
	Book    	      	string		`bun:"column:book"`
	ChapterNo			int 		`bun:"column:chapter_no"`
	BookId				int 		`bun:"column:book_id"`
} // Chapter()

func GetChapterByBookName( bookname string ) ( []Chapter, error ) {
	var chapter []Chapter
	err := Db.NewSelect().Model(&chapter).Where("book = ?", bookname).Scan(Ctx)
	return chapter, err
} // GetChapterByBookName()