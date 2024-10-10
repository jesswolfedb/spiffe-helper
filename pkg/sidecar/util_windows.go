//go:build windows
// +build windows

package sidecar

import (
	"errors"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

func (s *Sidecar) getWorkloadAPIAddress() workloadapi.ClientOption {
	return workloadapi.WithNamedPipeName(s.config.AgentAddress)
}

func (s *Sidecar) SignalProcess() error {
	// Signal to reload certs
	// TODO: it is not possible to get signal by name on windows,
	// we must provide int here
	return errors.New("sending signal is not supported on windows")
}
