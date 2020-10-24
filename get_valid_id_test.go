package main

import (
	"testing"
)

func TestGetValidID(t *testing.T) {
	// a valid DID should succeed
	validID := "OR9GoWA7txtKzeegiB52uVGjRWU00q_jE-eJC8ct1lE"
	validDID := "did:jlinc:OR9GoWA7txtKzeegiB52uVGjRWU00q_jE-eJC8ct1lE"
	id, ok := getValidID(validDID)
	if !ok {
		t.Errorf(validDID)
	}
	if id != validID {
		t.Errorf("returned incorrect id: got %v want %v", id, validID)
	}

	// an invalid DID should fail
	invalidDID := "did:notjlinc:OR9GoWA7txtKzeegiB52uVGjRWU00q_jE-eJC8ct1lE"
	badid, ok := getValidID(invalidDID)
	if ok {
		t.Errorf("succeeded where it should have failed: got %v want %v", ok, !ok)
	}
	if badid != "" {
		t.Errorf("returned incorrect id: got %v want %v", id, "")
	}
}
