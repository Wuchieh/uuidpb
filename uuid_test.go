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

	var uid2 *uuidpb.UUID
	if uid2.AsUUID() != uuid.Nil {
		t.Errorf("AsUUID returned a non-nil UUID")
	}
}

func TestUUID_AsNullUUID(t *testing.T) {
	uid := uuidpb.UUID{}
	if nuid := uid.AsNullUUID(); !nuid.Valid {
		t.Errorf("AsUUID returned a non-nil UUID")
	}

	var uid2 *uuidpb.UUID
	if nuid := uid2.AsNullUUID(); nuid.Valid {
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
