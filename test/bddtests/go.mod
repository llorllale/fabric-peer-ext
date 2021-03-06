// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/trustbloc/fabric-peer-ext/test/bddtests

require (
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
	github.com/cucumber/godog v0.8.1
	github.com/hyperledger/fabric-protos-go v0.0.0
	github.com/hyperledger/fabric-sdk-go v1.0.0-beta1.0.20200222173625-ff3bdd738791
	github.com/pkg/errors v0.8.1
	github.com/spf13/viper v1.0.2
	github.com/trustbloc/fabric-peer-test-common v0.1.2
)

replace github.com/hyperledger/fabric-protos-go => github.com/trustbloc/fabric-protos-go-ext v0.1.2

go 1.13
