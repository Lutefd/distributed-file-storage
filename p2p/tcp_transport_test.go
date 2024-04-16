package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T){
	tcpOpts := TCPTransportOpts{
		ListenAddr: ":3000",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder: GOBDecoder{},
	}
	tr := NewTCPTransport(tcpOpts)

	assert.Equal(t, tr.listenAddress, tcpOpts.ListenAddr)

	assert.Nil(t,tr.ListenAndAccept())

}