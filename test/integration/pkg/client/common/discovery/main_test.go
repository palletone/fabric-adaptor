/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package discovery

import (
	"testing"

	"github.com/palletone/fabric-adaptor/pkg/fabsdk"
	"github.com/palletone/fabric-adaptor/test/integration"
	"github.com/palletone/fabric-adaptor/test/integration/util/runner"
)

const (
	org1Name     = "Org1"
	org2Name     = "Org2"
	adminUser    = "Admin"
	org1User     = "User1"
	orgChannelID = "orgchannel"
)

var mainSDK *fabsdk.FabricSDK
var mainTestSetup *integration.BaseSetupImpl

func TestMain(m *testing.M) {
	r := runner.New()
	r.Initialize()
	mainSDK = r.SDK()
	mainTestSetup = r.TestSetup()

	r.Run(m)
}
