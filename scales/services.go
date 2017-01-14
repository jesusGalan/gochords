package goscales

import "encoding/json"

/*GetAll return all scale bytes*/
func GetAll(repo []byte) map[string]map[string]map[string][]string {
	scaleBytes := MapperJSONByte(repo)

	return scaleBytes
}

/*MapperJSONByte can read json bytes containing arrays of int
inside.*/
func MapperJSONByte(bit []byte) map[string]map[string]map[string][]string {
	m := map[string]map[string]map[string][]string{}
	err := json.Unmarshal(bit, &m)
	if err != nil {
		panic(err)
	}
	return m
}
