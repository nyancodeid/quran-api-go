## JSON Schema
```json
{
  "$schema": "http://json-schema.org/draft-06/schema#",
  "$ref": "#/definitions/Surah",
  "definitions": {
    "Surah": {
      "type": "object",
      "additionalProperties": false,
      "properties": {
        "ayat": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Ayat"
          }
        },
        "id": {
          "type": "integer"
        },
        "id_juz": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "name_arabic": {
          "type": "string"
        },
        "name_translation": {
          "type": "string"
        },
        "total": {
          "type": "integer"
        }
      },
      "required": [
        "ayat",
        "id",
        "id_juz",
        "name",
        "name_arabic",
        "name_translation",
        "total"
      ],
      "title": "Surah"
    },
    "Ayat": {
      "type": "object",
      "additionalProperties": false,
      "properties": {
        "audio": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "index": {
          "type": "integer"
        },
        "text": {
          "type": "string"
        },
        "text_translation": {
          "type": "string"
        }
      },
      "required": [
        "audio",
        "id",
        "index",
        "text",
        "text_translation"
      ],
      "title": "Ayat"
    }
  }
}
```

## Go
```go
package main

import "encoding/json"

func UnmarshalSurah(data []byte) (Surah, error) {
  var r Surah
  err := json.Unmarshal(data, &r)
  return r, err
}

func (r *Surah) Marshal() ([]byte, error) {
  return json.Marshal(r)
}

type Surah struct {
  Ayat            []Ayat `json:"ayat"`            
  ID              int64  `json:"id"`              
  IDJuz           int64  `json:"id_juz"`          
  Name            string `json:"name"`            
  NameArabic      string `json:"name_arabic"`     
  NameTranslation string `json:"name_translation"`
  Total           int64  `json:"total"`           
}

type Ayat struct {           
  Audio            string `json:"audio"`
  ID              int64  `json:"id"`              
  Index           int64  `json:"index"`           
  Text            string `json:"text"`            
  TextTranslation string `json:"text_translation"`
}
```

## Dart
```dart
class Surah {
    List<Ayat> ayat;
    int id;
    int idJuz;
    String name;
    String nameArabic;
    String nameTranslation;
    int total;

    Surah({
        this.ayat,
        this.id,
        this.idJuz,
        this.name,
        this.nameArabic,
        this.nameTranslation,
        this.total,
    });
}

class Ayat {
    String audio;
    int id;
    int index;
    String text;
    String textTranslation;

    Ayat({
        this.audio,
        this.id,
        this.index,
        this.text,
        this.textTranslation,
    });
}
```

## Python

```py
from typing import List


class Ayat:
    audio: str
    id: int
    index: int
    text: str
    text_translation: str

    def __init__(self, audio: str, id: int, index: int, text: str, text_translation: str) -> None:
        self.audio = audio
        self.id = id
        self.index = index
        self.text = text
        self.text_translation = text_translation


class Surah:
    ayat: List[Ayat]
    id: int
    id_juz: int
    name: str
    name_arabic: str
    name_translation: str
    total: int

    def __init__(self, ayat: List[Ayat], id: int, id_juz: int, name: str, name_arabic: str, name_translation: str, total: int) -> None:
        self.ayat = ayat
        self.id = id
        self.id_juz = id_juz
        self.name = name
        self.name_arabic = name_arabic
        self.name_translation = name_translation
        self.total = total
```

## Typescript
```ts
export interface Surah {
    ayat:             Ayat[];
    id:               number;
    id_juz:           number;
    name:             string;
    name_arabic:      string;
    name_translation: string;
    total:            number;
}

export interface Ayat {
    audio:            string;
    id:               number;
    index:            number;
    text:             string;
    text_translation: string;
}
```

## Ruby
```rb
require 'json'
require 'dry-types'
require 'dry-struct'

module Types
  include Dry::Types.module

  Int    = Strict::Int
  Hash   = Strict::Hash
  String = Strict::String
end

class Ayat < Dry::Struct
  attribute :audio,            Types::String
  attribute :id,               Types::Int
  attribute :index,            Types::Int
  attribute :text,             Types::String
  attribute :text_translation, Types::String
end

class Surah < Dry::Struct
  attribute :ayat,             Types.Array(Ayat)
  attribute :id,               Types::Int
  attribute :id_juz,           Types::Int
  attribute :surah_name,       Types::String
  attribute :name_arabic,      Types::String
  attribute :name_translation, Types::String
  attribute :total,            Types::Int
end
```
