package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Arrays
	ages := [3]int{23, 34, 23}
	fmt.Println(ages)

	ageSlices := []int{20, 66, 26, 39, 16, 99}
	sort.Ints(ageSlices)
	fmt.Println("After sort:", ageSlices)

	// Slices don't take array length
	names := []string{"derek", "kofi", "ama"}
	names[1] = "duafa"

	fmt.Println(names, len(names))

	names = append(names, "Emma")
	fmt.Println(names, len(names))

	rangeOne := names[1:3]
	fmt.Println(rangeOne, len(rangeOne))

	// strings from std
	message := "hello world in ghana"
	fmt.Println(strings.Contains(message, "hello"))
}


app := fiber.New()

	// GET /api/register
	app.Get("/api/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	})

	// GET /flights/LAX-SFO
	app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
		return c.SendString(msg) // => 💸 From: LAX, To: SFO
	})

	// GET /dictionary.txt
	app.Get("/:file.:ext", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
		return c.SendString(msg) // => 📃 dictionary.txt
	})

	// GET /john/75
	app.Get("/:name/:age/:gender?", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
		return c.SendString(msg) // => 👴 john is 75 years old
	})

	// GET /john
	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
		return c.SendString(msg) // => Hello john 👋!
	})

	log.Fatal(app.Listen(initializers.Config.ServerPort))