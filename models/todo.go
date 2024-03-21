package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ! Json içindeki verinin boş olmasına izin vermemek için "omitempty" yazılır.
// ! "primitive.ObjectID" mongo dbye ekleme yapacağımız için
type Todo struct {
	Id      primitive.ObjectID `json:"id,omitempty"` 
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}