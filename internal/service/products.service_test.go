package service

import (
	context "context"
	"testing"

	models "github.com/disturb/max-inventory/internal/models"
)

func TestAddProduct(t *testing.T) {
	testCases := []struct {
		Name          string
		Product       models.Product
		Email         string
		ExpectedError error
	}{
		{
			Name: "AddProduct_Success",
			Product: models.Product{
				Name:        "Test Product",
				Description: "Test Description",
				Price:       10.00,
			},
			Email:         "admin@email.com",
			ExpectedError: nil,
		},
		{
			Name: "AddProduct_InvalidPermissions",
			Product: models.Product{
				Name:        "Test Product",
				Description: "Test Description",
				Price:       10.00,
			},
			Email:         "customer@email.com",
			ExpectedError: ErrInvalidPermissions,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.AddProdcut(ctx, tc.Product, tc.Email)

			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
