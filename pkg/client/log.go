package client

import (
	"bufio"
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

type LogReader struct {
	stop chan struct{}
	*bufio.Scanner
}

func (logre *LogReader) Close() error {
	close(logre.stop)
	return logre.Err()
}
func NewLogReader(re io.Reader) (*LogReader, <-chan struct{}) {
	stop := make(chan struct{})
	return &LogReader{
		stop,
		bufio.NewScanner(re),
	}, stop
}

type GetPodLogsParams struct {
	Namespace, Pod, Container string
	Previous, Follow          bool
	Tail                      int
}

func (client *Client) GetPodLogs(params GetPodLogsParams) (*LogReader, error) {
	logUrl, err := client.podLogUrl(params)
	if err != nil {
		return nil, err
	}
	conn, err := client.newWebsocketConnection(logUrl)
	if err != nil {
		return nil, err
	}
	re, wr := io.Pipe()
	logReader, stop := NewLogReader(re)
	go client.logStream(stop, conn, wr)
	return logReader, nil
}

func (client *Client) podLogUrl(params GetPodLogsParams) (*url.URL, error) {
	queryUrl, err := url.Parse(client.APIurl)
	if err != nil {
		return nil, err
	}
	switch queryUrl.Scheme {
	case "http":
		queryUrl.Scheme = "ws"
	case "https":
		queryUrl.Scheme = "wss"
	}
	queryUrl.Path = fmt.Sprintf("/namespaces/%s/pods/%s/log", params.Namespace, params.Pod)
	queryUrl.Query().Set(followParam, strconv.FormatBool(params.Follow))
	queryUrl.Query().Set(previousParam, strconv.FormatBool(params.Previous))
	queryUrl.Query().Set(tailParam, strconv.Itoa(params.Tail))
	queryUrl.Query().Set(containerParam, params.Container)
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

func (client *Client) logStream(stop <-chan struct{}, conn *websocket.Conn, out *io.PipeWriter) {
	defer conn.Close()
	conn.SetReadLimit(maxMessageSize)
	for {
		select {
		case <-stop:
			out.Close()
			return
		default:
			mtype, data, err := conn.ReadMessage()
			if err != nil {
				out.CloseWithError(err)
			}
			switch mtype {
			case websocket.TextMessage, websocket.BinaryMessage:
				_, err := out.Write(data)
				if err != nil {
					out.CloseWithError(err)
				}
			default:
				continue
			}
		}
	}
}
