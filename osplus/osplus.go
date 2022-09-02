package osplus

import (
	"io/ioutil"
	"os"

	"github.com/electricface/go-stdlib-compat/io/fs"
)

// call ioutil.TempFile
func CreateTemp(dir, pattern string) (f *os.File, err error) {
	return ioutil.TempFile(dir, pattern)
}

// readDir reads the directory named by dirname and returns
// a sorted list of directory entries.
func ReadDir(dirname string) ([]fs.DirEntry, error) {

	// 标准实现
	// f, err := os.Open(dirname)
	// if err != nil {
	// 	return nil, err
	// }
	// dirs, err := f.Readdir(-1)
	// f.Close()
	// if err != nil {
	// 	return nil, err
	// }
	// sort.Slice(dirs, func(i, j int) bool { return dirs[i].Name() < dirs[j].Name() })

	fileInfos, err := ioutil.ReadDir(dirname)
	// fileInfos 已经排序
	if err != nil {
		return nil, err
	}
	dirs := make([]fs.DirEntry, 0, len(fileInfos))
	for _, fi := range fileInfos {
		dirs = append(dirs, fs.FileInfoToDirEntry(fi))
	}
	return dirs, nil
}
