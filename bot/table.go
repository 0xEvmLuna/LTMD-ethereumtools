package bot

import "github.com/cheynewallace/tabby"

func Table() {
	t := tabby.New()
	t.AddHeader("NAME", "TITLE", "DEPARTMENT")
	t.AddLine("John Smith", "Developer", "Engineering")
	t.Print()

}
