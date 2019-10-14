package cmd

import (
	"testing"

	"github.com/shirou/gopsutil/cpu"
)

type TimesStats []cpu.TimesStat

func TestCalcCPUPercent(t *testing.T) {
	cpu := cpu.TimesStat{
		CPU:"test",
		User:1.0,
		System:1.0,
		Idle:1.0,
		Nice:1.0,
		Iowait:1.0,
		Irq:1.0,
		Softirq:1.0,
		Steal:1.0,
		Guest:1.0,
		GuestNice:1.0,
	}

	// コードに合わせてテストデータを生成
	var testData TimesStats
	testData = append(testData, cpu)

	result := CalcCPUPercent(testData)
	expext := 90.0
	if result != expext {
		t.Error("\n実際： ", result, "\n理想： ", expext)
	}

	t.Log(cpu.CPU + "Test終了")
}