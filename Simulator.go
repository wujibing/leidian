package leidian

import (
	"context"
	"time"
)

type (
	Simulator struct {
		Index     int
		Name      string
		Pid       int
		RenderPid int
		Running   bool
		VboxPid   int
	}
)

func (s *Simulator) Quit() error {
	return Quit(s.Index)
}

//还原并重置硬件信息
func (s *Simulator) Restore(path, cpu, memory string) {
	s.Quit()
	time.Sleep(time.Second * 2) //休息2秒
	Restore(s.Index, path)
	Modify(s.Index, "--resolution", "960,540,160", "--cpu", cpu, "--memory", memory, "--imei", "auto", "--imsi", "auto", "--simserial", "auto", "--androidid", "autp", "--mac", "auto")
}

//运行app
func (s *Simulator) RunApp(ctx context.Context, packageName string) error {
	if err := RunApp(s.Index, packageName); err != nil {
		return err
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Minute):
			if err := RunApp(s.Index, packageName); err != nil {
				return err
			}
		case <-time.After(time.Millisecond * 100):
			if GetTopPackageName(s.Index) == packageName {
				return nil
			}
		}
	}
}

//安装apk 一直等待能查找到包名为止
func (s *Simulator) InstallApp(ctx context.Context, packageName, filename string) error {
	InstallApp(s.Index, filename)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Millisecond * 300):
			packages, err := GetPackages(s.Index)
			if err == nil && packages.Find(packageName) {
				return nil
			}
		}
	}
}

func (s *Simulator) KillApp(packageName string) {
	KillApp(s.Index, packageName)
}
