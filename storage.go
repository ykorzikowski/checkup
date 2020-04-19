package checkup

import (
	"encoding/json"
	"fmt"

	"github.com/ykorzikowski/checkup/storage/fs"
	"github.com/ykorzikowski/checkup/storage/github"
	"github.com/ykorzikowski/checkup/storage/s3"
	"github.com/ykorzikowski/checkup/storage/sql"
)

func storageDecode(typeName string, config json.RawMessage) (Storage, error) {
	switch typeName {
	case s3.Type:
		return s3.New(config)
	case github.Type:
		return github.New(config)
	case fs.Type:
		return fs.New(config)
	case sql.Type:
		return sql.New(config)
	default:
		return nil, fmt.Errorf(errUnknownStorageType, typeName)
	}
}
