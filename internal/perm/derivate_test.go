package perm_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
	"github.com/OpenSlides/openslides-permission-service/internal/test"
)

func TestDerivatePerm(t *testing.T) {
	tdp := test.NewDataProvider()
	tdp.AddUser(1)
	tdp.AddUserToMeeting(1, 1)
	tdp.AddUserToGroup(1, 1, 2)
	dp := dataprovider.DataProvider{External: tdp}

	p, err := perm.New(context.Background(), dp, 1, 1)

	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}
	if !p.Has("motion.can_see") {
		t.Errorf("User does not have can_see permission")
	}
}
