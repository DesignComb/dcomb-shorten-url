package server

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"main/model"
	"main/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func redirect(c *fiber.Ctx) error {
	shortenUrl := c.Params("redirect")
	url, err := model.FindByUrl(shortenUrl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not find url in DB " + err.Error(),
		})
	}
	// grab any stats you want...
	url.Clicked += 1
	err = model.UpdateUrlShorten(url)
	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map {
		"origin": url.Origin,
	})
}

func getAllUrlShorten(c *fiber.Ctx) error {
	urls, err := model.GetAllUrlShorten()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error getting all goly links " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(urls)
}

func getUrlShorten(c *fiber.Ctx) error {
	id, err := strconv.ParseUint( c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error could not parse id " + err.Error(),
		})
	}

	url, err := model.GetUrlShorten(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error could not retrieve url from db " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(url)
}

func createUrlShorten(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var urlShorten model.UrlShorten
	err := c.BodyParser(&urlShorten)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error parsing JSON " + err.Error(),
		})
	}

	if len(urlShorten.Origin) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map {
			"message": "origin column is required",
		})
	}

	if urlShorten.Random {
		urlShorten.Short = utils.RandomURL(8)
	}else {
		if len(urlShorten.Short) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map {
				"message": "short column is required",
			})
		}
	}

	urlShorten, err = model.CreateUrlShorten(urlShorten)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not create urlShorten in db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(urlShorten)
}

func updateUrlShorten(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var urlShorten model.UrlShorten

	err := c.BodyParser(&urlShorten)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not parse json " + err.Error(),
		})
	}

	err = model.UpdateUrlShorten(urlShorten)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not update urlShorten link in DB " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(urlShorten)
}

func deleteUrlShorten(c *fiber.Ctx) error {
	id, err := strconv.ParseUint( c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not parse id from url " + err.Error(),
		})
	}

	err = model.DeleteUrlShorten(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not delete from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map {
		"message": "urlShorten deleted.",
	})
}


func SetupAndListen() {

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "http://172.19.0.1, http://54.249.0.5",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Use(logger.New())
	router.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	router.Get("/r/:redirect", redirect)

	//router.Get("/urlShorten", getAllUrlShorten)
	router.Get("/urlShorten/:id", getUrlShorten)
	router.Post("/urlShorten", createUrlShorten)
	//router.Patch("/urlShorten", updateUrlShorten)
	//router.Delete("/urlShorten/:id", deleteUrlShorten)


	data, _ := json.MarshalIndent(router.Stack(), "", "  ")
	fmt.Println(string(data))

	router.Listen(":5000")

}