package pkg

import (
	"github.com/google/wire"
	"test/tests/pkg/context"
	"test/tests/pkg/testcontainer"
)

var ProviderSet = wire.NewSet(
	context.ProviderSet,
	testcontainer.ProviderSet,
)

var APIMockProviderSet = wire.NewSet(
	context.APIMockProviderSet,
)
