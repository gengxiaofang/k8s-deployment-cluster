package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"fmt"
	"os"
)

func main() {

	var b = true
	var pid []string //存放pid号

	path := "/proc"
	Dir,err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	// 取出pid号
	reg := regexp.MustCompile(`^[\d]{1,}`)
	for _,dir := range Dir {
		//通过regexp正则表达式过滤出pid只是数字的目录,并且返回true;
		b = reg.MatchString(dir.Name())
		// 通过if取出/proc下的所有目录，通过布尔类型(b)判断是否是数字;
		if dir.IsDir() && true == b{
			// 追加Pid号到 slice pid 中;
			pid = append(pid,dir.Name())
		}
	}

	// 实现打印功能
	for _,Pid := range pid {
		cmd_file := path + "/" + Pid + "/cmdline"

		cmd,err := os.Open(cmd_file)
		if err != nil {
			log.Fatal(err)
		}
		defer cmd.Close()

		//通过ioutil读取文件;
		cmd1,err := ioutil.ReadAll(cmd)
		fmt.Printf("pid:% -4v \t cmdline:%v\n",Pid,string(cmd1))
	}

}