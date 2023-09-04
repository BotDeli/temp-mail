package tests

import (
	"context"
	"errors"
	"github.com/go-redis/redismock/v9"
	"temp-mail/internal/storage/mails"
	"temp-mail/pkg/generators"
	"testing"
)

var (
	testError = errors.New("tested error")
	testCtx   = context.Background()
)

func TestGetEmptyUUID(t *testing.T) {
	testKey := ""
	testingGetResponseError(t, testKey)
}

func testingGetResponseError(t *testing.T, testKey string) {
	client, mock := redismock.NewClientMock()
	mock.ExpectGet(testKey).SetErr(testError)

	s := mails.RedisMails{Ctx: testCtx, Client: client}
	mail, err := s.GetMail(testKey)

	if mail != "" {
		t.Errorf("expected empty mail, got %s", mail)
	}

	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestGetDontSetKeys(t *testing.T) {
	testKeys := []string{
		"test",
		"testK2",
		"dontCreate!",
		"dontAddedKey",
		"error!!!",
	}
	for _, testKey := range testKeys {
		testingGetResponseError(t, testKey)
	}
}

func TestGetDontSetUUIDs(t *testing.T) {
	var testUUID string
	for i := 0; i < 10; i++ {
		testUUID = generators.NewUUID()
		testingGetResponseError(t, testUUID)
	}
}
