package service_test

import (
	"testing"

	service "github.com/rht6226/clip-url/service/encodec"
	"github.com/stretchr/testify/assert"
)

var (
	base62EncoDecService service.EncoDecService = service.NewBase62EncoDecService()
)

func TestBase62EncodecService(t *testing.T) {
	var id uint64 = 10056789342

	t.Run("SUCCESS", func(t *testing.T) {
		encodedUrl := base62EncoDecService.Encode(id)
		decodedId, err := base62EncoDecService.Decode(encodedUrl)

		if err != nil {
			t.Fatalf("Failed to decode")
		}

		assert.Equal(t, decodedId, id)
	})

	t.Run("INVALID_URL", func(t *testing.T) {
		url := "Cax-Dcv45!"

		_, err := base62EncoDecService.Decode(url)

		assert.Error(t, err)
	})
}
