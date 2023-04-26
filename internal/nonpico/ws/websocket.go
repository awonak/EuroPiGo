//go:build !pico
// +build !pico

package ws

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pongWait = time.Second * 60
)

type WebSocket struct {
	conn   *websocket.Conn
	mu     sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc
	err    error
}

func (w *WebSocket) Done() <-chan struct{} {
	if w.ctx == nil {
		return nil
	}

	return w.ctx.Done()
}

func (w *WebSocket) WriteJSON(msg any) error {
	if w.conn == nil {
		return errors.New("connection not established")
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	err := w.conn.WriteJSON(msg)
	if err != nil {
		w.SetError(err)
	}
	return err
}

func (w *WebSocket) ReadMessage() ([]byte, error) {
	_, blob, err := w.conn.ReadMessage()
	if err != nil && websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
		w.SetError(err)
	}
	return blob, err
}

func (w *WebSocket) SetError(err error) {
	w.err = err
	if w.cancel != nil {
		w.cancel()
	}
}

func (w *WebSocket) Close() error {
	if w.cancel != nil {
		w.cancel()
	}

	if w.conn == nil {
		return nil
	}

	return w.conn.Close()
}

var upgrader = websocket.Upgrader{}

func Upgrade(w http.ResponseWriter, r *http.Request) (*WebSocket, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	conn.SetReadLimit(2048)
	_ = conn.SetWriteDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(appData string) error {
		return conn.SetWriteDeadline(time.Now().Add(pongWait))
	})

	ws := &WebSocket{
		conn:   conn,
		ctx:    ctx,
		cancel: cancel,
	}

	return ws, nil
}
