package telego

type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	// https://core.telegram.org/bots/api#getting-updates

	ID      int    `json:"update_id"`
	Message string `json:"message"`
}
