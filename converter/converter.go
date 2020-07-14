package converter

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/signavio/plantuml-converter/cmd"
	"io/ioutil"
	"log"
)

type PlantUml struct {
	files []PlantUmlFile
}

type PlantUmlFile struct {
	filePath       string
	fileContent    string
	updatedContent string
	blocks         []PlantUmlBlock
}

type PlantUmlBlock struct {
	// to last line number of content block
	// import to remember where to put the link
	lineNumber int
	// plant uml block
	content string
	// generated markdown link
	markdownLink string
}

// generates a link to the rendered png image for given input
func GenerateLink(input string) string {
	hash := hex.EncodeToString([]byte(input))
	return fmt.Sprintf("%s/plantuml/png/~h%s", cmd.PlantUmlServer, hash)
}

// set the markdown link
func (p *PlantUmlBlock) GenerateMarkdownLink() {
	p.markdownLink = GenerateLink(p.content)
}

// filter files list for given pattern (cmd.FilePattern) that should be converted
// also reads the file content into PlantUmlFile.fileContent
func (p *PlantUml) SetFiles() {
	var files []PlantUmlFile
	// find files matching the pattern
	// set fileContent and filePath for each PlantUmlFile
	p.files = files
}

// adding links and set PlantUmlFile.updatedContent
func (f *PlantUmlFile) SetUpdatedContent() {
	// you can always update the markdown file because the hash will be the same
	// if the content does not change
}

// parse the plant uml blocks from a file
func (f *PlantUmlFile) SetBlocks() error {
	var blocks []PlantUmlBlock

	// read lines from file
	bytesRead, _ := ioutil.ReadFile(f.filePath)
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")

	var hasStart bool = false

	var myBlock PlantUmlBlock
	fmt.Println(lines)

	for i := 0; i < len(lines); i++ {
		vline := lines[i]
		if vline == "@startuml" {
			hasStart = true
		} else if vline == "@enduml" {
			if !hasStart {
				return errors.New("Adrian is nen ...")
			}
			myBlock.lineNumber = i + 1
			blocks = append(blocks, myBlock)
			myBlock = PlantUmlBlock{}
			hasStart = false
		} else if hasStart {
			myBlock.content = myBlock.content + vline
		}
	}
	// find @startuml and @enduml and map to blocks

	f.blocks = blocks
	return errors.New("Not implemented")
}
