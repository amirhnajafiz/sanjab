package storage

import (
	"fmt"

	"github.com/ceph/go-ceph/rados"
)

type Storage struct {
	conn   *rados.Conn
	ioConn *rados.IOContext
}

// NewConnection returns a rados connection
func NewConnection(confFile, pool string) (*Storage, error) {
	// create a connection to the Ceph cluster
	conn, err := rados.NewConn()
	if err != nil {
		return nil, fmt.Errorf("failed to create connection: %v", err)
	}

	// read the Ceph configuration from a file or provide it programmatically
	err = conn.ReadConfigFile(confFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read Ceph config file: %v", err)
	}

	// initialize the connection
	err = conn.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ceph cluster: %v", err)
	}

	// open pool
	ioConn, err := conn.OpenIOContext(pool)
	if err != nil {
		return nil, fmt.Errorf("failed to open pool connection: %v", err)
	}

	return &Storage{
		conn:   conn,
		ioConn: ioConn,
	}, nil
}
