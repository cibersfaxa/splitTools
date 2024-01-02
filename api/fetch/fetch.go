package fetch

import (
	"fmt"
	"github.com/cibersfaxa/splitTools/splitTools/api/model"
	"github.com/imroc/req/v3"
	"os"
)

type Fetch struct {
	Client *req.Client
}

type CallbackDocument struct {
	Err              error
	ChapterDocuments []model.ChapterDocument
}

func (f *Fetch) ChapterByID(chapterID string) (*model.ChapterDocument, error) {
	r, err := f.Client.R().SetQueryParam("chapter_id", chapterID).Get("chapter/info")
	if err != nil {
		return nil, err
	}
	var result model.ChapterDocument
	err = r.UnmarshalJson(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *CallbackDocument) ForEachDocument(callback func(chapter model.ChapterDocument)) error {
	for _, chapter := range c.ChapterDocuments {
		callback(chapter)
	}
	return nil
}

func (f *Fetch) ChaptersByBookID(bookID string) *CallbackDocument {
	r, err := f.Client.R().SetQueryParam("book_id", bookID).Get("chapter/list")
	if err != nil {
		return &CallbackDocument{Err: err}
	}
	var result []model.ChapterDocument
	err = r.UnmarshalJson(&result)
	if err != nil {
		return &CallbackDocument{Err: err}
	}
	return &CallbackDocument{ChapterDocuments: result, Err: nil}
}

func (f *Fetch) SimilarityBook(keyword string) (map[string]interface{}, error) {
	r, err := f.Client.R().SetQueryParam("keyword", keyword).Get("book/search")
	if err != nil {
		return nil, err
	}
	var book map[string]interface{}
	err = r.UnmarshalJson(&book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (f *Fetch) ChatCompletionHandler(chat string) (*model.ChatCompletion, error) {
	r, err := f.Client.R().SetFormData(map[string]string{
		"chat": chat,
	}).Post("book/correction")
	if err != nil {
		return nil, err
	}
	var result model.ChatCompletion
	err = r.UnmarshalJson(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil

}

func (f *Fetch) UploadFileAPI(filePath string) ([]model.Chapter, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	resp, err := f.Client.R().SetFile("file", filePath).Post("/book/upload")
	if err != nil {
		return nil, fmt.Errorf("failed to post request: %w", err)
	}
	var chapters []model.Chapter
	err = resp.UnmarshalJson(&chapters)
	if err != nil {
		return nil, err
	}
	return chapters, nil
}

func (f *Fetch) CountChaptersByBookID(bookID string) (int, error) {
	r, err := f.Client.R().SetQueryParam("book_id", bookID).Get("chapter/count")
	if err != nil {
		return 0, err
	}
	var result map[string]int
	err = r.UnmarshalJson(&result)
	if err != nil {
		return 0, err
	}
	return result["count"], nil
}

func (f *Fetch) SimilarityByTwoString(str1, str2 string) (map[string]interface{}, error) {
	r, err := f.Client.R().SetBody(map[string]string{
		"text1": str1,
		"text2": str2,
	}).Post("similarity")
	if err != nil {
		return nil, err
	}

	var book map[string]interface{}
	err = r.UnmarshalJson(&book)
	if err != nil {
		return nil, err
	}
	return book, nil
}
