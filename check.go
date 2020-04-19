package checkup

import (
	"encoding/json"
	"fmt"

	"github.com/ykorzikowski/checkup/check/dns"
	"github.com/ykorzikowski/checkup/check/exec"
	"github.com/ykorzikowski/checkup/check/http"
	"github.com/ykorzikowski/checkup/check/tcp"
	"github.com/ykorzikowski/checkup/check/tls"
)

func checkerDecode(typeName string, config json.RawMessage) (Checker, error) {
	switch typeName {
	case dns.Type:
		return dns.New(config)
	case exec.Type:
		return exec.New(config)
	case http.Type:
		return http.New(config)
	case tcp.Type:
		return tcp.New(config)
	case tls.Type:
		return tls.New(config)
	default:
		return nil, fmt.Errorf(errUnknownCheckerType, typeName)
	}
}
