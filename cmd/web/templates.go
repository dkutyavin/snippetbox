package main

import "snippetbox.dekutyavin.net/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
