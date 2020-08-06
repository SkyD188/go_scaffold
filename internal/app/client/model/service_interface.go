package model

type Service interface {
	Ping() string
}

type ServiceStruct struct {
}

func (s ServiceStruct) Ping() string {
	return "pong"
}
