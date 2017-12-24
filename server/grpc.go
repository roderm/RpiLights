package server

var instance Service

type Service struct{}

func GetService() *Service {
	return &instance
}
