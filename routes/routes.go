package routes

import (
    "library_api/handler"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/session"
)

func SetupRoutes(router fiber.Router) {

    sess := session.New()

    book := router.Group("/book")

	// inject session to request context
	book.Use(func(c *fiber.Ctx) error {
		c.Locals("session", sess)
		return c.Next()
	})
    
    // Book
    book.Get("/", handler.GetOneBook )
    book.Post("/tag", handler.GetBookByTag )


    chapter := router.Group("/chapter")
    // inject session to request context
	chapter.Use(func(c *fiber.Ctx) error {
		c.Locals("session", sess)
		return c.Next()
	})

    // Chapter
    chapter.Get("/", handler.GetOneBookAllChapter )

} // SetupNoteRoutes()