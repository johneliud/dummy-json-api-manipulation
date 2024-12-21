package utils

import (
	"html/template"
	"strings"
)

// RenderStars dynamically generates star icons based on the rating
func renderRatingStars(rating float64) template.HTML {
	var stars strings.Builder
	fullStars := int(rating)
	halfStar := rating-float64(fullStars) >= 0.5

	// Full stars
	for i := 0; i < fullStars; i++ {
		stars.WriteString(`<ion-icon name="star"></ion-icon>`)
	}

	// Half stars
	if halfStar {
		stars.WriteString(`<ion-icon name="star-half"></ion-icon>`)
	}

	// Empty stars
	for i := fullStars; i < 5; i++ {
		if halfStar && i == fullStars {
			continue
		}
		stars.WriteString(`<ion-icon name="star-outline"></ion-icon>`)
	}
	return template.HTML(stars.String())
}
