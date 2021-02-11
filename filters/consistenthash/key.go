package consistenthash

import (
	"github.com/zalando/skipper/eskip"
	"github.com/zalando/skipper/filters"
	"github.com/zalando/skipper/loadbalancer"
)

type consistentHashKey struct {
	template *eskip.Template
}

func NewConsistentHashKey() filters.Spec { return &consistentHashKey{} }
func (*consistentHashKey) Name() string  { return "consistentHashKey" }

func (*consistentHashKey) CreateFilter(args []interface{}) (filters.Filter, error) {
	if len(args) != 1 {
		return nil, filters.ErrInvalidFilterParameters
	}
	value, ok := args[0].(string)
	if !ok {
		return nil, filters.ErrInvalidFilterParameters
	}
	return &consistentHashKey{eskip.NewTemplate(value)}, nil
}

func (c *consistentHashKey) Request(ctx filters.FilterContext) {
	if key, ok := c.template.ApplyContext(ctx); ok {
		ctx.StateBag()[loadbalancer.ConsistentHashKey] = key
	}
}

func (*consistentHashKey) Response(filters.FilterContext) {}
