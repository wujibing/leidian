package leidian

import (
	"context"
	"testing"
)

func TestSimulator_RunApp(t *testing.T) {
	SetPath("F:\\LDPlayer\\LDPlayer64")
	simulator := &Simulator{
		Index: 0,
	}
	if err := simulator.RunApp(context.Background(), "com.sy.dldlhsdj.azt"); err != nil {
		t.Fatal(err)
	} else {
		t.Log("ok")
	}
}
