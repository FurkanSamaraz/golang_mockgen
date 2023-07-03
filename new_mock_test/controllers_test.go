package main

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"main/services"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	mocks "main/mocks/repository"

	"github.com/golang/mock/gomock"
)

func TestUserController_AddUser_Success(t *testing.T) {
	// Mock UserService nesnesi oluştur
	mockUserService := &MockUserService{}

	// UserController'ı mock nesnesi ile oluştur
	controller := NewUserController(mockUserService)

	// Fiber Ctx'i oluştur
	ctx := &fiber.Ctx{
		Request:  httptest.NewRequest(http.MethodPost, "/users", nil),
		Response: &fiber.Response{},
	}

	// UserService AddUser metoduna mock dönüş değerini ayarla
	mockUserService.EXPECT().AddUser("John").Return(nil)

	// UserController'ın AddUser metodunu çağır
	err := controller.AddUser(ctx)

	// Hata olmamalı
	assert.Nil(t, err)

	// Yanıtın başarılı olduğunu kontrol et
	assert.Equal(t, http.StatusOK, ctx.Response.StatusCode())
}

func TestUserController_AddUser_Failure(t *testing.T) {
	// Mock UserService nesnesi oluştur
	mockUserService := &MockUserService{}

	// UserController'ı mock nesnesi ile oluştur
	controller := NewUserController(mockUserService)

	// Fiber Ctx'i oluştur
	ctx := &fiber.Ctx{
		Request:  httptest.NewRequest(http.MethodPost, "/users", nil),
		Response: &fiber.Response{},
	}

	// UserService AddUser metoduna mock hata dönüş değerini ayarla
	mockUserService.EXPECT().AddUser("John").Return(errors.New("error"))

	// UserController'ın AddUser metodunu çağır
	err := controller.AddUser(ctx)

	// Hata olmalı
	assert.NotNil(t, err)

	// Yanıtın başarısız olduğunu kontrol et
	assert.Equal(t, http.StatusInternalServerError, ctx.Response.StatusCode())
}

func TestUserController_GetUser_Success(t *testing.T) {
	// Mock UserService nesnesi oluştur
	mockUserService := &MockUserService{}

	// UserController'ı mock nesnesi ile oluştur
	controller := NewUserController(mockUserService)

	// Fiber Ctx'i oluştur
	ctx := &fiber.Ctx{
		Request: httptest.NewRequest(http.MethodGet, "/users/1", nil),
		Params:  map[string]string{"id": "1"},
	}

	// UserService GetUserByID metoduna mock dönüş değerini ayarla
	mockUserService.EXPECT().GetUserByID(1).Return("John", nil)

	// UserController'ın GetUser metodunu çağır
	err := controller.GetUser(ctx)

	// Hata olmamalı
	assert.Nil(t, err)

	// Yanıtın başarılı olduğunu kontrol et
	assert.Equal(t, http.StatusOK, ctx.Response.StatusCode())

	// Yanıtın doğru olduğunu kontrol et
	expectedResponse := `{"user":"John"}`
	assert.Equal(t, expectedResponse, ctx.Response.Body())
}

func TestUserController_GetUser_Failure(t *testing.T) {
	// Mock UserService nesnesi oluştur
	mockUserService := &MockUserService{}

	// UserController'ı mock nesnesi ile oluştur
	controller := NewUserController(mockUserService)

	// Fiber Ctx'i oluştur
	ctx := &fiber.Ctx{
		Request: httptest.NewRequest(http.MethodGet, "/users/1", nil),
		Params:  map[string]string{"id": "1"},
	}

	// UserService GetUserByID metoduna mock hata dönüş değerini ayarla
	mockUserService.EXPECT().GetUserByID(1).Return("", errors.New("error"))

	// UserController'ın GetUser metodunu çağır
	err := controller.GetUser(ctx)

	// Hata olmalı
	assert.NotNil(t, err)

	// Yanıtın başarısız olduğunu kontrol et
	assert.Equal(t, http.StatusInternalServerError, ctx.Response.StatusCode())
}
