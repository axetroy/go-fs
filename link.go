package fs

import "os"

func Link(existingPath string, newPath string) (error) {
  return os.Link(existingPath, newPath)
}

func ReadLink(path string) (string, error) {
  return os.Readlink(path)
}

func Symlink(target string, path string) (error) {
  return os.Symlink(target, path)
}
