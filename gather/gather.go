package gather

import "github.com/bxgdchen/device-communication/db"

// Gather ...
type Gather interface {
	GatherData(db db.DataAccessor) (map[string]string, error)
}
