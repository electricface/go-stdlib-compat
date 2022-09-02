package osplus

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/electricface/go-stdlib-compat/io/fs"
)

func Stat(name string) (fs.FileInfo, error) {
	info, err := os.Stat(name)
	return ToFsFileInfo(info), err
}

func Lstat(name string) (fs.FileInfo, error) {
	info, err := os.Lstat(name)
	return ToFsFileInfo(info), err
}

func Open(name string) (fs.File, error) {
	f, err := os.Open(name)
	return ToFsFile(f), err
}

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
		// TODO
		// dirs = append(dirs, fs.FileInfoToDirEntry(fi))
		_ = fi
	}
	return dirs, nil
}

func (fiw *fsFileInfoWrap) Name() string {
	return fiw.info.Name()
}

func (fiw *fsFileInfoWrap) Size() int64 {
	return fiw.info.Size()
}

func (fiw *fsFileInfoWrap) Mode() fs.FileMode {
	return fs.FileMode(fiw.info.Mode())
}

func (fiw *fsFileInfoWrap) ModTime() time.Time {
	return fiw.info.ModTime()
}

func (fiw *fsFileInfoWrap) IsDir() bool {
	return fiw.info.IsDir()
}

func (fiw *fsFileInfoWrap) Sys() interface{} {
	return fiw.info.Sys()
}

type fsFileInfoWrap struct {
	info os.FileInfo
}

func ToFsFileInfo(fi os.FileInfo) fs.FileInfo {
	if fi == nil {
		return nil
	}
	return &fsFileInfoWrap{fi}
}

type osFileInfoWrap struct {
	info fs.FileInfo
}

// IsDir implements os.FileInfo
func (fiw *osFileInfoWrap) IsDir() bool {
	return fiw.info.IsDir()
}

// ModTime implements os.FileInfo
func (fiw *osFileInfoWrap) ModTime() time.Time {
	return fiw.info.ModTime()
}

// Mode implements os.FileInfo
func (fiw *osFileInfoWrap) Mode() os.FileMode {
	return os.FileMode(fiw.info.Mode())
}

// Name implements os.FileInfo
func (fiw *osFileInfoWrap) Name() string {
	return fiw.info.Name()
}

// Size implements os.FileInfo
func (fiw *osFileInfoWrap) Size() int64 {
	return fiw.info.Size()
}

// Sys implements os.FileInfo
func (fiw *osFileInfoWrap) Sys() interface{} {
	return fiw.info.Sys()
}

func ToOsFileInfo(fi fs.FileInfo) os.FileInfo {
	if fi == nil {
		return nil
	}
	return &osFileInfoWrap{fi}
}

type fileWrap struct {
	f *os.File
}

// Close implements fs.File
func (fw *fileWrap) Close() error {
	return fw.f.Close()
}

// Read implements fs.File
func (fw *fileWrap) Read(b []byte) (int, error) {
	return fw.Read(b)
}

// Stat implements fs.File
func (fw *fileWrap) Stat() (fs.FileInfo, error) {
	fi, err := fw.f.Stat()
	return ToFsFileInfo(fi), err
}

func ToFsFile(f *os.File) fs.File {
	return &fileWrap{f}
}
