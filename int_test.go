package mdgoenum

import (
	"testing"
)

func TestIntEnum_Add(t *testing.T) {
	enum := NewIntEnum()
	enum.Add(1, "One")
	enum.Add(2, "Two")

	if enum.Len() != 2 {
		t.Errorf("Expected 2 members, got %d", enum.Len())
	}

	member, err := enum.Get(1)
	if err != nil {
		t.Errorf("Error getting member: %v", err)
	}

	if member.GetValue() != 1 || member.GetDesc() != "One" {
		t.Errorf("Member does not match expected value or description")
	}
}

func TestIntEnum_MustAdd(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected panic, got nil")
		}
	}()

	enum := NewIntEnum()
	enum.MustAdd(0, "Zero value should panic")
}

func TestIntEnum_Get(t *testing.T) {
	enum := NewIntEnum()
	enum.Add(1, "One")

	_, err := enum.Get(2)
	if err == nil {
		t.Error("Expected error for non-existent member, got nil")
	}
}

func TestIntEnum_MustGet(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, got nil")
		}
	}()

	enum := NewIntEnum()
	enum.MustGet(3)
}

func TestIntEnum_Contains(t *testing.T) {
	enum := NewIntEnum()
	enum.Add(1, "One")

	if !enum.Contains(1) {
		t.Error("Expected enum to contain '1'")
	}

	if enum.Contains(2) {
		t.Error("Expected enum to not contain '2'")
	}
}

func TestIntEnum_Len(t *testing.T) {
	enum := NewIntEnum()
	enum.Add(1, "One")

	if enum.Len() != 1 {
		t.Errorf("Expected 1 member, got %d", enum.Len())
	}
}

func TestIntEnum_IsEmpty(t *testing.T) {
	enum := NewIntEnum()

	if !enum.IsEmpty() {
		t.Error("Expected enum to be empty")
	}

	enum.Add(1, "One")

	if enum.IsEmpty() {
		t.Error("Expected enum to not be empty")
	}
}
