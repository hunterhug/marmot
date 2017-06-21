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
package image

import (
	"errors"
	// this library is get by anther place
	"github.com/hunterhug/go-image/graphics"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

var (
	ExtNotSupportError = errors.New("ext of filename not support")
	FileNameError = errors.New("filename error")
	FileExistError = errors.New("File already exist error")
)

//按宽度和高度进行比例缩放
// scale by width and height from a image file to aother file
func ThumbnailF2F(filename string, savepath string, width int, height int) (err error) {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	src, filetype, err := LoadImage(filename)
	if err != nil {
		return
	}
	err = graphics.Thumbnail(dst, src)
	if err != nil {
		return
	}
	err = SaveImage(savepath, dst, filetype)
	return
}

//按宽度进行比例缩放
// just scale by width from a image file to other file ,maybe ugly
func ScaleF2F(filename string, savepath string, width int) (err error) {
	img, filetype, err := Scale(filename, width)
	if err != nil {
		return
	}
	err = SaveImage(savepath, img, filetype)
	if err != nil {
		return
	}
	return
}

//图像文件的真正名字
// a image file real filename ,such as a tt.jpg may be a tt.png
func RealImageName(filename string) (filerealname string, err error) {
	_, ext, err := LoadImage(filename)
	if err != nil {
		return
	}
	temp := strings.Split(filename, ".")
	if len(temp) < 2 {
		err = FileNameError
	}
	temp[len(temp) - 1] = ext
	filerealname = strings.Join(temp, ".")
	return
}

//文件改名,如果force为假,且新的文件名已经存在,那么抛出错误
// change a file's name,if force is False then if exist file thorw FileExistError
func ChangeImageName(oldname string, newname string, force bool) (err error) {
	if !force {
		_, err = os.Open(newname)
		if err == nil {
			err = FileExistError
			return
		}
	}
	err = os.Rename(oldname, newname)
	return

}


// 根据文件名打开图片,并编码,返回编码对象和文件类型
// Load a image by a filename and return it's type,such as png
func LoadImage(path string) (img image.Image, filetype string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, filetype, err = image.Decode(file)
	return
}

// 将编码对象存入文件中
// save a image object into a file just support png and jpg
func SaveImage(path string, img *image.RGBA, filetype string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return
	}
	if filetype == "png" {
		err = png.Encode(file, img)
	} else if filetype == "jpeg" {
		err = jpeg.Encode(file, img, nil)
	} else {
		err = ExtNotSupportError
	}
	defer file.Close()
	return
}

//对文件中的图片进行等比例变化,宽度为newdx,返回图像编码和文件类型
// see ScaleF2F
func Scale(filename string, newdx int) (dst *image.RGBA, filetype string, err error) {
	src, filetype, err := LoadImage(filename)
	if err != nil {
		return
	}
	bound := src.Bounds()
	dx := bound.Dx()
	dy := bound.Dy()
	dst = image.NewRGBA(image.Rect(0, 0, newdx, newdx * dy / dx))
	// 产生缩略图,等比例缩放
	err = graphics.Scale(dst, src)
	return
}
