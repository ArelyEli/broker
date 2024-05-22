package schemas

import (
	"testing"
)

func TestBusinessCreateRequest(t *testing.T) {
	t.Run("Name field set correctly", func(t *testing.T) {
		m := CreateBusinessRequest{Name: "Test Business"}
		if m.Name != "Test Business" {
			t.Errorf("Expected Name to be 'Test Business', got %v", m.Name)
		}
	})

	t.Run("Commission field set correctly", func(t *testing.T) {
		m := CreateBusinessRequest{Commission: 10.5}
		if m.Commission != 10.5 {
			t.Errorf("Expected Commission to be 10.5, got %v", m.Commission)
		}
	})

	t.Run("Commission field omitted correctly", func(t *testing.T) {
		m := CreateBusinessRequest{Name: "Test Business"}
		if m.Commission != 0 {
			t.Errorf("Expected Commission to be 0 (omitted), got %v", m.Commission)
		}
	})
}
