package book

import (
	"github.com/google/wire"
	"rest-api/internal/app/book/handler"
	"rest-api/internal/app/book/usecase/usecaseimpl"
)

var ModuleSet = wire.NewSet(
	usecaseSet,
	deliverySet,
)
var (
	usecaseSet = wire.NewSet(
		wire.Struct(new(usecaseimpl.Book), "*"),
		usecaseimpl.NewBook,
		//wire.Bind(new(usecase.Book), new(*usecaseimpl.Book)),
	)

	deliverySet = wire.NewSet(
		wire.Struct(new(handler.BookRegistryOpts), "*"),
		handler.NewBookRegistry,
	)
)
