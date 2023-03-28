package internal_test

import (
	"fmt"
	"testing"

	"github.com/ganvoa/chait/internal"
)

type FakeChatParticipant struct {
	Name  string
	Count int
	Error error
}

func (fp *FakeChatParticipant) Talk() (*internal.ChatMessage, error) {
	if fp.Error != nil {
		return nil, fp.Error
	}

	fp.Count++

	return &internal.ChatMessage{Name: fp.Name, Message: fmt.Sprintf("Message #%d From %s", fp.Count, fp.Name)}, nil

}

func (fp *FakeChatParticipant) Reply(message *internal.ChatMessage) (*internal.ChatMessage, error) {
	if fp.Error != nil {
		return nil, fp.Error
	}

	fp.Count++

	return &internal.ChatMessage{Name: fp.Name, Message: fmt.Sprintf("Message #%d From %s", fp.Count, fp.Name)}, nil

}

func Test_CreateChatRoom(t *testing.T) {

	t.Run("maxMessages must be at least 1", func(t *testing.T) {
		_, err := internal.NewChatRoom(0, nil, nil)
		if err == nil {
			t.Fatal("expected error, got none")
		}
	})

	t.Run("chat should stop on maxMessages", func(t *testing.T) {

		maxMessages := 10

		p1 := &FakeChatParticipant{
			Name:  "P1",
			Error: nil,
			Count: 0,
		}

		p2 := &FakeChatParticipant{
			Name:  "P2",
			Error: nil,
			Count: 0,
		}

		cr, err := internal.NewChatRoom(maxMessages, p1, p2)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		err = cr.Talk()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if cr.MessageCount != maxMessages {
			t.Fatalf("expected MessageCount to be %d, got %d", maxMessages, cr.MessageCount)
		}

		if len(cr.Conversation) != maxMessages*2 {
			t.Fatalf("expected Conversation length to be %d, got %d", maxMessages*2, len(cr.Conversation))
		}

		firstExpectedMessage := fmt.Sprintf("Message #%d From %s", 1, p1.Name)
		if cr.Conversation[0].Message != firstExpectedMessage {
			t.Fatalf("expected First Message to be %s, got %s", firstExpectedMessage, cr.Conversation[0].Message)
		}

		lastExpectedMessage := fmt.Sprintf("Message #%d From %s", 10, p2.Name)
		if cr.Conversation[len(cr.Conversation)-1].Message != lastExpectedMessage {
			t.Fatalf("expected last Message to be %s, got %s", lastExpectedMessage, cr.Conversation[len(cr.Conversation)-1].Message)
		}
	})
}
