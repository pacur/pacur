package parse

import (
	"bufio"
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/pack"
	"os"
	"regexp"
	"strings"
)

const (
	blockList = 1
	blockFunc = 2
)

var (
	itemReg = regexp.MustCompile("(\"[^\"]+\")|(`[^`]+`)")
)

func File(path string) (pac *pack.Pack, err error) {
	pac = &pack.Pack{}

	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	n := 0
	blockType := 0
	blockKey := ""
	blockData := ""
	blockItems := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		n += 1

		if line == "" || line[:1] == "#" {
			continue
		}

		if blockType == blockList {
			if line == ")" {
				for _, item := range itemReg.FindAllString(blockData, -1) {
					blockItems = append(blockItems, item[1:len(item)-1])
				}
				err = pac.AddItem(blockKey, blockItems, n, line)
				if err != nil {
					return
				}
				blockType = 0
				blockKey = ""
				blockData = ""
				blockItems = []string{}
				continue
			}

			blockData += strings.TrimSpace(line)
		} else if blockType == blockFunc {
			if line == "}" {
				err = pac.AddItem(blockKey, blockItems, n, line)
				if err != nil {
					return
				}
				blockType = 0
				blockKey = ""
				blockItems = []string{}
				continue
			}

			blockItems = append(blockItems, strings.TrimSpace(line))
		} else {
			switch line {
			case "build() {":
				blockType = blockFunc
				blockKey = "build"
			case "package() {":
				blockType = blockFunc
				blockKey = "package"
			case "preinst() {":
				blockType = blockFunc
				blockKey = "preinst"
			case "postinst() {":
				blockType = blockFunc
				blockKey = "postinst"
			case "prerm() {":
				blockType = blockFunc
				blockKey = "prerm"
			case "postrm() {":
				blockType = blockFunc
				blockKey = "postrm"
			default:
				parts := strings.SplitN(line, "=", 2)
				if len(parts) != 2 {
					err = &SyntaxError{
						errors.Newf("parse: Line missing '=' (%d: %s)",
							n, line),
					}
					return
				}

				key := parts[0]
				val := parts[1]

				if key[:1] == " " {
					err = &SyntaxError{
						errors.Newf("parse: Extra space padding (%d: %s)",
							n, line),
					}
					return
				} else if key[len(key)-1:] == " " {
					err = &SyntaxError{
						errors.Newf("parse: Extra space before '=' (%d: %s)",
							n, line),
					}
					return
				}

				switch val[:1] {
				case `"`, "`":
					err = pac.AddItem(key, val[1:len(val)-1], n, line)
					if err != nil {
						return
					}
				case "(":
					if val[len(val)-1:] == ")" {
						val = val[2 : len(val)-2]
						err = pac.AddItem(key, []string{val}, n, line)
						if err != nil {
							return
						}
					} else {
						blockType = blockList
						blockKey = key
					}
				case " ":
					err = &SyntaxError{
						errors.Newf("parse: Extra space after '=' (%d: %s)",
							n, line),
					}
					return
				default:
					err = &SyntaxError{
						errors.Newf("parse: Unexpected char '%s' expected "+
							"'\"' or '`' (%d: %s)", val[:1], n, line),
					}
					return
				}
			}
		}
	}

	return
}
