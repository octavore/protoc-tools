package proto3

type PageService interface {
	GetPage(*Timestamp) (*Page, error)
}
