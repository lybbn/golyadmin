// +-------------------------------------------------------------------
// | program: golyadmin
// +-------------------------------------------------------------------
// | Author: lybbn
// +-------------------------------------------------------------------
// | QQ: 1042594286
// +-------------------------------------------------------------------

// ------------------------------
// monitor系统命令封装
// ------------------------------
package utils

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type LySystemMonitorInfo struct {
	IsWindows bool    `json:"is_windows"`
	Time      string  `json:"time"`
	System    string  `json:"system"`
	Mem       memInfo `json:"mem"`
}

type diskInfo struct {
	UsedMB      int `json:"usedMb"`
	UsedGB      int `json:"usedGb"`
	TotalMB     int `json:"totalMb"`
	TotalGB     int `json:"totalGb"`
	UsedPercent int `json:"usedPercent"`
}

type memInfo struct {
	Free    float64 `json:"free"`
	Percent float64 `json:"percent"`
	Total   float64 `json:"total"`
	Used    float64 `json:"used"`
}

// 获取系统信息
func LyGetSystemAllInfo() LySystemMonitorInfo {
	var data LySystemMonitorInfo
	data.Mem = getMemInfo()
	// data['load_average'] = self.GetLoadAverage()
	// data['network'] = self.GetNetWork()
	// data['cpu'] = self.GetCpuInfo(1)
	// data['disk'] = self.GetDiskInfo()
	data.Time = getBootTime()
	data.System = getSystemVersion()
	data.IsWindows = isWindows()
	return data
}

func isWindows() bool {
	plat := strings.ToLower(runtime.GOOS)
	if plat == "windows" {
		return true
	}
	return false
}

func getBootTime() string {
	boottime, _ := host.BootTime()
	nowtime := time.Now().Unix()
	boottimeUnix := time.Unix(int64(boottime), 0).Unix()
	runDays := int((nowtime - boottimeUnix) / 86400)
	result := fmt.Sprintf("%d天", runDays)
	return result
}

func getSystemVersion() string {
	n, _ := host.Info()
	plat := n.Platform
	platV := n.PlatformVersion
	if isWindows() {
		platArr := strings.Fields(plat)
		plat = platArr[1] + " " + platArr[2]
		platVArr := strings.Split(platV, "Build")
		platV = "Build" + platVArr[1]
	}

	info := fmt.Sprintf("%v %v", plat, platV)
	sysversion := info + "(" + runtime.Version() + ")"
	return sysversion
}

func getMemInfo() (m memInfo) {
	if u, err := mem.VirtualMemory(); err != nil {
		return m
	} else {
		m.Free = float64(u.Free / MB)
		m.Used = float64(u.Used / MB)
		m.Total = float64(u.Total / MB)
		m.Percent = u.UsedPercent
	}
	return m
}

func getDiskInfo() (d diskInfo) {
	if u, err := disk.Usage("/"); err != nil {
		return d
	} else {
		d.UsedMB = int(u.Used) / MB
		d.UsedGB = int(u.Used) / GB
		d.TotalMB = int(u.Total) / MB
		d.TotalGB = int(u.Total) / GB
		d.UsedPercent = int(u.UsedPercent)
	}
	return d
}
