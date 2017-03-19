package model

import "html/template"

type Image struct {

	Type string
	Name string
	Directory string
	Url template.URL

}

type ImageData struct {
	Token string
	Images []Image
}
