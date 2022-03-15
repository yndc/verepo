package app_test

import (
	"fmt"
	"testing"

	"github.com/flowscan/repomaster-go/pkg/app"
)

func TestGetAll(t *testing.T) {
	apps, err := app.GetAll()
	if err != nil {
		t.Error(err)
	}

	for _, v := range apps {
		fmt.Println(v)
	}
}
