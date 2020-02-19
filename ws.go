package wspool

import (
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

//closeWebsocket .
func closeWebsocket(ctx context.Context, ws *websocket.Conn) error {
	err := ws.WriteMessage(websocket.CloseMessage, []byte{})
	if err != nil {
		return fmt.Errorf("failed to close websocket: %s", err.Error())
	}

	err = ws.Close()
	if err != nil {
		return fmt.Errorf("failed to close websocket: %s", err.Error())
	}

	return nil
}
