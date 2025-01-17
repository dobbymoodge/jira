package jiracmd

import (
	"github.com/coryb/figtree"
	"github.com/coryb/oreo"
	jira "github.com/go-jira/jira"
	"github.com/go-jira/jira/jiracli"
	"github.com/pkg/browser"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func CmdBrowseRegistry() *jiracli.CommandRegistryEntry {
	issue := ""

	return &jiracli.CommandRegistryEntry{
		"Open issue in browser",
		func(fig *figtree.FigTree, cmd *kingpin.CmdClause) error {
			cmd.Arg("ISSUE", "Issue to browse to").Required().StringVar(&issue)
			return nil
		},
		func(o *oreo.Client, globals *jiracli.GlobalOptions) error {
			return CmdBrowse(globals, issue)
		},
	}
}

// CmdBrowse open the default system browser to the provided issue
func CmdBrowse(globals *jiracli.GlobalOptions, issue string) error {
	return browser.OpenURL(jira.URLJoin(globals.Endpoint.Value, "browse", issue))
}
