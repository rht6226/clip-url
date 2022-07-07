package service

import (
	"math"
	"strings"

	apperrors "github.com/rht6226/clip-url/errors"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = uint64(len(alphabet))
)

type base62EncoDecService struct{}

// returns a new Encodec service based on Base62 service
func NewBase62EncoDecService() *base62EncoDecService {
	return &base62EncoDecService{}
}

// returns a base 62 encoded string for a given id
func (svc *base62EncoDecService) Encode(id uint64) string {
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(11)

	for ; id > 0; id = id / length {
		encodedBuilder.WriteByte(alphabet[(id % length)])
	}

	return encodedBuilder.String()
}

// returns the decoded id for a given string
// returns error if decoding not possible
func (svc *base62EncoDecService) Decode(shortLink string) (uint64, error) {
	var number uint64

	for i, symbol := range shortLink {
		alphabeticPosition := strings.IndexRune(alphabet, symbol)

		if alphabeticPosition == -1 {
			return uint64(alphabeticPosition), apperrors.NewBadRequest("Invalid short url")
		}
		number += uint64(alphabeticPosition) * uint64(math.Pow(float64(length), float64(i)))
	}

	return number, nil

}
