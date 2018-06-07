/*
参考链接：
https://github.com/bosun-monitor/bosun/blob/c3017f169dfb568c52c17abed9e1606096f91b73/cmd/scollector/collectors/processes_linux.go
https://github.com/prometheus/prometheus/blob/737ae60ceaebc4264358e28e377eda061861117e/vendor/github.com/prometheus/procfs/proc_stat.go
*/
package main

import (
	"os"
	"fmt"
	"log"
	"io/ioutil"
	"regexp"
	"bytes"
)

const userHZ = 100

type ProcStat struct {
	PID int
	Comm string
	State string
	PPID	int
	PGRP	int
	Session int
	TTY 	int
	TPGID	int
	Flags 	int
	MinFlt	uint
	CMinFlt	uint
	MajFlt	uint
	CMajFlt uint
	STime 	uint
	CUTime	uint
	CSTime	uint
	Prioyity int
	Nice	int
	NumThreads int
	StartTime	uint64
	Vsize 	int
	RSS		int
}

/*
type Proc struct {
	PID int
}
type Procs []Proc
*/

//获取PID路径
func PidNum() string {
	var b= true
	var Pid string
	var piddata []string

	Dir, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}

	reg := regexp.MustCompile(`^[\d]{1,}`)
	for _, dir := range Dir {
		b = reg.MatchString(dir.Name())
		if dir.IsDir() && true == b {
			piddata = append(piddata, dir.Name())
		}
	}
	for _, pid := range piddata {
		Pid = pid
	}
	return Pid
}

func PidStat(Pid string) (ProcStat,error) {
	f,err := os.Open("/proc/"+ Pid + "/"+"stat")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data,err := ioutil.ReadAll(f)
	if err != nil {
		return ProcStat{},err
	}

	var r = bytes.LastIndex(data,[]byte(")"))
	var s = ProcStat{}
	_,err = fmt.Fscan(
		bytes.NewBuffer(data[r+2:]),
		&s.State,
		&s.PPID,
		&s.PGRP,
		&s.Session,
		&s.TTY,
		&s.TPGID,
		&s.Flags,
		&s.MinFlt,
		&s.MajFlt,
		&s.CMajFlt,
		&s.STime,
		&s.CUTime,
		&s.CSTime,
		&s.Nice,
		&s.NumThreads,
		&s.Vsize,
		&s.RSS,
	)
	if err != nil {
		return ProcStat{},err
	}
	fmt.Println(s)
	return s,nil
}

func (s ProcStat) VirtualMemory() int {
	return s.Vsize
}

func (s ProcStat) ResidentMemory() int {
	return s.RSS * os.Getpagesize()
}

func main() {
		fmt.Println(PidNum())
	PidStat(PidNum())
}
