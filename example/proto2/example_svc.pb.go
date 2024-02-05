package proto2

import (
	context "context"
)

type PageService interface {
	GetPage(context.Context, *Timestamp) (*Page, error)
}
