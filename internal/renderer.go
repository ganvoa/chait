package internal

import (
	"io"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type TableRenderer struct {
	writer       io.Writer
	conversation []ChatMessage
}

func NewTableRenderer(c []ChatMessage, w io.Writer) *TableRenderer {
	return &TableRenderer{conversation: c, writer: w}
}

func (cr *TableRenderer) Render() error {
	tw := table.NewWriter()
	tw.SetTitle("Conversation")
	tw.AppendHeader(table.Row{"User", "Message"})
	tw.SetStyle(table.StyleLight)
	tw.Style().Options.SeparateRows = true
	tw.SetColumnConfigs([]table.ColumnConfig{
		{Name: "User", WidthMin: 6, WidthMax: 6, Align: text.AlignCenter},
		{Name: "Message", WidthMin: 60, WidthMax: 60},
	})
	for _, v := range cr.conversation {
		tw.AppendRow(table.Row{v.Name, v.Message})
	}

	render := tw.Render()
	_, err := cr.writer.Write([]byte(render))

	return err
}
