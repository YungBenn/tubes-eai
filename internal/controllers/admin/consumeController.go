package admin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type KotaResponse struct {
	Message string `json:"message"`
	Kota    []Kota `json:"kota"`
}

type Kota struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

func GetAllKota(c *fiber.Ctx) error {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://odd-pear-fish-vest.cyclic.app/kota", nil)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZW1haWwiOiJhZG1pbkBnbWFpbC5jb20iLCJpc0FkbWluIjp0cnVlLCJpYXQiOjE2ODY0NzI3MjIsImV4cCI6MTY4NjQ3NjMyMn0.VmtvxPKhgnYqmVp6FVTB20SnMN16Fqz2FDfeEJSKWaI"

	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer res.Body.Close()

    var response KotaResponse
    if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
        log.Fatal(err)
    }

    return c.Status(200).JSON(response)
}
