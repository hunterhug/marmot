/*
Copyright 2017 by GoSpider author.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package util

// 功能： 文件帮助功能
import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// 获取调用者的当前文件DIR
//Get the caller now directory
func CurDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}

// 获取当前执行二进制所在的位置
func GetBinaryCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)

	// go run 时会生成在临时文件中
	//fmt.Println(path)
	if err != nil {
		return "", err
	}

	if strings.Contains(path, "command-line-arguments") {
		return GetCurrentPath()
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

// 获取当前执行命令所在的位置
func GetCurrentPath() (string, error) {
	return os.Getwd()
}

//将字节数组保存到文件中去
//Save bytes into file
func SaveToFile(filepath string, content []byte) error {
	//全部权限写入文件
	err := ioutil.WriteFile(filepath, content, 0777)
	return err
}

// read bytes from file
func ReadfromFile(filepath string) ([]byte, error) {
	return ioutil.ReadFile(filepath)
}

// Get file info
func GetFilenameInfo(filepath string) (os.FileInfo, error) {
	fileinfo, err := os.Stat(filepath)
	return fileinfo, err
}

// rename
func Rename(oldfilename string, newfilename string) error {
	return os.Rename(oldfilename, newfilename)
}

//根据传入文件夹名字递归新建文件夹
//Create dir by recursion
func MakeDir(filedir string) error {
	return os.MkdirAll(filedir, 0777)
}

//根据传入文件名，递归创建文件夹
// ./dir/filename  /home/dir/filename
//Create dir by the filename
func MakeDirByFile(filepath string) error {
	temp := strings.Split(filepath, "/")
	if len(temp) <= 2 {
		return errors.New("please input complete file name like ./dir/filename or /home/dir/filename")
	}
	dirpath := strings.Join(temp[0:len(temp)-1], "/")
	return MakeDir(dirpath)
}

func FileExist(filename string) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		return false
	}
	if fi.IsDir() {
		return false
	} else {
		return true
	}
}

// 递归列出文件夹下文件（全路径）
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {

		if fi.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

// 列出文件夹下非递归文件全称
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	suffix = strings.ToUpper(suffix)
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+"/"+fi.Name())
		}
	}
	return files, nil
}

// 列出文件夹下文件名字
func ListDirOnlyName(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	suffix = strings.ToUpper(suffix)
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, fi.Name())
		}
	}
	return files, nil
}

//判断文件或文件夹是否存在
func HasFile(s string) bool {
	f, err := os.Open(s)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	f.Close()
	return true
}

//File-File复制文件
func CopyFF(src io.Reader, dst io.Writer) error {
	_, err := io.Copy(dst, src)
	return err
}

//File-String复制文件
func CopyFS(src io.Reader, dst string) error {
	f, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, src)
	return err
}

//判断是否是文件
func IsFile(filepath string) bool {
	fielinfo, err := os.Stat(filepath)
	if err != nil {
		return false
	} else {
		if fielinfo.IsDir() {
			return false
		} else {
			return true
		}
	}
}

//判断是否是文件夹
func IsDir(filepath string) bool {
	fielinfo, err := os.Stat(filepath)
	if err != nil {
		return false
	} else {
		if fielinfo.IsDir() {
			return true
		} else {
			return false
		}
	}
}

//文件状态
func FileStatus(filepath string) {
	fielinfo, err := os.Stat(filepath)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%v", fielinfo)
	}
}

//文件夹下数量
func SizeofDir(dirPth string) int {
	if IsDir(dirPth) {
		files, _ := ioutil.ReadDir(dirPth)
		return len(files)
	}

	return 0
}

//获取文件后缀
func GetFileSuffix(f string) string {
	fa := strings.Split(f, ".")
	if len(fa) == 0 {
		return ""
	} else {
		return fa[len(fa)-1]
	}
}

/*
# 去除标题中的非法字符 (Windows)
*/
func ValidFileName(filename string) string {
	patterns := []string{
		" ", "#01#",
		"\\", "#02#",
		"/", "#03#",
		":", "#04#",
		"\"", "#05#",
		"?", "#06#",
		"<", "#07#",
		">", "#08#",
		"|", "#09#",
	}
	r := strings.NewReplacer(patterns...)
	dudu := r.Replace(filename)
	return dudu
}

func ValidBackFileName(filename string) string {
	patterns := []string{
		"#01#", " ",
		"#02#", "\\",
		"#03#", "/",
		"#04#", ":",
		"#05#", "\"",
		"#06#", "?",
		"#07#", "<",
		"#08#", ">",
		"#09#", "|",
	}
	r := strings.NewReplacer(patterns...)
	dudu := r.Replace(filename)
	return dudu
}
