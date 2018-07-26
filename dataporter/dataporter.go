package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/y-tac/gosom/som"
)

// Request格納用
type traitAPIRequest struct {
	Unit som.Unit
}

func main() {
	fmt.Println("Start::DataClients")

	for {
		param := traitAPIRequest{}
		param.Unit.Red = memUsage()
		param.Unit.Blue = cpuData()
		param.Unit.Green = diskusage()
		fmt.Println(param.Unit)
		input, err := json.Marshal(param)
		resp, err := http.Post("http://localhost:3306/trait", "application/json", bytes.NewBuffer(input))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.println(resp.Body)
		time.Sleep(5 * time.Second)

	}

}
func cpuData() int {
	c, _ := cpu.Percent(5*time.Second, false)
	return int(c[0] / 100 * som.MaxValue)
}

func memUsage() int {
	m, _ := mem.VirtualMemory()
	return int(m.UsedPercent / 100 * som.MaxValue)
}
func diskusage() int {
	parts, _ := disk.Partitions(false)
	somTotal := float64(0)
	somUsed := float64(0)
	for _, part := range parts {
		u, _ := disk.Usage(part.Mountpoint)
		somTotal += float64(u.Total)
		somUsed += float64(u.Total)
	}
	return int(float64(som.MaxValue) * somUsed / somTotal)

}
