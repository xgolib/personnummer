package personnummer

import (
	"strings"
	"time"
	"unicode"

	pn "github.com/personnummer/go/v3"
)

// Personnummer represents the personnummer struct.
type Personnummer struct {
	ID                   string    `json:"id"`
	Birthday             time.Time `json:"birthday"`
	Gender               string    `json:"gender"`
	Delimiter            string    `json:"delimiter"`
	IsCoordinationNumber bool      `json:"is_coordination_number"`
}

func Parse(id string) (*Personnummer, error) {
	id = strings.Map(func(r rune) rune {
		if !unicode.IsSpace(r) {
			return r
		}
		return -1
	}, id)
	ret, err := pn.New(id)
	if err != nil {
		return nil, err
	}
	formatID, _ := ret.Format(true)
	gender := "female"
	if ret.IsMale() {
		gender = "male"
	}
	var t time.Time
	if ret.IsCoordinationNumber() {
		day := []byte(formatID[0:8])
		day[6] = day[6] - 6
		t, _ = time.Parse("20060102", string(day))
	} else {
		t, _ = time.Parse("20060102", formatID[0:8])
	}
	val := &Personnummer{ID: formatID, Gender: gender, Birthday: t,
		Delimiter: ret.Sep, IsCoordinationNumber: ret.IsCoordinationNumber()}
	return val, nil
}
