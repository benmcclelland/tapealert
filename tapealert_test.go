package tapealert_test

import (
	"testing"

	"github.com/benmcclelland/tapealert"
)

func TestAlert(t *testing.T) {
	ta, err := tapealert.New()
	if err != nil {
		t.Fatal(err)
	}
	alerts := ta.FromFlags(0x1, 0)
	if len(alerts) != 1 {
		t.Errorf("expected 1 alert, got %d", len(alerts))
	}
	if alerts[0].Name != "Read Warning" {
		t.Errorf("expected Read Warning, got %s", alerts[0].Name)
	}
	alerts = ta.FromFlags(0x1, 2)
	if len(alerts) != 0 {
		t.Errorf("expected 0 alerts, got %d", len(alerts))
	}
}
