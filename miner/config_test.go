// 
// 	Copyright 2017 by marmot author: gdccmcm14@live.com.
// 	Licensed under the Apache License, Version 2.0 (the "License");
// 	you may not use this file except in compliance with the License.
// 	You may obtain a copy of the License at
// 		http://www.apache.org/licenses/LICENSE-2.0
// 	Unless required by applicable law or agreed to in writing, software
// 	distributed under the License is distributed on an "AS IS" BASIS,
// 	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// 	See the License for the specific language governing permissions and
// 	limitations under the License
//

package miner

import (
	"context"
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	worker, _ := New(nil)
	fmt.Printf("%#v\n", worker)

	if worker.BeforeAction == nil {
		fmt.Println("good")
	}

	worker.BeforeAction = func(c context.Context, worker *Worker) {
		worker.SetHeaderParm("Marmot", "v2")
	}
}
