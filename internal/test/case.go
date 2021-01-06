package test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"

	models "github.com/OpenSlides/openslides-models-to-go"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
	"github.com/OpenSlides/openslides-permission-service/pkg/permission"
	"gopkg.in/yaml.v3"
)

// Case tests a request
type Case struct {
	name string
	Data map[string]json.RawMessage `yaml:"data"`

	UserID     int
	MeetingID  int
	Permission string
	Handler    string

	// only needed on action.
	Payload         map[string]string
	ExpectedAllowed bool `yaml:"is_allowed"`

	// only needed on collections.
	FQFields         []string
	ExpectedFQFields []string
}

// Test runs the test.Case.
//
// It runs a request with the given userID to the handler and makes sure the
// result is the same then expectedAllowed.
//
// If the userID is not set, the user with the ID 1337 is used. If the meetingID
// does not exist, the meetingID 1 is used.
//
// A group with the ID 1337 is created with the given permissions and the user
// is put into this group.
//
// If the user is in other groups, he also has the permissions from this other
// groups. To test, that a user is not allowed to do something, use a userID
// that is not in the exampleData.
func (c *Case) Test(t *testing.T) {
	data := make(map[string]json.RawMessage, len(c.Data)+len(models.ExampleData))
	for k, v := range models.ExampleData {
		data[k] = v
	}
	for k, v := range c.Data {
		data[k] = v
	}

	c.UserID = defaultInt(c.UserID, 1337)
	meetingID := defaultInt(c.MeetingID, 1)

	// Make sure the user does exists.
	userFQID := fmt.Sprintf("user/%d", c.UserID)
	if data[userFQID+"/id"] == nil {
		data[userFQID+"/id"] = []byte(strconv.Itoa(c.UserID))
	}

	// Make sure, the user is in the meeting.
	meetingFQID := fmt.Sprintf("meeting/%d", meetingID)
	data[meetingFQID+"/user_ids"] = jsonAddInt(data[meetingFQID+"/user_ids"], c.UserID)

	// Create group with the user and the given permissions.
	data["group/1337/id"] = []byte("1337")
	data[meetingFQID+"/group_ids"] = []byte("[1337]")
	data["group/1337/user_ids"] = []byte(fmt.Sprintf("[%d]", c.UserID))
	f := fmt.Sprintf("user/%d/group_$%d_ids", c.UserID, meetingID)
	data[f] = jsonAddInt(data[f], 1337)
	data["group/1337/meeting_id"] = []byte(strconv.Itoa(meetingID))
	if c.Permission != "" {
		data["group/1337/permissions"] = []byte(fmt.Sprintf(`["%s"]`, c.Permission))
	}

	p := permission.New(&DataProvider{Data: data})

	if strings.Contains(c.Handler, ".") {
		c.testWrite(t, p)
		return
	}
	c.testRead(t, p)
}

func convertPayload(in map[string]string) map[string]json.RawMessage {
	o := make(map[string]json.RawMessage, len(in))
	for k, v := range in {
		o[k] = []byte(v)
	}
	return o
}

func (c *Case) testWrite(t *testing.T, p *permission.Permission) {
	payload := []map[string]json.RawMessage{convertPayload(c.Payload)}

	_, err := p.IsAllowed(context.Background(), c.Handler, c.UserID, payload)

	if err != nil {
		var notAllowed perm.NotAllowedError
		if !errors.As(err, &notAllowed) {
			t.Fatalf("Got unexpected error: %v", err)
		}

		if c.ExpectedAllowed {
			t.Errorf("IsAllowed(%s) for case %s returned: %v", c.Handler, c.name, notAllowed)
		}
		return
	}

	if !c.ExpectedAllowed {
		t.Errorf("IsAllowed(%s) for case %s returned an unexpected allowed", c.Handler, c.name)
	}
}

func (c *Case) testRead(t *testing.T, p *permission.Permission) {
	got, err := p.RestrictFQFields(context.Background(), c.UserID, c.FQFields)
	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}

	if len(got) != len(c.ExpectedFQFields) {
		t.Errorf("RestrictFQField(%s) for case %s returned %v, expected %v", c.Handler, c.name, setToList(got), c.ExpectedAllowed)
		return
	}

	for _, f := range c.ExpectedFQFields {
		if !got[f] {
			t.Errorf("RestrictFQField(%s) for case %s did not allow %s", c.Handler, c.name, f)
		}
	}
}

// jsonAddInt adds the given int to the encoded json list.
//
// If the value exists in the list, the list is returned unchanged.
func jsonAddInt(list json.RawMessage, value int) json.RawMessage {
	var decoded []int
	if list != nil {
		json.Unmarshal(list, &decoded)
	}

	for _, i := range decoded {
		if i == value {
			return list
		}
	}

	decoded = append(decoded, value)
	list, _ = json.Marshal(decoded)
	return list
}

// defaultInt returns returns the given value or d, if value == 0
func defaultInt(value int, d int) int {
	if value == 0 {
		return d
	}
	return value
}

func setToList(s map[string]bool) []string {
	l := make([]string, 0, len(s))
	for k, v := range s {
		if !v {
			continue
		}
		l = append(l, k)
	}
	return l
}

// Cases holds a list of cases and knows how to be decoded from yaml.
type Cases struct {
	cases []Case
}

// UnmarshalYAML is the hock for the yml lib.
func (c *Cases) UnmarshalYAML(value *yaml.Node) error {
	var m map[string]Case
	if err := value.Decode(&m); err != nil {
		return fmt.Errorf("decoding cases: %w", err)
	}

	for k, v := range m {
		v.name = k
		c.cases = append(c.cases, v)
	}
	return nil
}
