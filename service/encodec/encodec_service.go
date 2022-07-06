package service

// encoder-decoder service takes care of encoding and decoding of the urls
type EncoDecService interface {
	Encode(uint64) string
	Decode(string) (uint64, error)
}