package main

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Aviator-Coding/HttpPLC/configs"
	"github.com/Aviator-Coding/HttpPLC/database"
	"github.com/Aviator-Coding/HttpPLC/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type HttpTest struct {
	description   string
	route         string // input route
	method        string // input method
	tokenString   string // input token
	body          string
	expectedError bool
	expectedCode  int
}

func TestPrivateRoutess(t *testing.T) {

	// Define a structure for specifying input and output data of a single test case.
	tests := []HttpTest{
		// {
		// 	description:   "Add User without Name",
		// 	route:         "/user",
		// 	method:        "POST",
		// 	tokenString:   "",
		// 	body:          `{"password":"lol","email": "info@aviator-coding.de"}`,
		// 	expectedError: false,
		// 	expectedCode:  403,
		// },
		// {
		// 	description:   "Add User without Password",
		// 	route:         "/user",
		// 	method:        "POST",
		// 	tokenString:   "",
		// 	body:          `{"name": "Aviator","email": "info@aviator-coding.de"}`,
		// 	expectedError: false,
		// 	expectedCode:  403,
		// },
		// {
		// 	description:   "Add User without Email",
		// 	route:         "/user",
		// 	method:        "POST",
		// 	tokenString:   "",
		// 	body:          `{"name": "Aviator","password":"lol"}`,
		// 	expectedError: true,
		// 	expectedCode:  403,
		// },
		{
			description:   "Add User without Email",
			route:         "/user",
			method:        "POST",
			tokenString:   "",
			body:          `{"name": "Aviator","password":"lol","email": "info@aviator-coding.de"}`,
			expectedError: false,
			expectedCode:  201,
		},
	}

	//Load Env Files
	configs.LoadConfig()
	// Connect to DB
	database.ConnectDB()

	// Define a new Fiber app.
	app := fiber.New()

	// Define routes.
	routes.UserPublicRoute(app)

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case.
		req := httptest.NewRequest(test.method, test.route, strings.NewReader(test.body))
		req.Header.Set("Authorization", test.tokenString)
		req.Header.Set("Content-Type", "application/json")

		// Perform the request plain with the app.
		resp, err := app.Test(req, -1) // the -1 disables request latency

		// Verify, that no error occurred, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses,
		// the next test case needs to be processed.
		if test.expectedError {
			continue
		}

		// Verify, if the status code is as expected.
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
