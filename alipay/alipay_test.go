package alipay

import (
	"fmt"
	"github.com/buhuang28/go-alipay/logs"
	"github.com/nsf/termbox-go"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/process"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	//"github.com/yusufpapurcu/wmi"
	"github.com/StackExchange/wmi"
	"testing"
	"time"
)

func TestTradePrecreate(t *testing.T) {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	counts, _ := cpu.Counts(true)
	fmt.Println(counts)
	percent, _ := cpu.Percent(time.Second*1, false)
	fmt.Println(percent)
	// convert to JSON. String() is also implemented
	fmt.Println(v)

	var rbs uint64
	var wbs uint64
	for i := 0; i < 10; i++ {
		counters, err := disk.IOCounters()
		if err != nil {
			panic(err)
		}
		for _, v := range counters {
			if v.Name != "C:" {
				continue
			}
			if rbs == 0 {
				rbs = v.ReadBytes
			} else {
				fmt.Println("readBytes:", v.ReadBytes-rbs)
			}
			if wbs == 0 {
				wbs = v.WriteBytes
			} else {
				fmt.Println("writeBytes:", v.WriteBytes-wbs)
			}
			fmt.Println(v.IoTime)
			time.Sleep(time.Second)
		}
	}
}

type Win32_PerfFormattedData_Tcpip_NetworkInterface struct {
	Name                string
	BytesReceivedPerSec uint32
	BytesSentPerSec     uint32
}

func print(x, y int, s string) {
	for _, r := range s {
		termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
		x += 1
	}
}

func TestTradePrecreate2(t *testing.T) {
	//err := termbox.Init()
	//if err != nil {
	//	panic(err)
	//}
	//defer termbox.Close()

	//event_queue := make(chan termbox.Event)
	//go func() {
	//	for {
	//		event_queue <- termbox.PollEvent()
	//	}
	//}()

	var dst []Win32_PerfFormattedData_Tcpip_NetworkInterface
	q := wmi.CreateQuery(&dst, `` /*`WHERE Name = "Realtek PCIe GBE Family Controller"`*/)

	for {
		select {
		case <-time.After(time.Second):
			var d []Win32_PerfFormattedData_Tcpip_NetworkInterface
			err := wmi.Query(q, &d)
			if err != nil {
				panic(err)
			}
			for _, v := range d {
				fmt.Println(fmt.Sprintf("interface name:%s,recv:%d,send:%d", v.Name, v.BytesReceivedPerSec, v.BytesSentPerSec))
			}

		}
	}

}

func TestGetProcess(t *testing.T) {
	processes, err := process.Processes()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, p := range processes {
		name, err := p.Name()
		exe, err := p.Exe()
		fmt.Println(exe)
		if err != nil {
			fmt.Println("Error getting process name:", err)
			continue
		}
		pid := p.Pid

		fmt.Printf("Process ID: %d, Name: %s\n", pid, name)
	}
}

func TestNet(t *testing.T) {
	logs.LogInit()
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range devices {
		if !strings.Contains(v.Description, "AX200") {
			continue
		}
		// 打开网络接口
		handle, err := pcap.OpenLive(v.Name, 65536, false, 3*time.Second)
		if err != nil {
			log.Fatal(err)
		}
		defer handle.Close()

		// 设置过滤器（可选）
		//var filter string = "port 49798"
		//err = handle.SetBPFFilter(filter)
		//if err != nil {
		//	log.Fatal(err)
		//}

		filter := getFilter(9966)
		err = handle.SetBPFFilter(filter)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Only capturing TCP port 49798 packets.")

		// 使用gopacket捕获数据包
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for {
			select {
			case x := <-packetSource.Packets():
				x.Layers()
				log.Info(x)
			}
		}
	}
}

// 定义过滤器
func getFilter(port uint16) string {
	filter := fmt.Sprintf("tcp and ((src port %v) or (dst port %v))", port, port)
	return filter
}

func TestFirewall(t *testing.T) {
	// 需要禁止联网的程序路径
	programPath := "C:\\data\\1.exe"

	// 创建防火墙规则的命令
	cmd := exec.Command("netsh", "advfirewall", "firewall", "add", "rule", "name=\"BlockProgram\"", "dir=out", "action=block", "program=\""+programPath+"\"")
	fmt.Println(cmd.String())
	// 运行命令
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Output:", string(output))
	}
}
