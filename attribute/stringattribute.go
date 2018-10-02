package attribute

import (
	"github.com/getyoti/yoti-go-sdk/anchor"
)

//StringAttribute is a Yoti attribute which returns a string as its value
type StringAttribute struct {
	Name    string
	Value   string
	Type    AttrType
	Anchors []*anchor.Anchor
}

//NewStringAttribute creates a new String attribute
func NewStringAttribute(a *Attribute) *StringAttribute {
	return &StringAttribute{
		Name:    a.Name,
		Value:   string(a.Value),
		Type:    AttrTypeString,
		Anchors: a.Anchors,
	}
}