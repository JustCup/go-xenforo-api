package api // import "github.com/JustCup/go-xenforo-api/api"
import (
	"fmt"

	"github.com/JustCup/go-xenforo-api/object"
)

// ConversationMessagesSend function.
//
// Replies to a conversation
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_conversation_messages_
func (xf *XF) ConversationMessagesSend(params Params) (response object.ConversationMessagesResponse, err error) {
	err = xf.RequestUnmarshal("POST", "conversation-messages", params, &response)
	return
}

// ConversationMessagesGet function.
//
// Gets the specified conversation message.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_conversation_messages_
func (xf *XF) ConversationMessagesGet(id int) (response object.ConversationMessagesResponse, err error) {
	err = xf.RequestUnmarshal("GET", fmt.Sprintf("conversation-messages/%d/", id), Params{}, &response)
	return
}

// ConversationsGet function.
//
// Gets the API user's list of conversations.
//
// https://xenforo.com/community/pages/api-endpoints/#route_get_conversations_
func (xf *XF) ConversationsGet(params Params) (response object.ConversationsResponse, err error) {
	err = xf.RequestUnmarshal("GET", "conversations", params, &response)
	return
}

// ConversationsCreate function.
//
// Creates a conversation.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_conversations_
func (xf *XF) ConversationsCreate(params Params) (response object.ConversationsResponse, err error) {
	err = xf.RequestUnmarshal("POST", "conversations", params, &response)
	return
}

// ConversationsGetID function.
//
// Gets information about the specified conversation.
//
// https://xenforo.com/community/pages/api-endpoints/#route_get_conversations_id_
func (xf *XF) ConversationsGetID(id int) (response object.ConversationsResponse, err error) {
	err = xf.RequestUnmarshal("GET", fmt.Sprintf("conversations/%d/", id), Params{}, &response)
	return
}
