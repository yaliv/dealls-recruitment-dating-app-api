package registration_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
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

	email1 := "MimosaBurrows@jourrapide.com"
	t.Log(email1, "- Email address is taken.")
	req := httptest.NewRequest("GET", "/"+email1, nil)

	res, _ := app.Test(req)
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	testinghelper.CheckHttpStatus(t, res.StatusCode, 200)
	testinghelper.CheckSuccess(t, resBody)

	testinghelper.CheckData(t, resBody, testinghelper.DataTests{
		"email":        testinghelper.PropertyTest{Type: jsonparser.String, Value: email1},
		"is_available": testinghelper.PropertyTest{Type: jsonparser.Boolean, Value: "false"},
	})

	email2 := "MyrtleHayward@jourrapide.com"
	t.Log(email2, "- Email address is available.")
	req = httptest.NewRequest("GET", "/"+email2, nil)

	res, _ = app.Test(req)
	defer res.Body.Close()

	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	testinghelper.CheckHttpStatus(t, res.StatusCode, 200)
	testinghelper.CheckSuccess(t, resBody)

	testinghelper.CheckData(t, resBody, testinghelper.DataTests{
		"email":        testinghelper.PropertyTest{Type: jsonparser.String, Value: email2},
		"is_available": testinghelper.PropertyTest{Type: jsonparser.Boolean, Value: "true"},
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
