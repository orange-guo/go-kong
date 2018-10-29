package kong

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Configuration represents a config of a plugin in Kong.
type Configuration map[string]interface{}

// DeepCopyInto copies the receiver, writing into out. in must be non-nil.
func (in Configuration) DeepCopyInto(out *Configuration) {
	// Resorting to JSON since interface{} cannot be DeepCopied easily.
	// This could be replaced using reflection-fu.
	// XXX Ignoring errors
	b, _ := json.Marshal(&in)
	_ = json.Unmarshal(b, out)
}

// DeepCopy copies the receiver, creating a new Configuration.
func (in Configuration) DeepCopy() Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return *out
}

// Plugin represents a Plugin in Kong.
// Read https://getkong.org/docs/0.13.x/admin-api/#Plugin-object
// +k8s:deepcopy-gen=true
type Plugin struct {
	CreatedAt  *int          `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	ID         *string       `json:"id,omitempty" yaml:"id,omitempty"`
	Name       *string       `json:"name,omitempty" yaml:"name,omitempty"`
	RouteID    *string       `json:"route_id,omitempty" yaml:"route_id,omitempty"`
	ServiceID  *string       `json:"service_id,omitempty" yaml:"service_id,omitempty"`
	APIID      *string       `json:"api_id,omitempty" yaml:"api_id,omitempty"`
	ConsumerID *string       `json:"consumer_id,omitempty" yaml:"consumer_id,omitempty"`
	Config     Configuration `json:"config,omitempty" yaml:"config,omitempty"`
	Enabled    *bool         `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

// Valid checks if all the fields in Plugin are valid.
func (p *Plugin) Valid() bool {
	return !isEmptyString(p.Name)
}

func (p *Plugin) String() string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	buf.WriteByte(' ')
	if isEmptyString(p.ID) {
		buf.WriteString("nil")
	} else {
		buf.WriteString(*p.ID)
	}
	buf.WriteByte(' ')
	if isEmptyString(p.Name) {
		buf.WriteString("nil")
	} else {
		buf.WriteString(*p.Name)
	}
	buf.WriteByte(' ')
	if isEmptyString(p.RouteID) {
		buf.WriteString("nil")
	} else {
		buf.WriteString(*p.RouteID)
	}
	buf.WriteByte(' ')
	if isEmptyString(p.ServiceID) {
		buf.WriteString("nil")
	} else {
		buf.WriteString(*p.ServiceID)
	}
	buf.WriteByte(' ')
	if isEmptyString(p.APIID) {
		buf.WriteString("nil")
	} else {
		buf.WriteString(*p.APIID)
	}
	buf.WriteByte(' ')
	if isEmptyString(p.ConsumerID) {
		buf.WriteString("nil")
	} else {
		buf.WriteString(*p.ConsumerID)
	}
	buf.WriteByte(' ')
	buf.WriteString(fmt.Sprint(p.Config))
	buf.WriteByte(' ')
	buf.WriteByte(']')
	return buf.String()
}
