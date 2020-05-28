package expert

import (
	"fmt"
	"testing"
)

func TestFindPicture(t *testing.T) {
	data := `
		
		https://imgsa.baidu.com/forum/w%3D580/sign=294db374d462853592e0d229a0ee76f2/e732c895d143ad4b630e8f4683025aafa40f0611.jpg
		
		https://imgsa.baidu.com/forum/whttps:/4d462853592e0d229a0ee76f2/e732c895d143ad4b630e8f4683025aafa40f0611.jpg
		https://img1.jpg
		http://im62853592e0d229a0ee76f2/e732c895d143ad4b630e8f4683025aafa40f0611.jpgsfsadfsda
		httpdb374://aafa40f0611.jpg

		src="http://s.jpg"

		src="https://s.jpg"

		src = "http://s.jpg"
		src="https:s.jpg"
		"https://img1.jpg" "https://img1.jpgsss","https://img1.jpgss","https://img1.jpgss"
		`
	result := FindPicture(data)
	fmt.Printf("%#v", result)
}
