package handler

import (
	"fmt"
	"library_api/model"
	// "fmt"

	"github.com/gofiber/fiber/v2"
)

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
} // delChar()

func GetOneBook(c *fiber.Ctx) error { // 傳入書名 去圖書館找尋這本書的資料

	bookname := c.Query("name")
	book, err := model.GetOneBookByName(bookname)
	if err != nil {
		return err
	} // if()

	fmt.Println(book)
	fmt.Printf("%v", book.Tag)
	//book.Tag = str

	// return c.JSON(fiber.Map{"status": "success", "message": "Get All Book", "data": book})
	return c.JSON(book)
} // GetAllBook()

func GetBookByTag(c *fiber.Ctx) error {

	str := new(TagYesNo)
	err := c.BodyParser(str) // BodyParser can only parse into struct
	if err != nil {
		fmt.Println(err)
		return err
	} // if()

	book, err := model.GetBookByTag(str.Yes, str.No)
	if err != nil {
		fmt.Println(err)
		return err
	} // if()

	fmt.Println(book)
	// return c.JSON(fiber.Map{"data": str})
	return c.JSON(book)
} // GetBookByTag()

type ParseTag struct {
	Tag []string `json:"tag"`
} // ParseTag()

type TagYesNo struct {
	Yes []string `json:"yesTag"`
	No  []string `json:"noTag"`
} // TagYesNo()
