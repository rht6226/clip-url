package idgen

import "math/rand"

// get random ID
type IdGeneratorService interface {
	GetId() uint64
}

type randomIdGeneratorService struct {
	start uint64
	end   uint64
}

// get new random number generator service
func NewRandomIdGeneratorService(start, end uint64) *randomIdGeneratorService {
	return &randomIdGeneratorService{
		start: start,
		end:   end,
	}
}

// return a new Id in a range
func (r *randomIdGeneratorService) GetId() uint64 {
	random := rand.Uint64()
	diff := r.end - r.start
	return r.start + (random % diff)
}
