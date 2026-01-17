package icool

import (
	"context"
	"net/url"

	"github.com/gorilla/websocket"
)

type Client struct {
	dialer *websocket.Dialer
	conn   *websocket.Conn
}

func Dial(ctx context.Context, host, room string) (*Client, error) {
	url := buildClientURL(host, room)

	client := &Client{
		dialer: websocket.DefaultDialer,
	}

	var err error
	client.conn, _, err = client.dialer.DialContext(ctx, url.String(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

// listen and do nothing
func (c *Client) Listen(ctx context.Context, msgChan chan string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			_, _, err := c.conn.ReadMessage()
			if err != nil {
				return err
			}
		}
	}
}

func buildClientURL(host, room string) url.URL {
	u := url.URL{Scheme: "ws", Host: host, Path: "/socket.io/"}
	q := u.Query()
	q.Add("token", room)
	q.Add("engine", "web")
	q.Add("EIO", "4")
	q.Add("transport", "websocket")
	u.RawQuery = q.Encode()

	return u
}
