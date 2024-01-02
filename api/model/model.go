package model

type Chapter struct {
	Index         int    `json:"index"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	ChapterLength int    `json:"chapter_length"`
}

type ChapterDocument struct {
	ID                 string `bson:"_id" json:"_id" binding:"required"`
	ChapterContent     string `bson:"chapter_content" json:"chapter_content" binding:"required"`
	OriginContent      string `bson:"origin_content" json:"origin_content" binding:"required"`
	ComparisonResult   int    `bson:"comparison_result" json:"comparison_result" binding:"required"`
	LocalContentLength int    `bson:"local_content_length" json:"local_content_length"`
	BookID             string `bson:"book_id" json:"book_id" binding:"required"`
	ChapterID          string `bson:"chapter_id" json:"chapter_id"`
	ChapterTitle       string `bson:"chapter_title" json:"chapter_title"`
	IsPaid             string `bson:"is_paid" json:"is_paid"`
}

type ChatCompletion struct {
	NewBookName string `json:"new_book_name"`
	OldBookName string `json:"old_book_name"`
}
