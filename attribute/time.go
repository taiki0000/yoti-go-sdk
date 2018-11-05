package attribute

import (
	"log"
	"time"

	"github.com/getyoti/yoti-go-sdk/anchor"
	"github.com/getyoti/yoti-go-sdk/yotiprotoattr_v3"
)

//Time is a Yoti attribute which returns a time as its value
type Time struct {
	*yotiprotoattr_v3.Attribute
	Value   *time.Time
	Anchors []*anchor.Anchor
}

//NewTime creates a new Time attribute
func NewTime(a *yotiprotoattr_v3.Attribute) (*Time, error) {
	parsedTime, err := time.Parse("2006-01-02", string(a.Value))
	if err != nil {
		log.Printf("Unable to parse time value of: %q. Error: %q", a.Value, err)
		parsedTime = time.Time{}
		return nil, err
	}

	return &Time{
		Attribute: &yotiprotoattr_v3.Attribute{
			Name:        a.Name,
			ContentType: a.ContentType,
		},
		Value:   &parsedTime,
		Anchors: anchor.ParseAnchors(a.Anchors),
	}, nil
}
