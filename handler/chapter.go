package handler

import (

	"library_api/model"
	"fmt"
	// "time"

	"github.com/gofiber/fiber/v2"

) // import()

func GetOneBookAllChapter( c *fiber.Ctx ) error {

	bookname := c.Query("name") // 給出要查出章節的書本名稱

	chapter, err := model.GetChapterByBookName( bookname )
	if err != nil {
		fmt.Println( err ) 
		return err 
	} // if()

	return c.JSON( chapter )
	
} // GetOneBookAllChapter()