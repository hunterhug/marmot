/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2021
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/
package util

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func CurDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}

func GetBinaryCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)

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
		return "", errors.New(`error: Can't find "/" or "\"`)
	}
	return string(path[0 : i+1]), nil
}

func GetCurrentPath() (string, error) {
	return os.Getwd()
}

func SaveToFile(filePath string, content []byte) error {
	err := ioutil.WriteFile(filePath, content, 0777)
	return err
}

func ReadFromFile(filepath string) ([]byte, error) {
	return ioutil.ReadFile(filepath)
}

func GetFilenameInfo(filepath string) (os.FileInfo, error) {
	info, err := os.Stat(filepath)
	return info, err
}

func Rename(oldFilename string, newFilename string) error {
	return os.Rename(oldFilename, newFilename)
}

func MakeDir(fileDir string) error {
	return os.MkdirAll(fileDir, 0777)
}

func MakeDirByFile(filepath string) error {
	temp := strings.Split(filepath, "/")
	if len(temp) <= 2 {
		return errors.New("please input complete file name like ./dir/filename or /home/dir/filename")
	}
	dirPath := strings.Join(temp[0:len(temp)-1], "/")
	return MakeDir(dirPath)
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
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, dirPth+"/"+fi.Name())
		}
	}
	return files, nil
}

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
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, fi.Name())
		}
	}
	return files, nil
}

func HasFile(s string) bool {
	f, err := os.Open(s)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	f.Close()
	return true
}

func CopyFF(src io.Reader, dst io.Writer) error {
	_, err := io.Copy(dst, src)
	return err
}

func CopyFS(src io.Reader, dst string) error {
	f, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, src)
	return err
}

func IsFile(filepath string) bool {
	info, err := os.Stat(filepath)
	if err != nil {
		return false
	} else {
		if info.IsDir() {
			return false
		} else {
			return true
		}
	}
}

func IsDir(filepath string) bool {
	info, err := os.Stat(filepath)
	if err != nil {
		return false
	} else {
		if info.IsDir() {
			return true
		} else {
			return false
		}
	}
}

func SizeofDir(dirPth string) int {
	if IsDir(dirPth) {
		files, _ := ioutil.ReadDir(dirPth)
		return len(files)
	}

	return 0
}

func GetFileSuffix(f string) string {
	fa := strings.Split(f, ".")
	if len(fa) == 0 {
		return ""
	} else {
		return fa[len(fa)-1]
	}
}

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
