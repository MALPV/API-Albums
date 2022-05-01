package data

import "malpvapps/web-services-gin/models"

// albums slice to seed record album data.
var Albums = []models.Album{
	{ID: "1", Title: "Reir o llorar", Artist: "Jacqueline Valenzuela", Price: 10.000},
	{ID: "2", Title: "Arcoiris", Artist: "Marcos Palma", Price: 17.99},
	{ID: "3", Title: "Sueños", Artist: "El niño", Price: 39.99},
}
