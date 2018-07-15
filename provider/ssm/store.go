// Copyright 2018 Sebastian Döll
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ssm

import (
	"context"

	config "github.com/andersnormal/penny/cfg"
	"github.com/andersnormal/penny/provider"
	t "github.com/andersnormal/penny/tools"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var _ provider.Provider = (*Provider)(nil)

// New returns a new Store with the fetched parameters, and service id
func New() *Provider {
	return &Provider{}
}

// Must returns a new Store with parameters configured
func Must(ssm *ssm.SSM, ssmCfg *config.SSMConfig) *Provider {
	var p = New()

	p.ssm = ssm
	p.cfg = ssmCfg

	return p
}

// ListWithContext is listing keys in a directory and uses a context
func (p *Provider) ListWithContext(ctx context.Context, dir string, recursive bool) ([]*provider.KVPair, error) {
	return p.listWithContext(ctx, dir, recursive, nil, nil)
}

// listWithContext is the private implementation, which is much more extensive then interface
func (p *Provider) listWithContext(ctx context.Context, dir string, recursive bool, keys []*provider.KVPair, nextToken *string) ([]*provider.KVPair, error) {
	var err error

	// input to the SSM to get parameters by path
	input := &ssm.GetParametersByPathInput{
		Path:           t.String(dir),
		Recursive:      t.Bool(recursive),
		WithDecryption: t.Bool(p.cfg.WithDecryption),
	}

	if nextToken != nil {
		input.NextToken = nextToken
	}

	output, err := p.ssm.GetParametersByPathWithContext(ctx, input)
	if err != nil {
		return keys, err
	}

	for _, param := range output.Parameters {
		keys = append(keys, parameterKVPair(param))
	}

	// s.parameters = append(s.parameters, output.Parameters...)

	if nextToken != nil {
		p.listWithContext(ctx, dir, recursive, keys, nextToken)
	}

	return keys, err
}

func parameterKVPair(param *ssm.Parameter) *provider.KVPair {
	return &provider.KVPair{
		Value: []byte(t.StringValue(param.Value)),
		Key:   t.StringValue(param.Name),
	}
}
