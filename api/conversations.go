package api // import "github.com/JustCup/xenforo-sdk/api"
import (
	"fmt"

	"github.com/JustCup/xenforo-sdk/object"
)

// ConversationMessagesSend function.
//
// Replies to a conversation.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_conversation_messages_
func (xf *XF) ConversationMessagesSend(params Params) (response object.ConversationsResponse, err error) {
	err = xf.RequestUnmarshal("POST", "conversation-messages", params, &response)
	return
}

// ConversationMessagesGet function.
//
// Gets the specified conversation message.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_conversation_messages_
func (xf *XF) ConversationMessagesGet(id int) (response object.ConversationsResponse, err error) {
	err = xf.RequestUnmarshal("GET", fmt.Sprintf("conversation-messages/%d/", id), Params{}, &response)
	return
}

// ConversationMessagesUpdate function.
//
// Updates the specified conversation message.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_conversation_messages_id_
func (xf *XF) ConversationMessagesUpdate(id int, params Params) (response object.ConversationsResponse, err error) {
	err = xf.RequestUnmarshal("POST", fmt.Sprintf("conversation-messages/%d/", id), params, &response)
	return
}

// ConversationMessagesReact function.
//
// Reacts to the specified conversation message.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_conversation_messages_id_react
func (xf *XF) ConversationMessagesReact(id int, params Params) (response object.ConversationsResponse, err error) {
	err = xf.RequestUnmarshal("POST", fmt.Sprintf("conversation-messages/%d/react", id), params, &response)
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
func (xf *XF) ConversationsGetID(id int, params Params) (response object.ConversationsResponse, err error) {
	err = xf.RequestUnmarshal("GET", fmt.Sprintf("conversations/%d/", id), params, &response)
	return
}

// ConversationsUpdate function.
//
// Updates the specified conversation.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_conversations_id_
func (xf *XF) ConversationsUpdate(id int, params Params) (response object.ConversationsResponse, err error) {
	err = xf.RequestUnmarshal("POST", fmt.Sprintf("conversations/%d/", id), params, &response)
	return
}

// ConversationsUpdate function.
//
// Deletes the specified conversation from the API user's list. Does not delete the conversation for other receivers.
//
// https://xenforo.com/community/pages/api-endpoints/#route_delete_conversations_id_
func (xf *XF) ConversationsDelete(id int, params Params) (response object.ConversationsResponse, err error) {
	err = xf.RequestUnmarshal("POST", fmt.Sprintf("conversations/%d/", id), params, &response)
	return
}

// ConversationsInvite function.
//
// Invites the specified users to this conversation.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_conversations_id_invite
func (xf *XF) ConversationsInvite(id int, params Params) (response object.ConversationsResponse, err error) {
	err = xf.RequestUnmarshal("POST", fmt.Sprintf("conversations/%d/invite", id), params, &response)
	return
}

// ConversationsMarkRead function.
//
// Marks the conversation as read up until the specified time. This cannot move the
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_conversations_id_mark_read
func (xf *XF) ConversationsMarkRead(id int, params Params) (response object.ConversationsResponse, err error) {
	err = xf.RequestUnmarshal("POST", fmt.Sprintf("conversations/%d/mark-read", id), params, &response)
	return
}

// ConversationsMarkUnread function.
//
// Marks a conversation as unread. This will mark all messages in the conversation as unread.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_conversations_id_mark_unread
func (xf *XF) ConversationsMarkUnread(id int) (response object.ConversationsResponse, err error) {
	err = xf.RequestUnmarshal("POST", fmt.Sprintf("conversations/%d/mark-unread", id), Params{}, &response)
	return
}

// ConversationsMessagesGet function.
//
// Gets a page of messages in the specified conversation.
//
// https://xenforo.com/community/pages/api-endpoints/#route_get_conversations_id_messages
func (xf *XF) ConversationsMessagesGet(id int, params Params) (response object.ConversationsResponse, err error) {
	err = xf.RequestUnmarshal("GET", fmt.Sprintf("conversations/%d/messages", id), params, &response)
	return
}

// ConversationsStar function.
//
// Sets the star status of the specified conversation.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_conversations_id_star
func (xf *XF) ConversationsStar(id int, params Params) (response object.ConversationsResponse, err error) {
	err = xf.RequestUnmarshal("POST", fmt.Sprintf("conversations/%d/star", id), params, &response)
	return
}
