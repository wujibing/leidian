package leidian

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

//运行app
func (s *Simulator) RunApp(packageName string) error {
	return RunApp(s.Index, packageName)
}

func (s *Simulator) InstallApp(filename string) {
	InstallApp(s.Index, filename)
}

func (s *Simulator) KillApp(packageName string) {
	KillApp(s.Index, packageName)
}
