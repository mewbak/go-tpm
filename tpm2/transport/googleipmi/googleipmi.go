// Package googleipmi provides functions for connecting to and initializing a
// Titan TPM IPMI device.
package googleipmi

import (
	"context"
	"fmt"

	"github.com/google/go-tpm/tpm2/transport"
)

const (
	ecTitanSendSendTPMCommand uint16 = 0x3e33
)

// titanTPM implements go-tpm's transport.TPM interface, for transmitting TPM
// commands to a Titan TPM.
type titanTPM struct {
	cd commandDispatcher
}

// Send implements the transport.TPM interface.
func (t *titanTPM) Send(input []byte) ([]byte, error) {
	ctx := context.Background()
	rsp, err := ecSendCommand(ctx, t.cd, ecCommand{ecTitanSendSendTPMCommand, 0}, input)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// Close closes the connection to the TPM.
func (t *titanTPM) Close() error {
	return t.cd.close()
}

// Open creates a TPM connection to a Titan device via IPMI.
func Open() (transport.TPMCloser, error) {
	opts := options{
		DevicePath: "/dev/ipmi0",
	}

	titan, err := new(opts)
	if err != nil {
		return nil, fmt.Errorf("[IPMI TPM] err: %v", err)
	}
	return &titanTPM{
		cd: titan,
	}, nil
}
