package wsmock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"git.containerum.net/ch/kube-client/pkg/websocket/wsmock/errwsmock"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
}

func (s *EchoServer) echoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		retErr := errwsmock.ErrUpgradeFailed().AddDetailsErr(err)
		w.WriteHeader(retErr.StatusHTTP)
		json.NewEncoder(w).Encode(retErr)
		return
	}

	for {
		msgType, data, err := conn.ReadMessage()
		if err != nil {
			conn.WriteMessage(msgType, []byte(fmt.Sprintf("read error: %v", err)))
			return
		}

		if err := conn.WriteMessage(msgType, data); err != nil {
			return
		}
	}
}

func (p *PeriodicServer) periodicHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		retErr := errwsmock.ErrUpgradeFailed().AddDetailsErr(err)
		w.WriteHeader(retErr.StatusHTTP)
		json.NewEncoder(w).Encode(retErr)
		return
	}

	ticker := time.NewTicker(p.cfg.MsgPeriod)

	for {
		<-ticker.C
		if err := conn.WriteMessage(websocket.TextMessage, []byte(p.cfg.MsgText)); err != nil {
			ticker.Stop()
			return
		}
	}
}
