package splitTools

import (
	"fmt"
	"github.com/imroc/req/v3"
	"os"
)

type Tools struct {
	client *req.Client
}

func NewTools() *Tools {
	client := req.C().SetBaseURL("http://localhost:8080")
	client.EnableDumpAll() // Optional: Enable this to debug requests and responses
	return &Tools{
		client: client,
	}
}

func (t *Tools) UploadFileAPI(filePath string) ([]Chapter, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	resp, err := t.client.R().SetFile("file", filePath).Post("/book/upload")
	if err != nil {
		return nil, fmt.Errorf("failed to post request: %w", err)
	}
	var chapters []Chapter
	err = t.handleResponse(resp, &chapters)
	if err != nil {
		return nil, err
	}
	return chapters, nil
}

func (t *Tools) AddChapterAPI(chapter ChapterDocument) (*ChapterDocument, error) {
	r, err := t.client.R().SetBody(chapter).Post("chapter/add")
	if err != nil {
		return nil, err
	}
	var result ChapterDocument
	err = t.handleResponse(r, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (t *Tools) GetChapterByID(chapterID string) (*ChapterDocument, error) {
	r, err := t.client.R().SetQueryParam("chapter_id", chapterID).Get("chapter/info")
	if err != nil {
		return nil, err
	}
	var result ChapterDocument
	err = t.handleResponse(r, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (t *Tools) GetChaptersByBookID(bookID string) ([]ChapterDocument, error) {
	r, err := t.client.R().SetQueryParam("book_id", bookID).Get("chapter/list")
	if err != nil {
		return nil, err
	}
	var result []ChapterDocument
	err = t.handleResponse(r, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t *Tools) CountChaptersByBookID(bookID string) (int, error) {
	r, err := t.client.R().SetQueryParam("book_id", bookID).Get("chapter/count")
	if err != nil {
		return 0, err
	}
	var result map[string]int
	err = t.handleResponse(r, &result)
	if err != nil {
		return 0, err
	}
	return result["count"], nil
}
func (t *Tools) GetSimilarityBook(keyword string) (map[string]interface{}, error) {
	r, err := t.client.R().SetQueryParam("keyword", keyword).Get("book/search")
	if err != nil {
		return nil, err
	}
	var book map[string]interface{}
	err = t.handleResponse(r, &book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (t *Tools) GetSimilarityByTwoString(str1, str2 string) (map[string]interface{}, error) {
	r, err := t.client.R().SetBody(map[string]string{
		"text1": str1,
		"text2": str2,
	}).Post("similarity")
	if err != nil {
		return nil, err
	}

	var book map[string]interface{}
	err = t.handleResponse(r, &book)
	if err != nil {
		return nil, err
	}
	return book, nil
}
func (t *Tools) ChatCompletionHandler(chat string) (*ChatCompletion, error) {
	r, err := t.client.R().SetFormData(map[string]string{
		"chat": chat,
	}).Post("book/correction")
	if err != nil {
		return nil, err
	}
	var result ChatCompletion
	err = t.handleResponse(r, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil

}

// AddChapterAPI, GetChapterByID, GetChaptersByBookID, CountChaptersByBookID,
// GetSimilarityBook, GetSimilarityByTwoString, and ChatCompletionHandler
// can be optimized similarly as UploadFileAPI.

// Use a common method to handle the response and unmarshalling
func (t *Tools) handleResponse(resp *req.Response, v interface{}) error {
	if resp.IsErrorState() {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	if err := resp.Unmarshal(v); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return nil
}
