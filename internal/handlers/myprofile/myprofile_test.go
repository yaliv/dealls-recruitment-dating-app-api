package myprofile_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"

	"yaliv/dating-app-api/internal/handlers/authorization"
	"yaliv/dating-app-api/internal/handlers/myprofile"
	"yaliv/dating-app-api/internal/helpers/testinghelper"
)

func TestShow(t *testing.T) {
	testinghelper.CompleteSetup(t)

	app := fiber.New()

	app.Use(authorization.New())
	app.Get("/", myprofile.Show)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", testinghelper.GetAuthorization(t, 1))

	res, _ := app.Test(req)
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	testinghelper.CheckHttpStatus(t, res.StatusCode, 200)
	testinghelper.CheckSuccess(t, resBody)

	testinghelper.CheckData(t, resBody, testinghelper.DataTests{
		"id":       testinghelper.PropertyTest{Type: jsonparser.Number, Value: "1"},
		"usser_id": testinghelper.PropertyTest{Type: jsonparser.Number, Value: "1"},
		"verified": testinghelper.PropertyTest{Type: jsonparser.Boolean, Value: "false"},
		"name":     testinghelper.PropertyTest{Type: jsonparser.String, Value: "Mimosa Burrows"},
		"age":      testinghelper.PropertyTest{Type: jsonparser.Number, Value: "27"},
		"bio":      testinghelper.PropertyTest{Type: jsonparser.String, Value: "I am not a player...I'm the game"},
		"pic_url":  testinghelper.PropertyTest{Type: jsonparser.String, Value: "https://picsum.photos/id/10/400/600"},
	})
}
