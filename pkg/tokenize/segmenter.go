// Copyright 2019 AmirQasemi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tokenize

import (
	"errors"
	"github.com/deckarep/golang-set"
	"max_match/internal/app/max_match/utils"
	"unicode/utf8"
)

type Segmenter interface {
	Segment(sentence string) ([]string, error)
}

type MaxMatchSegmenter struct {
	Tokens mapset.Set
}

func(receiver *MaxMatchSegmenter) Segment(sentence string) ([]string, error) {
		if len(sentence) == 0 {
			return make([]string, 0), nil
		}

		revStr := utils.Reverse(sentence)
		for i, w := 0, 0; i < len(sentence); i += w {
			_, width := utf8.DecodeRuneInString(revStr[i:])
			w = width

			firstWord := utils.Reverse(revStr[i:])
			remainder := utils.Reverse(revStr[:i])

			if receiver.Tokens.Contains(firstWord) {
				tmpRes, err := receiver.Segment(remainder)
				if err == nil {
					return append([]string{firstWord},tmpRes...), nil
				}
			}
		}

		/**
		* The above Code is equivalent of following code when sentence is in ASCII encoding
		*/
		/*for i := len(sentence) - 1; i >=0 ; i-- {
			firstWord := sentence[:i + 1]
			var remainder string
			if i == len(sentence) - 1 {
				remainder = ""
			} else {
				remainder = sentence[i + 1:]
			}
			if receiver.Tokens.Contains(firstWord) {
				tmpRes, err := receiver.Segment(remainder)
				if err == nil {
					return append([]string{firstWord},tmpRes...), nil
				}
			}

		}*/

		return nil, errors.New("ERROR")
}