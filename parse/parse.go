package parse

import (
	"bufio"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pacur/pacur/pack"
	"github.com/pacur/pacur/utils"
)

const (
	root      = "/pacur_build"
	blockList = 1
	blockFunc = 2
)

var (
	itemReg = regexp.MustCompile("(\"[^\"]+\")|(`[^`]+`)")
)

func File(distro, release, home string) (pac *pack.Pack, err error) {
	home, err = filepath.Abs(home)
	if err != nil {
		return
	}

	err = utils.ExistsMakeDir(root)
	if err != nil {
		return
	}

	err = utils.CopyFiles(home, root, false)
	if err != nil {
		return
	}
	path := filepath.Join(root, "PKGBUILD")

	pac = &pack.Pack{
		Distro:     distro,
		Release:    release,
		Root:       root,
		Home:       home,
		SourceDir:  filepath.Join(root, "src"),
		PackageDir: filepath.Join(root, "pkg"),
	}

	pac.Init()

	file, err := utils.Open(path)
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
		n++

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
			if strings.Contains(line, "() {") {
				blockType = blockFunc
				blockKey = strings.Split(line, "() {")[0]
			} else {
				parts := strings.SplitN(line, "=", 2)
				if len(parts) != 2 {
					return
				}

				key := parts[0]
				val := parts[1]

				if key[:1] == " " {
					return
				} else if key[len(key)-1:] == " " {
					return
				}

				valLen := len(val)
				switch val[:1] {
				case `"`, "`":
					if val[valLen-1:] != val[:1] {
						return
					}

					err = pac.AddItem(key, val[1:valLen-1], n, line)
					if err != nil {
						return
					}
				case "(":
					if val[valLen-1:] == ")" {
						if val[1:2] != `"` && val[1:2] != "`" {
							return
						}

						if val[valLen-2:valLen-1] != val[1:2] {
							return
						}

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
					return
				default:
					return
				}
			}
		}
	}

	return
}
