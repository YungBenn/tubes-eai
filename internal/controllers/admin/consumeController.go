package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BundleResponse struct {
	Message string   `json:"message"`
	Bundle  []Bundle `json:"kota"`
}

type Bundle struct {
	ID               int    `json:"id"`
	User_id          int    `json:"user_id"`
	Nama_user        string `json:"nama_user"`
	Tiket_id         int    `json:"tiket_id"`
	Destination_name string `json:"destination_name"`
	Quantity         int    `json:"quantity"`
	Kota             string `json:"kota"`
	Hotel_id         string `json:"hotel_id"`
	Hotel_name       string `json:"hotel_name"`
	Nopol            string `json:"noPol"`
	KursiTersedia    int    `json:"kursiTersedia"`
}

func GetAllBundle(c *fiber.Ctx) error {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://tourism-api-production.up.railway.app/booking_bundle", nil)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer res.Body.Close()

	var response []Bundle
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatal(err)
	}

	return c.Status(200).JSON(response)
}

func GetBundleByID(c *fiber.Ctx) error {
	id := c.Params("id")
	bundleID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	client := http.Client{}

	url := fmt.Sprintf("https://tourism-api-production.up.railway.app/booking_bundle/%d", bundleID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer res.Body.Close()

	var response Bundle
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatal(err)
	}

	return c.Status(200).JSON(response)
}

type BookingRequest struct {
	User_id          int    `json:"user_id"`
	Nama_user        string `json:"nama_user"`
	Tiket_id         int    `json:"tiket_id"`
	Destination_name string `json:"destination_name"`
	Quantity         int    `json:"quantity"`
	Kota             string `json:"kota"`
	Hotel_id         string `json:"hotel_id"`
	Hotel_name       string `json:"hotel_name"`
	Nopol            string `json:"noPol"`
	KursiTersedia    int    `json:"kursiTersedia"`
}

func PutBookBundle(c *fiber.Ctx) error {
	id := c.Params("id")
	bundleID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	reqbody := BookingRequest{}

	if err := c.BodyParser(&reqbody); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	bookingReq := BookingRequest{
		User_id:          reqbody.User_id,
		Nama_user:        reqbody.Nama_user,
		Tiket_id:         reqbody.Tiket_id,
		Destination_name: reqbody.Destination_name,
		Quantity:         reqbody.Quantity,
		Kota:             reqbody.Kota,
		Hotel_id:         reqbody.Hotel_id,
		Hotel_name:       reqbody.Hotel_name,
		Nopol:            reqbody.Nopol,
		KursiTersedia:    reqbody.KursiTersedia,
	}

	// Mengonversi data permintaan ke JSON
	jsonReq, err := json.Marshal(bookingReq)
	if err != nil {
		log.Println("Error marshaling request:", err)
		return err
	}

	// Membuat permintaan PUT
	url := fmt.Sprintf("https://tourism-api-production.up.railway.app/booking_bundle/%d", bundleID)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Println("Error creating request:", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	// Membuat klien HTTP dan mengirim permintaan
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	// Memeriksa kode status respons
	if resp.StatusCode != http.StatusOK {
		log.Println("Request failed with status code:", resp.StatusCode)
		return err
	}

	// Membaca respons body
	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		log.Println("Error decoding response:", err)
		return err
	}

	// Menampilkan respons
	fmt.Println("Response:", responseBody)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Book bundle updated",
	})
}
