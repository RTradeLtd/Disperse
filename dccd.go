package disperse

import (
	"errors"
	"fmt"
	"time"

	ipfsapi "github.com/RTradeLtd/go-ipfs-api"
)

type DCCDManager struct {
	Shell    *ipfsapi.Shell
	Gateways map[int]string
	TimeOut  time.Duration
}

// NewDCCDManager establishes our initial connection to our local IPFS node
func NewDCCDManager(connectionURL string, timeout time.Duration) *DCCDManager {
	if connectionURL == "" {
		// load a default api
		connectionURL = "localhost:5001"
	}
	c := generateClient(timeout)
	shell := ipfsapi.NewShellWithClient(connectionURL, c)
	manager := &DCCDManager{Shell: shell}
	manager.ParseGateways()
	return manager
}

func (dc *DCCDManager) ParseGateways() {
	indexes := make(map[int]string)
	for k, v := range GateArrays {
		indexes[k] = v
	}
	dc.Gateways = indexes
}

func (dc *DCCDManager) ReconnectShell(connectionURL string) error {
	if connectionURL == "" {
		return errors.New("please provide a valid connection url")
	}
	c := generateClient(dc.TimeOut)
	shell := ipfsapi.NewShellWithClient(connectionURL, c)
	dc.Shell = shell
	return nil
}

func (dc *DCCDManager) DisperseContentWithShell(contentHash string) (map[string]bool, error) {
	m := make(map[string]bool)
	for _, v := range GateArrays {
		err := dc.ReconnectShell(v)

		r, err := dc.Shell.CatGet(contentHash)
		if err != nil {
			m[v] = false
			fmt.Println("dispersal failed for host ", v)
			continue
		}
		err = r.Close()
		if err != nil {
			fmt.Println("failed to close handler ", err)
		}
		fmt.Println("dispersal suceeded for host ", v)
		m[v] = true
	}
	return m, nil
}
