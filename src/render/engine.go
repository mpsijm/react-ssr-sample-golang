package render

import (
	"fmt"
	"os"
	"strings"

	"github.com/mpsijm/go-node"
)

type Engine struct {
	scriptContents   string
	templateContents string
	vm               node.VM
}

func NewEngine(scriptLocation, templateLocation string) *Engine {

	e := new(Engine)
	e.scriptContents = loadFileContents(scriptLocation)
	e.templateContents = loadFileContents(templateLocation)

	// TODO All of the below might panic
	e.vm = node.New(nil)

	if _, err := e.vm.Run(e.scriptContents); err != nil {
		panic(err)
	}

	return e
}

func loadFileContents(filePath string) string {

	scriptContents, err := os.ReadFile(filePath)
	if err != nil {
		// TODO Handle this properly
		fmt.Println(err)
		panic(err)
	}
	return string(scriptContents)
}

// We create and teardown a new Node VM, bit wasteful but likely not thread safe otherwise
func (e *Engine) Render(currentPath string, serverSideState string) string {

	// TODO All of the below might panic
	result, err := e.vm.Run(fmt.Sprintf("global.render(`%s`, `%s`, `%s`)",
		strings.ReplaceAll(e.templateContents, "`", "\\`"),
		strings.ReplaceAll(currentPath, "`", "\\`"),
		strings.ReplaceAll(serverSideState, "`", "\\`"),
	))

	if err != nil {
		panic(err)
	}

	return result
}
