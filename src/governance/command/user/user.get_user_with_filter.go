package user

import "github.com/niko-labs/libs-go/helper/paginator"

type CommandGetUserWithFilter struct {
	Pagination paginator.Pagination
}

func (c CommandGetUserWithFilter) IsCommand() {}

func (c CommandGetUserWithFilter) Data() any {
	return &c
}
