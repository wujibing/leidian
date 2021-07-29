package leidian

import "testing"

func TestKillApp(t *testing.T) {
	SetPath("F:\\LDPlayer\\LDPlayer64")
	KillApp(0, "com.sy.dldlhsdj.azt")
}
func TestRunApp(t *testing.T) {
	SetPath("F:\\LDPlayer\\LDPlayer64")
	if err := RunApp(0, "com.sy.dldlhsdj.azt"); err != nil {
		t.Fatal(err)
	} else {
		t.Log("ok")
	}
}

func TestGetSimulator(t *testing.T) {
	SetPath("F:\\LDPlayer\\LDPlayer64")
	LoadSimulator()
	t.Log(GetSimulator(0))
}

func TestGetTopPackageName(t *testing.T) {
	SetPath("F:\\LDPlayer\\LDPlayer64")
	t.Log(GetTopPackageName(0))
}

func TestLaunch(t *testing.T) {
	SetPath("F:\\LDPlayer\\LDPlayer64")
	Backup(0, "E:\\project\\golang\\gua\\douluo\\bin\\back.ldbk")
}

func TestSort(t *testing.T) {
	SetPath("F:\\LDPlayer\\LDPlayer64")
	Sort()
}

func TestCopy(t *testing.T) {
	SetPath("F:\\LDPlayer\\LDPlayer64")
	Copy(0)
}

func TestGetPackages(t *testing.T) {
	SetPath("F:\\LDPlayer\\LDPlayer64")
	packages, err := GetPackages(0)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(packages.Find("com.sy.dldlhsdj.azt"))
	}
}
