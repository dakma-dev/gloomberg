package openrarity

type ItemRank struct {
	Rank  int     `json:"rank"`
	Score float64 `json:"score"`
}

//
//func LoadRaritiesFromJSONs() map[string]map[int]ItemRank {
//	rarities := make(map[string]map[int]ItemRank, 0)
//
//	rarityJSONsPath := "internal/assets/rarities/"
//
//	files, err := os.ReadDir(rarityJSONsPath)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, file := range files {
//		slug := strings.Replace(file.Name(), ".json", "", 1)
//
//		content, err := fs.ReadFile(os.DirFS(rarityJSONsPath), file.Name())
//		if err != nil {
//			log.Fatal("error when opening file: ", err)
//		}
//
//		var payload map[int]ItemRank
//
//		err = json.Unmarshal(content, &payload)
//		if err != nil {
//			log.Fatal("Error during Unmarshal(): ", err)
//		}
//
//		rarities[slug] = payload
//	}
//
//	return rarities
//}
