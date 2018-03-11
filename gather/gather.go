package gather

import "git.sy2020.cn/sykjcgq/device-communication/db"

// Gather ...
type Gather interface {
	GatherData(db db.DataAccessor) (map[string]string, error)
}
