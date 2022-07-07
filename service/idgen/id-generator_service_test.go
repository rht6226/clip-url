package idgen_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rht6226/clip-url/service/idgen"
)

var (
	start uint64 = 1000000000
	end   uint64 = 3000000000
)

var (
	idService idgen.IdGeneratorService = idgen.NewRandomIdGeneratorService(start, end)
)

func TestRandomIdGeneratorService(t *testing.T) {
	t.Run("SUCESS", func(t *testing.T) {
		randomId := idService.GetId()
		assert.Equal(t, true, randomId > start)
		assert.Equal(t, true, randomId <= end)
	})
}
