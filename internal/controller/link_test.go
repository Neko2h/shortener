package controller

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Neko2h/shortener/internal/entity"
	"github.com/Neko2h/shortener/internal/usecase/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// go test -coverprofile coverage.out ./...
// go tool cover -html coverage.out

/*


func NewControllerManager(ctx context.Context, usecase *usecase.UsecasesManager) *ControllersManager {
	return &ControllersManager{
		LinkController: *NewLinkController(ctx, usecase),
	}
}
*/

func TestCreateLink(t *testing.T) {

	testLink := &entity.Link{
		Origin: "https://google.com",
	}

	tests := []struct {
		testName     string
		expectations func(ctx context.Context, uc *mocks.LinkUsecase)
		input        string
		route        string
		err          error
		code         int
		param        string
	}{
		{
			testName:     "create empty struct",
			expectations: func(ctx context.Context, uc *mocks.LinkUsecase) {},
			input:        `{}`,
			route:        "create",
			code:         http.StatusBadRequest,
		},

		{
			testName:     "create bad binding",
			expectations: func(ctx context.Context, uc *mocks.LinkUsecase) {},
			input:        `{some"}`,
			route:        "create",
			code:         http.StatusBadRequest,
		},

		{
			testName: "create usecase error",
			expectations: func(ctx context.Context, uc *mocks.LinkUsecase) {
				uc.On("New", ctx, testLink).Return(nil, errors.New("some error"))
			},
			input: `{"origin":"https://google.com"}`,
			route: "create",
			code:  http.StatusInternalServerError,
		},

		{
			testName: "create valid",
			expectations: func(ctx context.Context, uc *mocks.LinkUsecase) {
				uc.On("New", ctx, testLink).Return(testLink, nil)
			},
			input: `{"origin":"https://google.com"}`,
			route: "create",
			code:  http.StatusCreated,
		},

		{
			testName: "get missing param",
			expectations: func(ctx context.Context, uc *mocks.LinkUsecase) {
			},
			input: "",
			route: "get",
			code:  http.StatusBadRequest,
		},

		{
			testName: "get usecase internal error",
			expectations: func(ctx context.Context, uc *mocks.LinkUsecase) {
				uc.On("Get", ctx, mock.AnythingOfType("string")).Return(nil, errors.New("usecase error"))
			},
			input: string(mock.AnythingOfType("string")),
			route: "get",
			code:  http.StatusBadRequest,
		},

		{
			testName: "get not found",
			expectations: func(ctx context.Context, uc *mocks.LinkUsecase) {
				uc.On("Get", ctx, mock.AnythingOfType("string")).Return(nil, nil)
			},
			input: string(mock.AnythingOfType("string")),
			route: "get",
			code:  http.StatusNotFound,
		},

		{
			testName: "get not found",
			expectations: func(ctx context.Context, uc *mocks.LinkUsecase) {
				uc.On("Get", ctx, mock.AnythingOfType("string")).Return(nil, nil)
			},
			input: string(mock.AnythingOfType("string")),
			route: "get",
			code:  http.StatusNotFound,
		},

		{
			testName: "found, redirect",
			expectations: func(ctx context.Context, uc *mocks.LinkUsecase) {
				uc.On("Get", ctx, mock.AnythingOfType("string")).Return(testLink, nil)
			},
			input: string(mock.AnythingOfType("string")),
			route: "get",
			code:  http.StatusMovedPermanently,
		},
	}

	for _, test := range tests {
		t.Logf("running %v", test.testName)
		ctx := context.Background()
		uc := &mocks.LinkUsecase{}

		controller := &LinkController{uc}

		test.expectations(ctx, uc)

		if test.route == "create" {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test.input))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := controller.Create(c)
			if err != nil {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.code, rec.Code)

		} else if test.route == "get" {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetPath("/:hash")
			c.SetParamNames("hash")
			c.SetParamValues(test.input)

			err := controller.GetLink(c)
			if err != nil {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.code, rec.Code)
		}

		uc.AssertExpectations(t)
	}
}
