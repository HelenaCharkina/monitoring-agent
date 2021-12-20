package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func getStats() (*Stats, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("VirtualMemory error: ", err)
		return nil, err
	}

	// Disk - start from "/" mount point for Linux
	// might have to change for Windows!!
	// don't have a Window to test this out, if detect OS == windows
	// then use "\" instead of "/"

	diskStat, err := disk.Usage("/")
	if err != nil {
		fmt.Println("Disk.Usage error: ", err)
		return nil, err
	}

	// Cpu - get CPU number of cores and speed
	//cpuStat, err := Cpu.Info()
	//if err != nil {
	//	fmt.Println("Cpu.Info error: ", err)
	//	return nil, err
	//}
	cpuStats, err := cpu.Info()
	if err != nil {
		fmt.Println("Cpu.Info error: ", err)
		return nil, err
	}

	percentage, err := cpu.Percent(0, true)
	if err != nil {
		fmt.Println("Cpu.Percent error: ", err)
		return nil, err
	}
	hostStat, err := host.Info()
	if err != nil {
		fmt.Println("Host.Info error: ", err)
		return nil, err
	}

	stats := Stats{}

	stats.VmStat = VmStat{
		Total:       vmStat.Total,
		Free:        vmStat.Free,
		UsedPercent: vmStat.UsedPercent,
	}
	stats.Disk = Disk{
		Total:       diskStat.Total,
		Used:        diskStat.Used,
		Free:        diskStat.Free,
		UsedPercent: diskStat.UsedPercent,
	}
	stats.Cpu = Cpu{
		Percentage: percentage,
		Model:      cpuStats[0].ModelName,
		Cores:      int(cpuStats[0].Cores),
	}
	stats.Host = Host{
		Procs:           hostStat.Procs,
		OS:              hostStat.OS,
		Platform:        hostStat.Platform,
		PlatformVersion: hostStat.PlatformVersion,
	}

	return &stats, nil
}

type Stats struct {
	VmStat VmStat
	Disk   Disk
	Cpu    Cpu
	Host   Host
}
type Host struct {
	Procs           uint64
	OS              string
	PlatformVersion string
	Platform        string
}
type Cpu struct {
	Percentage []float64
	Model      string
	Cores      int
}
type Disk struct {
	Total       uint64
	Free        uint64
	Used        uint64
	UsedPercent float64
}

type VmStat struct {
	Total       uint64
	Free        uint64
	UsedPercent float64
}
