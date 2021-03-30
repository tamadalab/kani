package utils

// copy from https://qiita.com/Nabetani/items/6f1785290270409486d6

import (
	"runtime"
	"syscall"
	"time"
)

// emulate touch command.
func Touch(path string) error {
	now := time.Now().UnixNano()
	const nano = 1000 * 1000 * 1000
	t := syscall.Timespec{Sec: now / nano, Nsec: now % nano}
	ut := []syscall.Timespec{t, t}
	err := syscall.UtimesNano(path, ut)
	if err == nil {
		return nil
	}
	isNoEnt := func(e error) bool {
		if e == syscall.ENOENT {
			return true
		}
		switch runtime.GOOS {
		case "windows":
			// Windows では syscall.ERROR_PATH_NOT_FOUND を返すことがある。
			// しかし、非 Windows 環境では syscall.ERROR_PATH_NOT_FOUND は定義されない。
			// ifdef も使えない。別ファイルにするのはめんどくさい。
			// 仕方ないので自分で定義する。
			const ERROR_PATH_NOT_FOUND = 3 // 「don't use ALL_CAPS in Go names」と言われるが、 syscall に合わせる。
			if errno, ok := e.(syscall.Errno); ok && errno == ERROR_PATH_NOT_FOUND {
				return true
			}
			return false
		case "darwin":
			return false
		case "linux":
			// たぶん OK だけど、テストしてない。
			return false
		default:
			panic("you should write something here")
		}
	}
	if !isNoEnt(err) {
		return err
	}
	fd, err := syscall.Open(path, syscall.O_CREAT|syscall.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer syscall.Close(fd)
	return nil
}
