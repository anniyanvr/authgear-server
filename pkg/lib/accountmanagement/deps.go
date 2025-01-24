package accountmanagement

import (
	"github.com/google/wire"
)

var DependencySet = wire.NewSet(
	wire.Struct(new(Service), "*"),
	wire.Struct(new(RedisStore), "*"),
	wire.Bind(new(Store), new(*RedisStore)),
	wire.Struct(new(RateLimitMiddleware), "*"),
)
