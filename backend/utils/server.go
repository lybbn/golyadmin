// +-------------------------------------------------------------------
// | program: golyadmin
// +-------------------------------------------------------------------
// | Author: lybbn
// +-------------------------------------------------------------------
// | QQ: 1042594286
// +-------------------------------------------------------------------
// | Version: 1.0
// +-------------------------------------------------------------------
// | Date: 2023/08/21
// +-------------------------------------------------------------------

// ------------------------------
// monitor系统命令封装
// ------------------------------
package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"time"

	"gitee.com/lybbn/golyadmin/global"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type LySystemMonitorInfo struct {
	IsWindows   bool            `json:"is_windows"`
	Time        string          `json:"time"`
	System      string          `json:"system"`
	Mem         memInfo         `json:"mem"`
	Disk        []diskInfo      `json:"disk"`
	Cpu         []interface{}   `json:"cpu"`
	LoadAverage loadAverageInfo `json:"load_average"`
	Network     networkInfo     `json:"network"`
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

type loadAverageInfo struct {
	One     float64 `json:"one"`
	Five    float64 `json:"five"`
	Fifteen float64 `json:"fifteen"`
	Max     int     `json:"max"`
	Limit   int     `json:"limit"`
	Safe    float64 `json:"safe"`
	Percent float64 `json:"percent"`
}

type networkInfo struct {
	Network     interface{} `json:"network"`
	UpTotal     uint64      `json:"upTotal"`
	DownTotal   uint64      `json:"downTotal"`
	Up          uint64      `json:"up"`
	Down        uint64      `json:"down"`
	DownPackets uint64      `json:"downPackets"`
	UpPackets   uint64      `json:"upPackets"`
	Iostat      interface{} `json:"iostat"`
}

type networkInnerInfo struct {
	UpTotal     uint64 `json:"upTotal"`
	DownTotal   uint64 `json:"downTotal"`
	Up          uint64 `json:"up"`
	Down        uint64 `json:"down"`
	DownPackets uint64 `json:"downPackets"`
	UpPackets   uint64 `json:"upPackets"`
}

// 获取系统信息
func LyGetSystemAllInfo() LySystemMonitorInfo {
	var data LySystemMonitorInfo
	data.Mem = GetMemInfo()
	data.LoadAverage = GetLoadAverage()
	data.Network = GetNetWork()
	data.Cpu = GetCpuInfo()
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

func GetCpuInfo() (c []interface{}) {
	if cpuInfos, err := cpu.Info(); err != nil {
		return c
	} else {
		cpuCount, _ := cpu.Counts(true) //false cpu物理核心 、 true cpu逻辑核心数量， cpu.Counts(false) //runtime.NumCPU()获取机器的CPU核心数(逻辑处理器)
		cpuNum, _ := cpu.Counts(false)
		cpus, _ := cpu.Percent(time.Duration(time.Second), true) //每秒刷新一次
		var used_total float64
		// 显示每个核心的使用率
		for i := 0; i < cpuCount; i++ {
			used_total = used_total + cpus[i]
		}
		used := FormatFloat64ToFloat64(used_total/float64(cpuCount), 1)
		c = append(c, used)
		c = append(c, cpuCount)
		c = append(c, cpus)
		c = append(c, cpuInfos[0].ModelName+" * "+FormatInt2String(len(cpuInfos)))
		cpuW := len(cpuInfos)
		tmp := 0
		if cpuW > 0 {
			tmp = cpuNum / cpuW
		}
		c = append(c, tmp)
		c = append(c, cpuW)
	}
	return c
}

func GetDiskInfo() (d []diskInfo) {
	if IsWindows() {
		diskParts, _ := disk.Partitions(true)
		for _, value := range diskParts {
			var dInfo diskInfo
			dInfo.Path = value.Mountpoint + "/"
			dInfo.Type = value.Fstype
			diskInfoms, _ := disk.Usage(value.Mountpoint)
			dInfo.Size = append(dInfo.Size, FormatInt2String(int(diskInfoms.Total)/GB)+"G")
			dInfo.Size = append(dInfo.Size, FormatInt2String(int(diskInfoms.Used)/GB)+"G")
			dInfo.Size = append(dInfo.Size, FormatInt2String(int(diskInfoms.Free)/GB)+"G")
			dInfo.Size = append(dInfo.Size, FormatFloat2String(float64(diskInfoms.UsedPercent), 1))
			d = append(d, dInfo)
		}
		return d
	} else {
		diskParts, _ := disk.Partitions(true)
		for _, value := range diskParts {
			var dInfo diskInfo
			dInfo.Path = value.Mountpoint
			dInfo.Type = value.Fstype
			diskInfoms, _ := disk.Usage(value.Mountpoint)
			dInfo.Size = append(dInfo.Size, FormatInt2String(int(diskInfoms.Total)/GB)+"G")
			dInfo.Size = append(dInfo.Size, FormatInt2String(int(diskInfoms.Used)/GB)+"G")
			dInfo.Size = append(dInfo.Size, FormatInt2String(int(diskInfoms.Free)/GB)+"G")
			dInfo.Size = append(dInfo.Size, FormatFloat2String(float64(diskInfoms.UsedPercent), 1))
			d = append(d, dInfo)
		}
		return d
	}

}

func GetLoadAverage() (l loadAverageInfo) {
	if !IsWindows() {
		loadavg, err := ioutil.ReadFile("/proc/loadavg") // 读取/proc/loadavg文件
		if err != nil {
			fields := strings.Fields(string(loadavg)) // 切分字符串
			l.One = FormatString2Float64(fields[0])
			l.Five = FormatString2Float64(fields[1])
			l.Fifteen = FormatString2Float64(fields[2])
		} else {
			loadInfo, _ := load.Avg()
			l.One = loadInfo.Load1
			l.Five = loadInfo.Load5
			l.Fifteen = loadInfo.Load15
		}
	} else {
		loadInfo, _ := load.Avg()
		l.One = loadInfo.Load1
		l.Five = loadInfo.Load5
		l.Fifteen = loadInfo.Load15
	}
	cpuCount, _ := cpu.Counts(true)
	l.Max = cpuCount * 2
	l.Limit = cpuCount * 2
	l.Safe = float64(cpuCount*2) * 0.75
	var temppercent float64
	if l.Max != 0 {
		temppercent = FormatFloat64ToFloat64(float64(l.One)/float64(l.Max)*100, 2)
	}
	if temppercent > 100 {
		temppercent = 100
	}
	l.Percent = temppercent
	return l
}

var ctx = context.Background()

func GetNetWork() (n networkInfo) {
	otime_key := "golyadmin_otime"
	redisExpiration := time.Second * 86400
	otime, _ := global.GL_REDIS.Get(ctx, otime_key).Result()
	ntime := time.Now().Unix()
	netInfos, _ := net.IOCounters(true)
	var networks = make(map[string]interface{})
	for _, value := range netInfos {
		if strings.Contains(value.Name, "Loopback") || strings.Contains(value.Name, "Teredo") {
			continue
		}
		up_key := fmt.Sprintf("%v_up", value.Name)
		down_key := fmt.Sprintf("%v_down", value.Name)
		if otime == "" {
			otime = FormatInt642String(time.Now().Unix())
			global.GL_REDIS.Set(ctx, up_key, value.BytesSent, redisExpiration)
			global.GL_REDIS.Set(ctx, down_key, value.BytesRecv, redisExpiration)
			global.GL_REDIS.Set(ctx, otime_key, otime, redisExpiration)
		}
		up, _ := global.GL_REDIS.Get(ctx, up_key).Result()
		down, _ := global.GL_REDIS.Get(ctx, down_key).Result()
		if up == "" {
			up = FormatUint2String(value.BytesSent)
		}
		if down == "" {
			down = FormatUint2String(value.BytesRecv)
		}
		var ntw networkInnerInfo
		ntw.UpTotal = value.BytesSent
		ntw.DownTotal = value.BytesRecv
		ntw.UpPackets = value.PacketsSent
		ntw.DownPackets = value.PacketsRecv
		internels := uint64(ntime - FormatString2Int64(otime))
		if internels == 0 {
			ntw.Up = 0
			ntw.Down = 0
		} else {
			ntw.Up = (value.BytesSent - FormatString2Uint64(up)) / 1024 / internels
			ntw.Down = (value.BytesRecv - FormatString2Uint64(down)) / 1024 / internels
		}

		networks[value.Name] = ntw
		n.UpTotal = n.UpTotal + value.BytesSent
		n.DownTotal = n.DownTotal + value.BytesRecv
		n.UpPackets = n.UpPackets + value.PacketsSent
		n.DownPackets = n.DownPackets + value.PacketsRecv
		n.Up = n.Up + ntw.Up
		n.Down = n.Down + ntw.Down

		global.GL_REDIS.Set(ctx, up_key, value.BytesSent, redisExpiration)
		global.GL_REDIS.Set(ctx, down_key, value.BytesRecv, redisExpiration)
		global.GL_REDIS.Set(ctx, otime_key, FormatInt642String(time.Now().Unix()), redisExpiration)

	}
	n.Network = networks
	return n
}
