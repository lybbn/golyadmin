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
	IsWindows bool       `json:"is_windows"`
	Time      string     `json:"time"`
	System    string     `json:"system"`
	Mem       memInfo    `json:"mem"`
	Disk      []diskInfo `json:"disk"`
}

type diskInfo struct {
	Path       string   `json:"path"`
	Filesystem string   `json:"filesystem"`
	Type       string   `json:"type"`
	Size       []string `json:"size"`
	Inodes     []string `json:"inodes"`
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
	data.Mem = GetMemInfo()
	// data['load_average'] = self.GetLoadAverage()
	// data['network'] = self.GetNetWork()
	// data['cpu'] = self.GetCpuInfo(1)
	data.Disk = GetDiskInfo()
	data.Time = GetBootTime()
	data.System = GetSystemVersion()
	data.IsWindows = IsWindows()
	return data
}

func IsWindows() bool {
	plat := strings.ToLower(runtime.GOOS)
	if plat == "windows" {
		return true
	}
	return false
}

func GetBootTime() string {
	boottime, _ := host.BootTime()
	nowtime := time.Now().Unix()
	boottimeUnix := time.Unix(int64(boottime), 0).Unix()
	runDays := int((nowtime - boottimeUnix) / 86400)
	result := fmt.Sprintf("%d天", runDays)
	return result
}

func GetSystemVersion() string {
	n, _ := host.Info()
	plat := n.Platform
	platV := n.PlatformVersion
	if IsWindows() {
		platArr := strings.Fields(plat)
		plat = platArr[1] + " " + platArr[2]
		platVArr := strings.Split(platV, "Build")
		platV = "Build" + platVArr[1]
	}

	info := fmt.Sprintf("%v %v", plat, platV)
	sysversion := info + "(" + runtime.Version() + ")"
	return sysversion
}

func GetMemInfo() (m memInfo) {
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

func GetDiskInfo() (d []diskInfo) {
	if IsWindows() {
		diskParts, _ := disk.Partitions(true)
		for _, value := range diskParts {
			var dInfo diskInfo
			dInfo.Path = value.Mountpoint + "/"
			dInfo.Filesystem = value.Fstype
			diskInfoms, _ := disk.Usage(value.Mountpoint)
			dInfo.Size = append(dInfo.Size, FormatInt2String(int(diskInfoms.Total)/GB))
			dInfo.Size = append(dInfo.Size, FormatInt2String(int(diskInfoms.Used)/GB))
			dInfo.Size = append(dInfo.Size, FormatInt2String(int(diskInfoms.Free)/GB))
			dInfo.Size = append(dInfo.Size, FormatFloat2String(float64(diskInfoms.UsedPercent), 1))
			d = append(d, dInfo)
		}
		return d
	} else {
		return d
	}

}

func GetLoadAverage() {

}
