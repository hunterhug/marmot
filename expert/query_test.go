/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:459527502
	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:459527502
	2017.7 by hunterhug
*/
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
