package fs

import "os"

/**
stat a file
 */
func Stat(name string) (info os.FileInfo, err error) {
  return os.Stat(name)
}

/**
stat a file
 */
func LStat(name string) (info os.FileInfo, err error) {
  return os.Lstat(name)
}
