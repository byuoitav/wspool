package wspool

import (
	"github.com/gorilla/websocket"
)

//closeWebsocket .
func closeWebsocket(ws *websocket.Conn, log Logger) {
	err := ws.WriteMessage(websocket.CloseMessage, []byte{})
	if err != nil {
		if log != nil {
			log.Warnf("failed to write close message: %s", err.Error())
		}
	}

	err = ws.Close()
	if err != nil {
		if log != nil {
			log.Warnf("failed to close websocket: %s", err.Error())
		}
	}

}
