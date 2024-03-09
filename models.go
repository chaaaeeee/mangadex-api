package main

type MangaResponse struct {
	Response string  `json:"response"`
	Data     []Manga `json:"data"`
}

type Manga struct {
	ID            string          `json:"id"`
	Type          string          `json:"type"`
	Attributes    MangaAttributes `json:"attributes"`
	Relationships []Relationships `json:"relationships"`
}

type MangaAttributes struct {
	Title                  map[string]string         `json:"title"`
	AltTitles              []MangaAltTitles          `json:"altTitles"`
	Description            MangaAttributeDescription `json:"description"`
	PublicationDemographic string                    `json:"publicationDemographic"`
	Tags                   []MangaAttributeTags      `json:"tags"`
	Status                 string                    `json:"status"`
	Year                   int                       `json:"year"`
}

type Relationships struct {
	ID         string                   `json:"id"`
	Type       string                   `json:"type"`
	Attributes []RelationshipAttributes `json:"attributes"`
}

type RelationshipAttributes struct {
	FileName  string `json:"fileName"`
	CreatedAt string `json:"createdAt"`
}

type MangaAltTitles struct {
	En   string `json:"en"`
	JaRo string `json:"ja-ro"`
	Ja   string `json:"ja"`
}

type MangaAttributeDescription struct {
	En string `json:"en"`
}

type MangaAttributeTags struct {
	ID         string             `json:"id"`
	Type       string             `json:"type"`
	Attributes MangaTagAttributes `json:"attributes"`
}

type MangaTagAttributes struct {
	Name TagAttributeName `json:"name"`
}

type TagAttributeName struct {
	En string `json:"en"`
}

type Search struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
