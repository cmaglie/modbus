// Copyright 2014 Quoc-Viet Nguyen. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license. See the LICENSE file for details.

package modbus

// RTUTCPClientHandler implements Packager and Transporter interface.
type RTUTCPClientHandler struct {
	rtuPackager
	tcpTransporter
}

// NewRTUTCPClientHandler allocates and initializes a RTUTCPClientHandler.
func NewRTUTCPClientHandler(address string) *RTUTCPClientHandler {
	handler := &RTUTCPClientHandler{}
	handler.Address = address
	handler.Timeout = tcpTimeout
	handler.IdleTimeout = tcpIdleTimeout
	return handler
}
