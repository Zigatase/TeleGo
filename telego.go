package telego

import (
	"encoding/json"
	"github.com/Zigatase/telego/e"
	"github.com/Zigatase/telego/types"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	tgBotApiHost = "api.telegram.org"

	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

// Client занимается 2 вещами получает update и отправляет сообщения
type Client struct {
	host     string // Host Api Telegram
	basePath string // bot<TOKEN>/method
	client   http.Client
}

func New(token string) *Client {
	return &Client{
		host:     tgBotApiHost,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) ([]types.Update, error) {
	// https://core.telegram.org/bots/api#getupdates
	q := url.Values{}

	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	// do request
	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	var res types.UpdatesResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

func (c *Client) SendMessageText(chatId int, msg string) error {
	// https://core.telegram.org/bots/api#sendmessage
	q := url.Values{}

	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("text", msg)

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return e.Wrap("Can't send message", err)
	}

	return nil
}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, e.Wrap("Can't do request", err)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, e.Wrap("Can't do request", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
