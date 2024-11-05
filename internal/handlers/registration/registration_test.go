package registration_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"

	"yaliv/dating-app-api/internal/handlers/registration"
	"yaliv/dating-app-api/internal/handlers/registration/registrationform"
	"yaliv/dating-app-api/internal/helpers/testinghelper"
)

func TestUserStatus(t *testing.T) {
	testinghelper.CompleteSetup(t)

	app := fiber.New()

	app.Get("/:email", registration.UserStatus)

	t.Run("Email address is taken", func(t *testing.T) {
		tUserStatus(t, app, "MimosaBurrows@jourrapide.com", false)
	})
	t.Run("Email address is available", func(t *testing.T) {
		tUserStatus(t, app, "MyrtleHayward@jourrapide.com", true)
	})
}

func tUserStatus(t *testing.T, app *fiber.App, email string, isAvailable bool) {
	req := httptest.NewRequest("GET", "/"+email, nil)

	res, _ := app.Test(req)
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	testinghelper.CheckHttpStatus(t, res.StatusCode, 200)
	testinghelper.CheckSuccess(t, resBody)

	testinghelper.CheckData(t, resBody, testinghelper.DataTests{
		"email":        testinghelper.PropertyTest{Type: jsonparser.String, Value: email},
		"is_available": testinghelper.PropertyTest{Type: jsonparser.Boolean, Value: strconv.FormatBool(isAvailable)},
	})
}

func TestRegister(t *testing.T) {
	testinghelper.MainSetup(t)
	testinghelper.ClearData()

	app := fiber.New()

	app.Post("/", registrationform.ParseRegister, registration.Register)

	reqBody := map[string]any{
		"email":    "MyrtleHayward@jourrapide.com",
		"password": "Eigh6Ufatai",
	}
	reqBodyJson := new(bytes.Buffer)
	json.NewEncoder(reqBodyJson).Encode(reqBody)

	req := httptest.NewRequest("POST", "/", reqBodyJson)
	req.Header.Set("Content-Type", "application/json")

	res, _ := app.Test(req)
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	testinghelper.CheckHttpStatus(t, res.StatusCode, 201)
	testinghelper.CheckSuccess(t, resBody)

	testinghelper.CheckData(t, resBody, testinghelper.DataTests{
		"id":    testinghelper.PropertyTest{Type: jsonparser.Number, Value: "1"},
		"email": testinghelper.PropertyTest{Type: jsonparser.String, Value: reqBody["email"]},
	})
}
