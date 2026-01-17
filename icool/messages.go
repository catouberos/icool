package icool

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func (c *Client) Queue(name, imageURL, youtubeID string) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	req := fmt.Sprintf(`42["playlist_add",{"data":{"isPriority":false,"media":{"id":"%s","songId":null,"name":"%s","duration":0,"imageUrl":"%s","isRecording":false,"singers":[],"type":"KTVYoutubeMedia","youtubeId":"%s","genres":[],"isSoundCloud":true}}}]`, id, name, imageURL, youtubeID)

	err = c.conn.WriteMessage(websocket.TextMessage, []byte(req))
	if err != nil {
		return err
	}

	return nil
}
