/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2020
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/
package util

import (
	"encoding/json"
)

func JsonBack(s []byte) ([]byte, error) {
	temp := new(interface{})
	json.Unmarshal(s, temp)
	result, err := json.Marshal(temp)
	return result, err
}
