package pablo

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

const (
	MainSection = "Main"
)

type Section struct {
	Name        string
	Files       []string
	Vakue       string
	subSections []*Section
}

type Book struct {
}

func LoadSections(src string) ([]*Section, error) {
	var result []*Section
	mainSect := &Section{
		Name: MainSection,
	}
	infos, err := ioutil.ReadDir(src)
	if err != nil {
		return nil, err
	}
	for _, v := range infos {
		fPath := filepath.Join(src, v.Name())
		if v.IsDir() {
			sub, err := LoadSection(fPath)
			if err != nil {
				return nil, err
			}
			result = append(result, sub)
			continue
		}
		mainSect.Files = append(mainSect.Files, fPath)
	}
	return result, nil
}

func LoadSection(dir string) (*Section, error) {
	sect := &Section{}
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	sect.Name = dir
	for _, v := range infos {
		fPath := filepath.Join(dir, v.Name())
		if v.IsDir() {
			sub, err := LoadSection(fPath)
			if err != nil {
				return nil, err
			}
			sect.subSections = append(sect.subSections, sub)
			continue
		}
		sect.Files = append(sect.Files, fPath)
		fmt.Printf("%s : %s \n", sect.Name, fPath)
	}
	return sect, nil
}
