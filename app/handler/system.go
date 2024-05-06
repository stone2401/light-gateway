package handler

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/stone2401/light-gateway/app/middleware"
	"github.com/stone2401/light-gateway/app/model/dto"
)

type OsStat struct {
	stat *dto.Stat
	mu   sync.RWMutex
}

func RegisterSystem(router *gin.RouterGroup) {
	osStat := new(OsStat)
	go osStat.WatchStat()
	router.GET("/sse/:uid", SystemSse)
	router.GET("/serve/stat", osStat.ServeStat)
}

func SystemSse(ctx *gin.Context) {
	w := ctx.Writer

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		middleware.ResponseError(ctx, 2000, errors.New("streaming not supported"))
		return
	}
	i := 1
	for {
		time.Sleep(5 * time.Second)
		w.Write([]byte(fmt.Sprintf("id: %d\ndata: %s\n\n", i, "{'type': 'ping'}")))
		i++
		flusher.Flush()
	}
}

func (o *OsStat) ServeStat(ctx *gin.Context) {
	if o.stat == nil {
		o.UpdateStat()
	}
	o.mu.RLock()
	defer o.mu.RUnlock()
	middleware.ResponseSuccess(ctx, o.stat)
}

func (o *OsStat) WatchStat() {
	// 每5秒获取一次
	// 1. 获取cpu信息
	// 2. 获取负载信息
	// 3. 获取内存信息
	// 4. 获取磁盘信息
	ticker := time.NewTicker(7 * time.Second)
	for range ticker.C {
		o.UpdateStat()
	}
}

func (o *OsStat) UpdateStat() {
	o.mu.Lock()
	defer o.mu.Unlock()
	// 获取cpu信息
	cpuInfo, _ := cpu.Info()
	// 获取负载信息
	allLoad, _ := cpu.Percent(time.Second, false)
	itemLoad, _ := cpu.Percent(time.Second, true)
	coresLoad := make([]dto.CoresLoad, 0)
	for _, load := range itemLoad {
		coresLoad = append(coresLoad, dto.CoresLoad{
			CoreLoad: float64(int(load*10)) / 10,
		})
	}
	memory, _ := mem.VirtualMemory()
	diskOs, _ := disk.Usage("/")
	stat := &dto.Stat{
		Runtime: dto.Runtime{
			NpmVersion:  "10.2.4",
			NodeVersion: "20.11.1",
			GoVersion:   runtime.Version(),
			Os:          runtime.GOOS,
			Arch:        runtime.GOARCH,
		},
		CPU: dto.CPU{
			Manufacturer:  cpuInfo[0].VendorID,
			Brand:         cpuInfo[0].ModelName,
			PhysicalCores: runtime.NumCPU(),
			Speed:         math.Round(cpuInfo[0].Mhz/100) / 10,
			AllLoad:       float64(int(allLoad[0]*10)) / 10,
			CoresLoad:     coresLoad,
		},
		Disk: dto.Disk{
			Size:            int64(diskOs.Total),
			Used:            int64(diskOs.Used),
			Available:       int64(diskOs.Total) - int64(diskOs.Used),
			ConstructorName: "Disk",
		},
		Memory: dto.Memory{
			Total:     int64(memory.Total),
			Available: int64(memory.Total) - int64(memory.Used),
		},
	}
	o.stat = stat
}
