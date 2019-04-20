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

package main

import "max_match/internal/app/max_match"

// Input and Output Parameters
var (
	asset string = "assets/"

	testTokens   string = asset + "testTokens"
	farsiTokens   string = asset + "fa.words.txt"
	englishTokens string = asset + "en.tokens.en"

	testMergedTokens   string = asset + "testMergedTokens"
	farsiMergedTokens   string = asset + "mergedTokens.fa"
	englishMergedTokens string = asset + "mergedTokens.en"

	outputPrefix string = "94463147_Assignment1_"

	testOutput   string = asset + outputPrefix + "TEST"
	farsiOutput   string = asset + outputPrefix + "FA"
	englishOutput string = asset + outputPrefix + "EN"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	//Running on Test test sample
	testTestSample := max_match.TestSample{TokensPath: testTokens, TestSetPath: testMergedTokens, OutputPath: testOutput}
	testTestSample.RunTestSample(false)

	// Running on Farsi test sample
	farsiTestSample := max_match.TestSample{TokensPath: farsiTokens, TestSetPath: farsiMergedTokens, OutputPath: farsiOutput}
	farsiTestSample.RunTestSample(true)

	// Running on English test sample
	englishTestSample := max_match.TestSample{TokensPath: englishTokens, TestSetPath: englishMergedTokens, OutputPath: englishOutput}
	englishTestSample.RunTestSample(false)
}
