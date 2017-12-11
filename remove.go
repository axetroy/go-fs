package fs

import (
  "os"
)

/**
remove a file or a dir
 */
func Remove(name string) (error) {
  return os.RemoveAll(name)
}
