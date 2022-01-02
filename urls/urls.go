package urls

import (
	"math/rand"
)

// User holds data for single user
type Url struct {
	ID       string `json: "id"`
	LongUrl  string `json:"longurl"`
	ShortUrl string `json:"shorturl"`
}

func AllUrls() ([]Url, error) {
	urls := []Url{}
	return urls, nil
}

func (u *Url) GenrtaeUrl() {
	u.ID = randStringRunes()
	u.ShortUrl = "http://127.0.0.1:8081/" + u.ID
}

func randStringRunes() string {
	NUM_CHARS_SHORT_LINK := 7
	ALPHABET := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, NUM_CHARS_SHORT_LINK)
	for i := range b {
		b[i] = ALPHABET[rand.Intn(len(ALPHABET))]
	}
	return string(b)
}
