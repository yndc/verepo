package project_test

import (
	"fmt"
	"testing"

	"github.com/flowscan/repomaster-go/pkg/project"
)

func TestGetAll(t *testing.T) {
	projects, err := project.GetAll()
	if err != nil {
		t.Error(err)
	}

	for _, v := range projects {
		fmt.Println(v)
	}
}
