package internal

import (
	"io"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type TableRenderer struct {
	writer       io.Writer
	conversation []ChatMessage
	config       *Config
}

func NewTableRenderer(c []ChatMessage, config *Config, w io.Writer) *TableRenderer {
	return &TableRenderer{conversation: c, writer: w, config: config}
}

func (cr *TableRenderer) Render() error {
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	tw := table.NewWriter()
	tw.SetStyle(table.StyleLight)
	tw.Style().Options.SeparateRows = true
	tw.SetTitle("Conversation")
	tw.SetCaption("roleU1: %s\nroleU2: %s\niterations: %d\n", cr.config.Chait.RoleU1, cr.config.Chait.RoleU2, cr.config.Chait.Iterations)
	tw.AppendSeparator()
	tw.AppendHeader(table.Row{"User", "Message"}, rowConfigAutoMerge)
	tw.SetColumnConfigs([]table.ColumnConfig{
		{Name: "User", WidthMin: 10, WidthMax: 10, Align: text.AlignCenter},
		{Name: "Message", WidthMin: 80, WidthMax: 80},
	})
	for _, v := range cr.conversation {
		tw.AppendRow(table.Row{v.Name, v.Message}, rowConfigAutoMerge)
	}

	render := tw.Render()
	_, err := cr.writer.Write([]byte(render))

	return err
}
