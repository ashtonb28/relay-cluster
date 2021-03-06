/*

  Copyright 2017 Loopring Project Ltd (Loopring Foundation).

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

package dao

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
)

func marshalHashListToStr(hashlist []common.Hash) string {
	list := []string{}
	for _, v := range hashlist {
		list = append(list, v.Hex())
	}
	if bs, err := json.Marshal(list); err != nil {
		return ""
	} else {
		return string(bs)
	}
}

func unmarshalStrToHashList(str string) []common.Hash {
	var (
		list    []common.Hash
		strlist []string
	)
	if str == "" {
		return list
	}

	if err := json.Unmarshal([]byte(str), &strlist); err != nil {
		return list
	}
	for _, v := range strlist {
		list = append(list, common.HexToHash(v))
	}
	return list
}
