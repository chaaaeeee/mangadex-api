package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func getMangaTitlesByTitle(c *fiber.Ctx) error {
	var search Search

	if err := c.BodyParser(&search); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	params := url.Values{}
	params.Set("title", search.Title)

	mangaUrl := baseUrl + "/manga?" + params.Encode()

	response, err := http.Get(mangaUrl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var mangaData MangaResponse

	if err := json.Unmarshal(responseData, &mangaData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(mangaData)
}
