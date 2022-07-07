package setting

import (
	"net/http"

	"github.com/fikryfahrezy/adea/los-postgre/file"
	"github.com/fikryfahrezy/adea/los-postgre/resp"
)

type SettingApp struct {
	file file.File
}

func NewSetting(file file.File) *SettingApp {
	return &SettingApp{
		file: file,
	}
}

func (a *SettingApp) ZipTmp(w http.ResponseWriter, r *http.Request) {
	contentDisposition := "attachment; filename=tmp.zip"
	w.Header().Set("Content-Disposition", contentDisposition)
	w.WriteHeader(http.StatusOK)
	if err := a.file.ZipSource("./tmp", w); err != nil {
		resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
		return
	}
}

func (a *SettingApp) LoadZipTmp(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(1024); err != nil {
		resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
		return
	}

	file, header, err := r.FormFile("zip")
	if err != nil {
		resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
		return
	}
	defer file.Close()

	if err := a.file.Unzip("./tmp", file, header.Size); err != nil {
		resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
		return
	}

	resp.NewResponse(http.StatusOK, "file inserted", nil).HttpJSON(w, resp.NewHttpBody(nil))
}
