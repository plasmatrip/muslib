package model

import (
	"strings"
	"time"
)

type Song struct {
	Group string `json:"group"`
	Song  string `json:"song"`
	SongDetail
}

type SongDetail struct {
	ReleaseDate ReleaseDate `json:"releaseDate"`
	Text        string      `json:"text,omitempty"`
	Link        string      `json:"link,omitempty"`
}

type Filter struct {
	Group       *string
	Song        *string
	Text        *string
	Link        *string
	ReleaseFrom *time.Time
	ReleaseTo   *time.Time
	Limit       int
	Offset      int
}

type VerseResponse struct {
	Song        string `json:"song"`
	Group       string `json:"group"`
	Verse       string `json:"verse"`
	VerseNum    int    `json:"verse_num"`
	TotalVerses int    `json:"total_verses"`
}

type ReleaseDate time.Time

func (c *ReleaseDate) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("02-01-2006", value)
	if err != nil {
		return err
	}
	*c = ReleaseDate(t)
	return nil
}

func (c ReleaseDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format("02-01-2006") + `"`), nil
}

func (c ReleaseDate) NilIfZero() interface{} {
	if time.Time(c).IsZero() {
		return nil
	}
	return time.Time(c)
}
