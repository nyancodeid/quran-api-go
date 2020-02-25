package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/novalagung/gubrak/v2"
)

func main() {
	surahResult := fetch("http://quran.kemenag.go.id/index.php/api/v1/surat/0/114")
	// Parse `surahResult` to Surah struct
	surah, _ := UnmarshalSurah(surahResult)
	var surahList []map[string]interface{}

	for i := range surah.Data {
		detail := surah.Data[i]

		ayatResult := fetch(
			fmt.Sprintf("https://quran.kemenag.go.id/index.php/api/v1/ayatweb/%d/0/0/%d",
				detail.ID,
				detail.CountAyat,
			),
		)
		// Parse `ayatResult` to Ayat struct
		ayat, err := UnmarshalAyat(ayatResult)
		if err != nil {
			log.Println(string(ayatResult))
			log.Fatalln(err)
		}

		log.Println(
			fmt.Sprintf("[HTTP][OK] Surah=%s IdSurah=%d Length=%d", detail.SuratName, detail.ID, len(ayat.Data)),
		)

		surahDetail := jsonStoreAyat(detail, ayat)
		// Append `surahDetail` to `surahList`
		surahList = append(surahList, surahDetail)

		// Log spacer
		fmt.Println("")
	}
	// Store surah list
	jsonStoreSurah(surahList)
}

// Fetch Function make response available offline by save
// it as json file with url hash as a filename
func fetch(url string) []byte {
	hash := makeHash(url)
	cacheName := fmt.Sprintf("caches/%s.json", hash)

	// Check availablity cache file
	if _, err := os.Stat(cacheName); err == nil {
		log.Println("[CACHE][HIT] Hash=" + hash)

		data, _ := ioutil.ReadFile(cacheName)
		// Return immediately after got cache file
		return data
	}

	log.Println("[CACHE][MISS] Hash=" + hash)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	// Close body after function done
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	} else {
		// Write cache file
		ioutil.WriteFile(cacheName, body, os.ModePerm)
	}

	return body
}

// Customize key values to then be saved in the json file
func jsonStoreAyat(surah SurahDatum, ayats Ayat) map[string]interface{} {
	// Map function on gubrak is helpful
	dataAyats := gubrak.From(ayats.Data).Map(func(each AyatDatum) map[string]interface{} {
		return map[string]interface{}{
			"id":               each.AyaID,
			"index":            each.AyaNumber,
			"text":             each.AyaText,
			"text_translation": each.TranslationAyaText,
			"audio":            fmt.Sprintf("https://quran.kemenag.go.id/cmsq/source/s01/%03d%03d.mp3", each.SuraID, each.AyaNumber),
		}
	}).Result()

	dataSurah := map[string]interface{}{
		"id":               ayats.Data[0].SuraID,
		"id_juz":           ayats.Data[0].JuzID,
		"name":             surah.SuratName,
		"name_arabic":      surah.SuratText,
		"name_translation": surah.SuratTerjemahan,
		"total":            surah.CountAyat,
		"ayat":             dataAyats,
	}

	// Encode to JSON with indent style 2 spaces
	dataJson, _ := json.MarshalIndent(dataSurah, "", "  ")
	pathJson := fmt.Sprintf("dist/surah/%d.json", surah.ID)
	// Write json file
	err := ioutil.WriteFile(pathJson, dataJson, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("[EXPORT][OK] Surah=" + surah.SuratName)
	// Remove `ayat` key from `surah`
	delete(dataSurah, "ayat")

	return dataSurah
}

// Store list of Surah as a JSON file
func jsonStoreSurah(surah []map[string]interface{}) {
	dataJson, _ := json.MarshalIndent(surah, "", "  ")
	pathJson := fmt.Sprintf("dist/surah.json")

	err := ioutil.WriteFile(pathJson, dataJson, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
}

// Generate MD5 Hash from string
func makeHash(rawData string) string {
	data := []byte(rawData)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
