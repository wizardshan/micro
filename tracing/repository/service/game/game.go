package game

import "tracing/pkg/httpquery"

type SignField struct {
	Source  int    `url:"source"`
	SignKey string `url:"sign_key"`
}

func (s *SignField) signKey(v string) {
	s.SignKey = v
}

func (s *SignField) source(v int) {
	s.Source = v
}

type Encryptor interface {
	httpquery.Builder
	Encode() string
	source(source int)
	signKey(signKey string)
}
