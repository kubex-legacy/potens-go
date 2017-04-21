package services

type Service struct {
	key string
}

func (s Service) Key() string {
	return s.key
}
