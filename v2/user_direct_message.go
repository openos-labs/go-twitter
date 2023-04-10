package twitter

type UserDirectMessageData struct {
	ID               string `json:"id"`
	Text             string `json:"text"`
	EventType        string `json:"event_type"`
	DmConversationId string `json:"dm_conversation_id"`
	CreatedAt        string `json:"created_at,omitempty"`
	SenderId         string `json:"sender_id"`
}

type UserDirectMessageCreateOpt struct {
	Text string `json:"text"`
}

type DmConversationMessageData struct {
	DmConversationId string `json:"dm_conversation_id"`
	DmEventId        string `json:"dm_event_id"`
}

type UserDirectMessageLookupResponse struct {
	MessageData *UserDirectMessageData `json:"data"`
	RateLimit   *RateLimit
}

type UserDirectMessageCreateResponse struct {
	MessageData *DmConversationMessageData `json:"data"`
	RateLimit   *RateLimit
}
