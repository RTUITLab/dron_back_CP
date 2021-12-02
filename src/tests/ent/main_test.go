package enttests_test

import (
	"context"
	"testing"

	"github.com/0B1t322/CP-Rosseti-Back/db"
)

func TestFunc_Init(t *testing.T) {
	client, err := db.CreateClient("root:root@/site")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Log(err)
		t.FailNow()
	}

	if _, err := client.Role.Create().SetRole("student").Save(context.Background()); err != nil {
		t.Log(err)
	}

	// 
	// 

}