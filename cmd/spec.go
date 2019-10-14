package cmd

import (
	"fmt"
	"math"

	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
)

func cmd()  {
	mem, cpu, load := getMachineSpec()

	fmt.Println(mem)
	fmt.Println(cpu)
	fmt.Println(load)
}

func getMachineSpec() (float64, float64, float64){
	memory, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}

	// 全コアの情報をまとめて取得する
	cpu, err := cpu.Times(false)
	if err != nil {
		panic(err)
	}

	cpuPersent := CalcCpuPercent(cpu)

	// ロードアベレージを取得
	load, err := load.Avg()
	if err != nil {
		panic(err)
	}

	return memory.UsedPercent, cpuPersent, load.Load1
}

/**
ここの処理を移植
https://github.com/shirou/gopsutil/blob/1c09419d4b1c4c19e06d9f48b9406bda27d32acd/cpu/cpu.go#L106
*/
func CalcCpuPercent(cpu []cpu.TimesStat) (float64) {
	busy := cpu[0].User + cpu[0].System + cpu[0].Nice + cpu[0].Iowait + cpu[0].Irq +
		cpu[0].Softirq + cpu[0].Steal + cpu[0].Guest + cpu[0].GuestNice

	return math.Min(100, math.Max(0, busy /(busy + cpu[0].Idle) * 100))
}