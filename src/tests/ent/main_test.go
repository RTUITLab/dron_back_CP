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

	_, err = client.Answer.Create().SetAnswer("asdas").SetCorrect(true).Save(context.Background())
	if err != nil {
		t.Log(err)
	}
}