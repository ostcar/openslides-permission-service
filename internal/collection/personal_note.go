package collection

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// PersonalNote handels permissions for personal notes.
type PersonalNote struct {
	dp dataprovider.DataProvider
}

// NewPersonalNote initializes a personal note.
func NewPersonalNote(dp dataprovider.DataProvider) *PersonalNote {
	return &PersonalNote{
		dp: dp,
	}
}

// Connect creates the routes.
func (p *PersonalNote) Connect(s perm.HandlerStore) {
	s.RegisterWriteHandler("personal_note.create", perm.WriteCheckerFunc(p.write))
	s.RegisterWriteHandler("personal_note.update", perm.WriteCheckerFunc(p.write))
	s.RegisterWriteHandler("personal_note.delete", perm.WriteCheckerFunc(p.write))

	s.RegisterReadHandler("personal_note", p)
}

func (p PersonalNote) write(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	if userID == 0 {
		perm.LogNotAllowedf("Anonymous can not create personal notes.")
		return false, nil
	}
	return true, nil
}

// RestrictFQFields checks for read permissions.
func (p PersonalNote) RestrictFQFields(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	var noteUserID int
	var lastID int
	for _, fqfield := range fqfields {
		if lastID != fqfield.ID {
			lastID = fqfield.ID
			key := fmt.Sprintf("personal_note/%d/user_id", fqfield.ID)
			if err := p.dp.Get(ctx, key, &noteUserID); err != nil {
				return fmt.Errorf("getting %s from datastore: %w", key, err)
			}
		}

		if noteUserID != userID {
			continue
		}

		result[fqfield.String()] = true
	}
	return nil
}
