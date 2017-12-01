package fs

import (
  "os"
)

/**
remove a file or a dir
 */
func Remove(name string) (err error) {
  return os.RemoveAll(name)
}
