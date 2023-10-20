package utilities

import (
	"github.com/google/uuid"
	"testing"
)

func TestGenerateUUID_ValidUUID(t *testing.T) {
	generatedUUID := GenerateUUID()
	_, err := uuid.Parse(generatedUUID)
	if err != nil {
		t.Errorf("Expected a valid UUID, but got an error: %v", err)
	}
}

func TestGenerateUUID_Length(t *testing.T) {
	generatedUUID := GenerateUUID()
	if len(generatedUUID) != 36 {
		t.Errorf("Expected UUID length to be 36, but got: %d", len(generatedUUID))
	}
}
