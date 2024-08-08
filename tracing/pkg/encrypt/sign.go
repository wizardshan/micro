package encrypt

type SignField struct {
	Source  int    `url:"source"`
	SignKey string `url:"sign_key"`
}

func (s *SignField) Build(source int, signKey string) {
	s.Source = source
	s.SignKey = signKey
}

type Signer interface {
	Encode() string
	Build(source int, signKey string)
}

func Sign(value any) string {
	v, ok := value.(Signer)
	if !ok {
		panic("value must be encoder")
	}

	return MD5(v.Encode())
}

//func Sign(value any, salt string) string {
//	query := ""
//	v, ok := value.(Encoder)
//	if ok {
//		query = v.Encode()
//	}
//
//	return MD5(fmt.Sprintf("%s|%s", query, salt))
//}
