package leidian

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var (
	leidianPath string
	consolePath string
	simulators  = make(map[int]*Simulator)
	lock        sync.RWMutex
)

func SetPath(path string) {
	leidianPath = path
	consolePath = filepath.Join(path, "ldconsole.exe")

}

// 运行模拟器
func Launch(index int) error {
	_, err := run("launch", "--index", fmt.Sprintf("%d", index))
	return err
}

// 退出模拟器
func Quit(index int) error {
	_, err := run("quit", "--index", fmt.Sprintf("%d", index))
	return err
}

func QuitAll() {
	run("quitall")
}

func Remove(index int) {
	run("remove", "--index", fmt.Sprintf("%d", index))
}

func Copy(index int) {
	run("copy", "--from", fmt.Sprintf("%d", index))
}

func Add() {
	run("add")
}

func Sort() {
	run("sortWnd")
}

func Restore(index int, file string) {
	run("restore", "--index", fmt.Sprintf("%d", index), "--file", file)
}

//备份
func Backup(index int, file string) {
	run("backup", "--index", fmt.Sprintf("%d", index), "--file", file)
}

func Modify(index int, values ...string) {
	newValues := []string{"modify", "--index", fmt.Sprintf("%d", index)}
	newValues = append(newValues, values...)
	run(newValues...)
}

func InstallApp(index int, filename string) {
	run("installapp", "--index", fmt.Sprintf("%d", index), "--filename", filename)
}

func KillApp(index int, packageName string) {
	run("killapp", "--index", fmt.Sprintf("%d", index), "--packagename", packageName)
}
func RunApp(index int, packageName string) error {
	_, err := run("runapp", "--index", fmt.Sprintf("%d", index), "--packagename", packageName)
	return err
}

func GetSimulator(index int) *Simulator {
	lock.RLock()
	defer lock.RUnlock()
	return simulators[index]
}

func LoadSimulator() error {
	reader, err := run("list2")
	if err != nil {
		return err
	}
	lock.Lock()
	defer lock.Unlock()
	rd := bufio.NewReader(reader)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		lines := strings.Split(string(line), ",")
		index, _ := strconv.Atoi(lines[0])
		if simulators[index] == nil {
			simulators[index] = new(Simulator)
		}
		temp := simulators[index]
		temp.Index = index
		temp.Name = lines[1]
		temp.Pid, _ = strconv.Atoi(lines[2])
		temp.RenderPid, _ = strconv.Atoi(lines[3])
		temp.Running = lines[4] == "1"
		temp.VboxPid, _ = strconv.Atoi(lines[6])
	}
	return nil
}

func run(arg ...string) (io.Reader, error) {
	cmd := exec.Command(consolePath, arg...)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	err = cmd.Wait()
	return buf, err
}
