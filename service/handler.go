package service

type Handler interface {
	Initialize(svc Service)
}
