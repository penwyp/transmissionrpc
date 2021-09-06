package transmissionrpc

import (
	"context"
	"fmt"
)

/*
	Port Checking
	https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L584
*/

// PortTest allows tests to see if your incoming peer port is accessible from the outside world.
func (c *Client) PortTest(ctx context.Context) (open bool, err error) {
	var result portTestAnswer
	// Send request
	if err = c.rpcCall(ctx, "port-test", nil, &result); err == nil {
		open = result.PortOpen
	} else {
		err = fmt.Errorf("'port-test' rpc method failed: %w", err)
	}
	return
}

type portTestAnswer struct {
	PortOpen bool `json:"port-is-open"`
}
