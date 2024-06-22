package app

import (
	"github.com/google/wire"
	"rest-api/internal/app/book"
)

func InjectApp() (func(), error) {
	panic(wire.Build(
		book.ModuleSet,
	))
}
