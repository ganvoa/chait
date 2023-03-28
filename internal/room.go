package internal

import "errors"

type ChatParticipant interface {
	Talk() (*ChatMessage, error)
	Reply(message *ChatMessage) (*ChatMessage, error)
}

type ChatMessage struct {
	Name    string
	Message string
}

type ChatRoom struct {
	maxMessages       int
	secondParticipant ChatParticipant
	firstParticipant  ChatParticipant
	MessageCount      int
	Conversation      []ChatMessage
}

func NewChatRoom(maxMessages int, firstParticipant ChatParticipant, secondParticipant ChatParticipant) (*ChatRoom, error) {

	if maxMessages <= 0 {
		return nil, errors.New("maxMessages must be at least 1")
	}

	var messages []ChatMessage
	return &ChatRoom{
		maxMessages:       maxMessages,
		Conversation:      messages,
		secondParticipant: secondParticipant,
		firstParticipant:  firstParticipant,
	}, nil
}

func (cr *ChatRoom) StartConversation() error {
	cr.MessageCount = 0

	var scm *ChatMessage

	for {
		if cr.MessageCount == 0 {
			fcm, err := cr.firstParticipant.Talk()
			if err != nil {
				return err
			}
			cr.Conversation = append(cr.Conversation, *fcm)
			scm, err = cr.secondParticipant.Reply(fcm)
			if err != nil {
				return err
			}
			cr.Conversation = append(cr.Conversation, *scm)
		} else {
			fcm, err := cr.firstParticipant.Reply(scm)
			if err != nil {
				return err
			}
			cr.Conversation = append(cr.Conversation, *fcm)
			scm, err = cr.secondParticipant.Reply(fcm)
			if err != nil {
				return err
			}
			cr.Conversation = append(cr.Conversation, *scm)
		}

		cr.MessageCount++
		if cr.MessageCount >= cr.maxMessages {
			break
		}

	}

	return nil
}
