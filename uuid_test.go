package uuidpb_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/wuchieh/uuidpb"
)

func TestUUID_AsUUID(t *testing.T) {
	uid := uuidpb.UUID{}
	if uid.AsUUID() != uuid.Nil {
		t.Errorf("AsUUID returned a non-nil UUID")
	}
}

func TestNew(t *testing.T) {
	uid := uuid.New()
	uidpb := uuidpb.New(uid)
	if uid != uidpb.AsUUID() {
		t.Errorf("New uuid does not match")
	}
}
