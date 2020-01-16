package utility

// ZipSheetsData takes a slice and a two-dimensional slice and returns
// what in JSON we would think of as an array of objects. The key in each object
// is an item in the keys slice. The value of each object is from
// the data two-dimensional slice. It's assumed that the keys and data are
// in the order in which they're to be zipped.
func ZipSheetsData(keys []interface{}, data [][]interface{}) []map[string]interface{} {
	var jsonData []map[string]interface{}

	for _, d := range data {
		playerData := make(map[string]interface{})

		for i, k := range keys {
			key := k.(string)
			playerData[key] = d[i]
		}

		jsonData = append(jsonData, playerData)
	}

	return jsonData
}
