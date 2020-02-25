package main

import (
	"encoding/json"
)

func UnmarshalSurah(data []byte) (Surah, error) {
	var r Surah
	err := json.Unmarshal(data, &r)
	return r, err
}
func (r *Surah) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Surah struct {
	Msg  string       `json:"msg"`
	Data []SurahDatum `json:"data"`
}
type SurahDatum struct {
	ID              int64  `json:"id"`
	SuratName       string `json:"surat_name"`
	SuratText       string `json:"surat_text"`
	SuratTerjemahan string `json:"surat_terjemahan"`
	CountAyat       int64  `json:"count_ayat"`
}

func UnmarshalAyat(data []byte) (Ayat, error) {
	var r Ayat
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Ayat) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Ayat struct {
	Data []AyatDatum `json:"data"`
}

type AyatDatum struct {
	AyaID              int64  `json:"aya_id"`
	AyaNumber          int64  `json:"aya_number"`
	AyaText            string `json:"aya_text"`
	SuraID             int64  `json:"sura_id"`
	JuzID              int64  `json:"juz_id"`
	PageNumber         int64  `json:"page_number"`
	TranslationAyaText string `json:"translation_aya_text"`
}
