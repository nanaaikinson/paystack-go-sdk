package paystack_test

import (
	"fmt"
	"os"
	"testing"

	"context"

	_ "github.com/joho/godotenv/autoload"
	p "github.com/nanaaikinson/paystack-sdk-go"
	"github.com/nanaaikinson/paystack-sdk-go/models"
)

func TestApplePayClient(t *testing.T) {
	secretKey := loadSecretKey(t)
	client := p.NewApplePayClient(p.WithSecretKey(secretKey))
	var r models.Response[any]
	if err := client.All(context.TODO(), &r); err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v", r)
}

func loadSecretKey(t *testing.T) string {
	t.Helper()
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if secretKey == "" {
		t.Fatal("PAYSTACK_SECRET_KEY in not provided")
	}
	return secretKey
}
