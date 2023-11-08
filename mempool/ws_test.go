package mempool

import (
	"github.com/gorilla/websocket"
	"os"
	"testing"
)

// TestListen is an example of listening to QuickNode mempool.
//
//	RPC=abc go test -v ./mempool
func TestListen(t *testing.T) {
	url := os.Getenv("RPC")
	t.Logf("url = %s", url)
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("dial error: %s", err)
	}
	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			t.Log("read:", err)
			return
		}
		t.Logf("recv: %s", message)
	}
}
