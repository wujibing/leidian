package leidian

import (
	"bufio"
	"io"
)

type (
	Packages []string
)

//查找包名
func (p *Packages) Find(packageName string) bool {
	for _, name := range *p {
		if name == packageName {
			return true
		}
	}
	return false
}

//解析包名
func NewPackages(reader io.Reader) (*Packages, error) {
	rd := bufio.NewReader(reader)
	packages := make(Packages, 0, 1)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		} else {
			packages = append(packages, string(line[8:len(line)-1]))
		}
	}
	return &packages, nil
}
