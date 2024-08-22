package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func run() {
	fmt.Printf("Running %v as %d \n", os.Args[2:], os.Getpid())

	cmd := exec.Command("proc/self/exe", append([]string {"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr {
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	cmd.Run()
}

func child() {
       //  fmt.Printf("Running %v \n", os.Args[2:], os.Getpid())
       // syscall.Sethostname([]byte("container"))
  
        must(syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, ""))
	      must(os.MkdirAll("rootfs/oldrootfs", 0700))
	      must(syscall.PivotRoot("rootfs", "rootfs/oldrootfs"))
	      must(os.Chdir("/"))

        cmd := exec.Command(os.Args[2], os.Args[3:]...)
        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        if err := cmd.Run(); err != nil {
		    fmt.Println("ERROR", err)
		    os.Exit(1)
	      }
  
        cmd.Run()
}


func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("What should I do?")
	}
}

func must(err error)	{
	if err != nil {
		panic(err)
	}
}

