package graph

import (
	"github.com/ledgerwatch/erigon-lib/kv"
	"github.com/optimism-java/erigon/turbo/jsonrpc"
	"github.com/optimism-java/erigon/turbo/rpchelper"
	"github.com/optimism-java/erigon/turbo/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	GraphQLAPI  jsonrpc.GraphQLAPI
	db          kv.RoDB
	filters     *rpchelper.Filters
	blockReader services.FullBlockReader
}
