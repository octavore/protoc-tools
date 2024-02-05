package proto2

type PageService interface {
	GetPage(*Timestamp) (*Page, error)
}
