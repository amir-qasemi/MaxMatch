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

package file

import (
	"bufio"
	"fmt"
	"github.com/deckarep/golang-set"
	"os"
	"regexp"
	"strings"
)

func GetTokensOfFile(filePath string) (mapset.Set, error) {
	result := mapset.NewSet()

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result.Add(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func GetSentences(filePath string) ([]string, error) {
	result := make([]string, 0)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func WriteResultToFile(filePath string, result map[string][]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for k, v := range result {
		fmt.Fprintln(file, k+"\t"+strings.Join(v, " "))
	}

	return nil
}

func GetCleanedTokens(filePath string) (mapset.Set, error) {
	var re = regexp.MustCompile(`\d+$`)

	result, err := GetTokensOfFile(filePath)
	if err != nil {
		return nil, err
	}

	cleanedResult := mapset.NewSet()
	for element := range result.Iter() {
		//fmt.Println(re.ReplaceAllString(element.(string), ``))
		cleanedResult.Add(re.ReplaceAllString(element.(string), ``))
	}

	return cleanedResult, nil
}
