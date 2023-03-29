package internal_test

import (
	"testing"

	"github.com/ganvoa/chait/internal"
	"github.com/stretchr/testify/assert"
)

type mockRenderer struct {
	expectedOutput string
	expectedError  error
	t              *testing.T
}

func (mr *mockRenderer) Write(m []byte) (int, error) {

	if mr.expectedError != nil {
		return 0, mr.expectedError
	}

	assert.Equal(mr.t, mr.expectedOutput, string(m))

	return 1, nil
}

func Test_TableRenderer(t *testing.T) {

	c := []internal.ChatMessage{
		{Name: "U1", Message: "Message 1"},
		{Name: "U2", Message: "Message 2"},
		{Name: "U1", Message: "Message 3"},
	}

	expectedTable :=
		`┌───────────────────────────────────────────────────────────────────────────────────────────────┐
│ Conversation                                                                                  │
├────────────┬──────────────────────────────────────────────────────────────────────────────────┤
│ USER       │ MESSAGE                                                                          │
├────────────┼──────────────────────────────────────────────────────────────────────────────────┤
│     U1     │ Message 1                                                                        │
├────────────┼──────────────────────────────────────────────────────────────────────────────────┤
│     U2     │ Message 2                                                                        │
├────────────┼──────────────────────────────────────────────────────────────────────────────────┤
│     U1     │ Message 3                                                                        │
└────────────┴──────────────────────────────────────────────────────────────────────────────────┘
roleU1: This is the roleU1
roleU2: This is the roleU2
iterations: 2
`

	mr := &mockRenderer{
		expectedOutput: expectedTable,
		expectedError:  nil,
		t:              t,
	}

	config := &internal.Config{}
	config.Chait.Iterations = 2
	config.Chait.RoleU1 = "This is the roleU1"
	config.Chait.RoleU2 = "This is the roleU2"

	r := internal.NewTableRenderer(c, config, mr)
	r.Render()

}
