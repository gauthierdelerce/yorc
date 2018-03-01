// Copyright 2018 Bull S.A.S. Atos Technologies - Bull, Rue Jean Jaures, B.P.68, 78340, Les Clayes-sous-Bois, France.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

import (
	"context"
	"testing"

	plugin "github.com/hashicorp/go-plugin"
)

func Test_clientMonitorContextCancellation(t *testing.T) {
	type args struct {
		ctx       context.Context
		closeChan chan struct{}
		id        uint32
		broker    *plugin.MuxBroker
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientMonitorContextCancellation(tt.args.ctx, tt.args.closeChan, tt.args.id, tt.args.broker)
		})
	}
}

func TestRPCContextCanceller_CancelContext(t *testing.T) {
	type args struct {
		nothing interface{}
		resp    *CancelContextResponse
	}
	tests := []struct {
		name    string
		r       *RPCContextCanceller
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.CancelContext(tt.args.nothing, tt.args.resp); (err != nil) != tt.wantErr {
				t.Errorf("RPCContextCanceller.CancelContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
