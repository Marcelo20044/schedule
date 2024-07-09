package dto

import (
	"strings"
	"time"
)

type CustomDate time.Time
type CustomTime time.Time

func (d *CustomDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse("02.01.2006", s)
	if err != nil {
		return err
	}
	*d = CustomDate(t)
	return nil
}

func (t *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	parsedTime, err := time.Parse("15:04", s)
	if err != nil {
		return err
	}
	*t = CustomTime(parsedTime)
	return nil
}
