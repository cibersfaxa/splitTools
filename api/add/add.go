package add

import (
	"github.com/cibersfaxa/splitTools/splitTools/api/model"
	"github.com/imroc/req/v3"
)

type Add struct {
	Client *req.Client
}

func (a *Add) AddChapterAPI(chapter model.ChapterDocument) (*model.ChapterDocument, error) {
	r, err := a.Client.R().SetBody(chapter).Post("chapter/add")
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
