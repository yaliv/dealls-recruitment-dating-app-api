package access_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"

	"yaliv/dating-app-api/internal/handlers/access"
	"yaliv/dating-app-api/internal/handlers/access/accessform"
	"yaliv/dating-app-api/internal/helpers/testinghelper"
)

func TestLogin(t *testing.T) {
	testinghelper.CompleteSetup(t)

	app := fiber.New()

	app.Post("/", accessform.ParseLogin, access.Login)

	reqBody := map[string]any{
		"email":    "MimosaBurrows@jourrapide.com",
		"password": "Husoh0EeP",
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

	testinghelper.CheckHttpStatus(t, res.StatusCode, 200)
	testinghelper.CheckSuccess(t, resBody)

	testinghelper.CheckData(t, resBody, testinghelper.DataTests{
		"access_token": testinghelper.PropertyTest{Type: jsonparser.Object, Value: nil},
		"id":           testinghelper.PropertyTest{Type: jsonparser.Number, Value: "1"},
		"email":        testinghelper.PropertyTest{Type: jsonparser.String, Value: reqBody["email"]},
	})

	testinghelper.CheckData(t, resBody, testinghelper.DataTests{
		"value":      testinghelper.PropertyTest{Type: jsonparser.String, Value: nil},
		"expired_at": testinghelper.PropertyTest{Type: jsonparser.String, Value: nil},
	}, "data", "access_token")
}

func TestLoginInactive(t *testing.T) {
	testinghelper.CompleteSetup(t)

	app := fiber.New()

	app.Post("/", accessform.ParseLogin, access.Login)

	reqBody := map[string]any{
		"email":    "DiamandaHornblower@dayrep.com",
		"password": "iZ2mohghae",
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

	testinghelper.CheckHttpStatus(t, res.StatusCode, 401)
	testinghelper.CheckSuccess(t, resBody, false)

	testinghelper.CheckData(t, resBody, testinghelper.DataTests{
		"code":    testinghelper.PropertyTest{Type: jsonparser.String, Value: "ERR_LOGIN_CREDENTIALS"},
		"message": testinghelper.PropertyTest{Type: jsonparser.String, Value: "Invalid email and/or password, or inactive account."},
	}, "error")
}
