package initializers

import (
	"github.com/slimloans/golly"
	"github.com/slimloans/golly-skeleton/app/controllers"
)

// Preboots lists the preboots
var Preboots = []golly.PrebootFunc{
	configPreboot,
	// redis.Preboot,
}

// Initializers default app initializers - not sure if i like this yet
// id like eto keep the seperated for cleanliness
var Initializers = []golly.GollyAppFunc{
	// orm.Initializer, // See golly readme on how to configure
	// Always Last
	controllers.Initializer,
}
