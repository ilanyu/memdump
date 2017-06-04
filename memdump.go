package main

import (
	"syscall"
	"fmt"
	"os"
	"strconv"
	"flag"
)

func main() {
	pid := flag.Int("pid", 1, "PID")
	start_addr_hex := flag.String("saddr", "0", "start addr(hex)")
	end_addr_hex := flag.String("eaddr", "0", "end addr(hex)")
	fileName := flag.String("filename", "memdump.mem", "filename")
	flag.Parse()
	start_addr, _ := strconv.ParseInt(*start_addr_hex, 16, 64)
	end_addr, _ := strconv.ParseInt(*end_addr_hex, 16, 64)
	length := end_addr - start_addr
	if err := syscall.PtraceAttach(*pid); err != nil {
		fmt.Errorf("PtraceAttach is err: %s", err)
	}
	var ws syscall.WaitStatus
	syscall.Wait4(*pid, &ws, syscall.WNOHANG, nil)
	path := fmt.Sprintf("/proc/%d/mem", *pid)
	fp, err := os.OpenFile(path, os.O_RDONLY, 0755)
	defer fp.Close()
	if err != nil {
		fmt.Errorf("OpenFile is err: %s", err)
	}
	fp.Seek(start_addr, os.SEEK_SET)
	buf := make([]byte, length)
	n, err := fp.Read(buf)
	if err != nil {
		fmt.Errorf("Read is err: %s", err)
	}
	fmt.Println("read " + strconv.Itoa(n))
	fpw, err := os.OpenFile(*fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	defer fpw.Close()
	fpw.Write(buf)
	if err := syscall.PtraceDetach(*pid); err != nil {
		fmt.Errorf("PtraceDetach is err: %s", err)
	}
}
