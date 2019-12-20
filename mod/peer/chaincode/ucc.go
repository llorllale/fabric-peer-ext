/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chaincode

import (
	"github.com/hyperledger/fabric/extensions/chaincode/api"
	"github.com/trustbloc/fabric-peer-ext/pkg/chaincode/ucc"
)

// GetUCC returns the in-process user chaincode for the given ID
func GetUCC(ccID string) (api.UserCC, bool) {
	return ucc.Get(ccID)
}

// Chaincodes returns all registered in-process chaincodes
func Chaincodes() []api.UserCC {
	return ucc.Chaincodes()
}

// WaitForReady blocks until the chaincodes are all registered
func WaitForReady() {
	ucc.WaitForReady()
}
