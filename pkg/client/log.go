package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"git.containerum.net/ch/kube-client/pkg/cherry"
	"git.containerum.net/ch/kube-client/pkg/rest"
	"github.com/gorilla/websocket"
)

const (
	followParam    = "follow"   // bool
	previousParam  = "previous" // bool
	tailParam      = "tail"     // int
	containerParam = "container"
)

const (
	// Maximum message size allowed from peer.
	maxMessageSize = 1024
)

func (client *Client) GetPodLogs(namespace, pod, container string, previous, follow bool, tail int) (*io.PipeReader, error) {
	logUrl, err := client.podLogUrl(namespace, pod, container, previous, follow, tail)
	if err != nil {
		return nil, err
	}

	conn, err := client.newWebsocketConnection(logUrl)
	if err != nil {
		return nil, err
	}

	pipeRd, pipeWr := io.Pipe()
	go client.logStream(conn, pipeWr)

	return pipeRd, nil
}

func (client *Client) podLogUrl(ns, pod, container string, previous, follow bool, tail int) (*url.URL, error) {
	queryUrl, err := url.Parse(client.APIurl)
	if err != nil {
		return nil, err
	}
	queryUrl.Path = fmt.Sprintf("/namespaces/%s/pods/%s/log", ns, pod)
	queryUrl.Query().Set(followParam, strconv.FormatBool(follow))
	queryUrl.Query().Set(previousParam, strconv.FormatBool(previous))
	queryUrl.Query().Set(tailParam, strconv.Itoa(tail))
	queryUrl.Query().Set(containerParam, container)
	return queryUrl, nil
}

func (client *Client) newWebsocketConnection(url *url.URL) (*websocket.Conn, error) {
	conn, httpResp, err := client.WSDialer.Dial(url.String(), http.Header{
		rest.HeaderUserFingerprint: {client.RestAPI.GetFingerprint()},
		rest.HeaderUserToken:       {client.RestAPI.GetToken()},
	})
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode >= 400 {
		var cherr cherry.Err
		if err := json.NewDecoder(httpResp.Body).Decode(&cherr); err != nil {
			return nil, err
		}
		return nil, &cherr
	}

	return conn, nil
}

func (client *Client) logStream(conn *websocket.Conn, wr *io.PipeWriter) {
	defer conn.Close()
	conn.SetReadLimit(maxMessageSize)
	for {
		mtype, data, err := conn.ReadMessage()
		if err != nil {
			wr.CloseWithError(err)
			return
		}
		switch mtype {
		case websocket.TextMessage, websocket.BinaryMessage:
		default:
			continue
		}
		if _, err := wr.Write(data); err != nil {
			return
		}
	}
}
