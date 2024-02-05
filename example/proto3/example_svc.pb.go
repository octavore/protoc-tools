package proto3

import (
	context "context"
)

type PageService interface {
	GetPage(context.Context, *Timestamp) (*Page, error)
}
