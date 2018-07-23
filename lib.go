package sh

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"
	"syscall"
)

// ================================================================================
func Cd(dir string) {
	err := os.Chdir(dir)
	ck(err)
}

func Pwd() string {
	dir, err := os.Getwd()
	ck(err)
	return dir
}

func ReadLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// ================================================================================
func absoluteFile(file string) string {
	fullFile := file
	if !strings.HasPrefix(file, "/") {
		fullFile = path.Join(Pwd(), file)
	}
	return fullFile
}

func fileMode(file string) os.FileMode {
	stat, err := os.Stat(file)
	ck(err)
	return stat.Mode()
}

func IsDir(file string) bool {
	fullFile := absoluteFile(file)
	mode := fileMode(fullFile)
	return mode.IsDir()
}

func IsFile(file string) bool {
	fullFile := absoluteFile(file)
	mode := fileMode(fullFile)
	return mode.IsRegular()
}

func FilePermForCaller(file string) (readable, writable, executable bool) {
	stat, err := os.Stat(file)
	ck(err)
	linuxStat := stat.Sys().(*syscall.Stat_t)

	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	var r, w, x uint
	if fmt.Sprintf("%d", linuxStat.Uid) == u.Uid {
		r, w, x = 8, 7, 6
	} else if fmt.Sprintf("%d", linuxStat.Gid) == u.Gid {
		r, w, x = 5, 4, 3
	} else {
		r, w, x = 2, 1, 0
	}

	perm := stat.Mode().Perm()
	return perm&(1<<r) > 0, perm&(1<<w) > 0, perm&(1<<x) > 0
}

// ================================================================================
func Run(script string) {
}

// ================================================================================
func ck(err error) {
	if err != nil {
		panic(err)
	}
}
