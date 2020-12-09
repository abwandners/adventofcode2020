package solutions

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readInputFile(fileName string) []map[string]string {
	rawInput, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	rawDataSets := strings.Split(string(rawInput), "\n\n")

	dataSets := make([]map[string]string,0)

	for _, rawDataSet := range rawDataSets {
		dataSets = append(dataSets, parseDataset(rawDataSet))
	}

	return dataSets
}

func parseDataset(dataset string) map[string]string {
	data := make(map[string]string)
	dataset = strings.Replace(dataset, "\n", " ", -1)
	rawKVs := strings.Split(dataset, " ")
	for _, rawKV := range rawKVs {
		kv := strings.Split(rawKV, ":")
		if len(kv) != 2 {
			panic(fmt.Sprintf("incomplete dataset: %#v", kv))
		}
		data[kv[0]] = kv[1]
	}
	return data
}
