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

package max_match

import (
	"github.com/deckarep/golang-set"
	"log"
	"max_match/internal/app/max_match/utils/file"
	"max_match/pkg/tokenize"
)

type TestSample struct {
	TokensPath  string
	TestSetPath string
	OutputPath  string
}


func (receiver * TestSample) RunTestSample(cleanTokens bool) {
	resultMap := make(map[string][]string)

	var tokens mapset.Set
	var err error
	if cleanTokens {
		tokens, err = file.GetCleanedTokens(receiver.TokensPath)
	} else {
		tokens, err = file.GetTokensOfFile(receiver.TokensPath)
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	segmenter := &tokenize.MaxMatchSegmenter{Tokens: tokens}
	sentences, err := file.GetSentences(receiver.TestSetPath)
	if err != nil {
		log.Fatal("%v", err)
		return
	}

	for _, sentence  := range sentences {
		result, err := segmenter.Segment(sentence)
		if err == nil {
			resultMap[sentence] = result
		} else {
			resultMap[sentence] = []string{err.Error()}
		}
	}

	err = file.WriteResultToFile(receiver.OutputPath, resultMap)
	if err != nil {
		log.Fatal("%v", err)
	}
}