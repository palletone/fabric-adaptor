/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package preferpeer

import (
	"github.com/palletone/fabric-adaptor/pkg/common/providers/context"
	"github.com/palletone/fabric-adaptor/pkg/fab/events/client/lbp"
	"github.com/palletone/fabric-adaptor/pkg/fab/events/client/peerresolver"
)

type params struct {
	loadBalancePolicy lbp.LoadBalancePolicy
}

func defaultParams(context context.Client, channelID string) *params {
	return &params{
		loadBalancePolicy: peerresolver.GetBalancer(context.EndpointConfig().ChannelConfig(channelID).Policies.EventService),
	}
}

func (p *params) SetLoadBalancePolicy(value lbp.LoadBalancePolicy) {
	logger.Debugf("LoadBalancePolicy: %#v", value)
	p.loadBalancePolicy = value
}