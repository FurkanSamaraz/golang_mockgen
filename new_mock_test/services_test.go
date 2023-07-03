package main

import (
	"testing"

	"main/services"

	mocks "main/mocks/repository"

	"github.com/golang/mock/gomock"
)

func TestUserService_AddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockUserRepository(ctrl)
	mockService := services.NewUserService(mockRepository)

	// Mock Repository'nin AddUser metodunu istediğiniz şekilde ayarlayın
	mockRepository.EXPECT().AddUser("username").Return(nil)

	// Servis metodu çağırın
	err := mockService.AddUser("username")

	// Hata kontrolü yapın
	if err != nil {
		t.Errorf("AddUser returned an unexpected error: %v", err)
	}
}

func TestUserService_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockUserRepository(ctrl)
	mockService := services.NewUserService(mockRepository)

	// Mock Repository'nin GetUserByID metodunu istediğiniz şekilde ayarlayın
	mockRepository.EXPECT().GetUserByID(1).Return("username", nil)

	// Servis metodu çağırın
	username, err := mockService.GetUserByID(1)

	// Hata veya sonuç kontrolü yapın
	if err != nil {
		t.Errorf("GetUserByID returned an unexpected error: %v", err)
	}
	if username != "username" {
		t.Errorf("GetUserByID returned an unexpected username: got %s, want %s", username, "username")
	}
}
