package mdgoenum

import (
	"testing"
)

func TestStringEnum_Add(t *testing.T) {
	enum := NewStringEnum()
	enum.Add("Active", "Represents an active state")
	enum.Add("Inactive", "Represents an inactive state")

	if enum.Len() != 2 {
		t.Errorf("Expected 2 members, got %d", enum.Len())
	}

	member, err := enum.Get("Active")
	if err != nil {
		t.Errorf("Error getting member: %v", err)
	}

	if member.GetValue() != "Active" || member.GetDesc() != "Represents an active state" {
		t.Errorf("Member does not match expected value or description")
	}
}

func TestStringEnum_MustAdd(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected panic, got nil")
		}
	}()

	enum := NewStringEnum()
	enum.MustAdd("", "Empty value should panic")
}

func TestStringEnum_Get(t *testing.T) {
	enum := NewStringEnum()
	enum.Add("Active", "Represents an active state")

	_, err := enum.Get("Inactive")
	if err == nil {
		t.Error("Expected error for non-existent member, got nil")
	}
}

func TestStringEnum_MustGet(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, got nil")
		}
	}()

	enum := NewStringEnum()
	enum.MustGet("NonExistent")
}

func TestStringEnum_Contains(t *testing.T) {
	enum := NewStringEnum()
	enum.Add("Active", "Represents an active state")

	if !enum.Contains("Active") {
		t.Error("Expected enum to contain 'Active'")
	}

	if enum.Contains("Inactive") {
		t.Error("Expected enum to not contain 'Inactive'")
	}
}

func TestStringEnum_Len(t *testing.T) {
	enum := NewStringEnum()
	enum.Add("Active", "Represents an active state")

	if enum.Len() != 1 {
		t.Errorf("Expected 1 member, got %d", enum.Len())
	}
}

func TestStringEnum_IsEmpty(t *testing.T) {
	enum := NewStringEnum()

	if !enum.IsEmpty() {
		t.Error("Expected enum to be empty")
	}

	enum.Add("Active", "Represents an active state")

	if enum.IsEmpty() {
		t.Error("Expected enum to not be empty")
	}
}
