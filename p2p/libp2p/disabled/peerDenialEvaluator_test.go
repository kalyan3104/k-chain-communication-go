package disabled_test

import (
	"testing"
	"time"

	"github.com/kalyan3104/k-chain-communication-go/p2p/libp2p/disabled"
	"github.com/kalyan3104/k-chain-core-go/core/check"
	"github.com/stretchr/testify/assert"
)

func TestPeerDenialEvaluator_ShouldWork(t *testing.T) {
	t.Parallel()

	pde := &disabled.PeerDenialEvaluator{}

	assert.False(t, check.IfNil(pde))
	assert.Nil(t, pde.UpsertPeerID("", time.Second))
	assert.False(t, pde.IsDenied(""))
}
