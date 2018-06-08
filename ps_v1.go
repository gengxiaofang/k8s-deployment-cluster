/*
参考链接：
https://github.com/bosun-monitor/bosun/blob/c3017f169dfb568c52c17abed9e1606096f91b73/cmd/scollector/collectors/processes_linux.go
https://github.com/prometheus/prometheus/blob/737ae60ceaebc4264358e28e377eda061861117e/vendor/github.com/prometheus/procfs/proc_stat.go
*/
package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"strconv"
)
//获取PID路径
func PidNum() []string {
	var b= true
	var PidData []string

	Dir, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}

	reg := regexp.MustCompile(`^[\d]{1,}`)
	for _, dir := range Dir {
		b = reg.MatchString(dir.Name())
		if dir.IsDir() && true == b {
			PidData = append(PidData, dir.Name())
		}
	}

	return PidData
}

func PidStat(Pid []string)  {
	var totalVirtualMem int64
	var totalRSSMem int64
	var osPageSize = os.Getpagesize()
	var TotalScollectorMemoryMB uint64
	for _,pid := range Pid{
		file_status,err := os.Stat("/proc/"+ pid)
		if err != nil {
			log.Fatal(err)
			continue
		}
		start_ts := file_status.ModTime().Unix()

		stats_file,err := ioutil.ReadFile("/proc/" + pid + "/stat")
		if err !=nil {
			log.Fatal(err)
			continue
		}
		stats := strings.Fields(string(stats_file))
		if len(stats) < 24 {
			err = fmt.Errorf("stats too short")
			continue
		}

		user,err := strconv.ParseInt(stats[13],10,64)
		if err != nil {
			log.Fatal(err)
			continue
		}
		sys,err := strconv.ParseInt(stats[14],10,64)
		if err != nil {
			log.Fatal(err)
			continue
		}

		virtual,err := strconv.ParseInt(stats[22],10,64)
		if err != nil {
			log.Fatal(err)
		}
		totalVirtualMem += virtual

		rss,err := strconv.ParseInt(stats[23],10,64)
		if err != nil {
			log.Fatal(err)
		}
		if pid == string(os.Getegid()) {
			TotalScollectorMemoryMB = uint64(rss) * uint64(osPageSize) / 1024 / 1024
		}

		totalRSSMem += rss
		//		fmt.Printf("Pid:%v ProcStartTime:%v ProcUptime:%v User:%v Sys:%v totalVirtualMem:%v TotalScollectorMemoryMB:%v",pid,start_ts,new()-start_ts,user,sys,totalVirtualMem,TotalScollectorMemoryMB)
		fmt.Printf("Pid:%v StartTime:%v us:%v sys:%v totalVirtualMem:%v TotalScollectorMemoryMB:%v",pid,start_ts,user,sys,totalVirtualMem,TotalScollectorMemoryMB)
		}

}

func main()  {
	PidStat(PidNum())
}
