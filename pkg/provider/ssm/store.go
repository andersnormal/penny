package ssm

import (
	"context"
	"sync"

	"github.com/andersnormal/penny/pkg/provider"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var _ provider.Provider = (*Provider)(nil)

// Provider holds the SSM provider
type Provider struct {
	provider.AbstractProvider

	ssm *ssm.SSM

	mux sync.Mutex
}

// New returns a new Store with the fetched parameters, and service id
func New() provider.Provider {
	p := newProvider()

	return p
}

// ListKeysWithContext is listing keys in a directory and uses a context
func (p *Provider) ListKeysWithContext(ctx context.Context, dir string, recursive bool) ([]*provider.KVPair, error) {
	return p.listWithContext(ctx, dir, recursive, nil, nil)
}

// listWithContext is the private implementation, which is much more extensive then interface
func (p *Provider) listWithContext(ctx context.Context, dir string, recursive bool, keys []*provider.KVPair, nextToken *string) ([]*provider.KVPair, error) {
	var err error

	// input to the SSM to get parameters by path
	input := &ssm.GetParametersByPathInput{
		Path:      aws.String(dir),
		Recursive: aws.Bool(recursive),
		// WithDecryption: t.Bool(p.cfg.WithDecryption),
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
		Value: []byte(aws.StringValue(param.Value)),
		Key:   aws.StringValue(param.Name),
	}
}

func newProvider() *Provider {
	p := new(Provider)

	return p
}
