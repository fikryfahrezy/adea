package setting

import (
	"net/http"

	"github.com/fikryfahrezy/adea/los-inmen/data"
	"github.com/fikryfahrezy/adea/los-inmen/file"
	"github.com/fikryfahrezy/adea/los-inmen/resp"
)

type SettingApp struct {
	file file.File
	db   *data.JsonFile
}

func NewSetting(file file.File, db *data.JsonFile) *SettingApp {
	return &SettingApp{
		file: file,
		db:   db,
	}
}

func (a *SettingApp) LoadJsonDB(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(1024); err != nil {
		resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
		return
	}

	dbname := r.FormValue("dbname")
	file, _, err := r.FormFile("json")
	if err != nil {
		resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
		return
	}
	defer file.Close()

	if err := a.db.ScanToMap(file, dbname); err != nil {
		resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
		return
	}

	resp.NewResponse(http.StatusOK, "file inserted", nil).HttpJSON(w, resp.NewHttpBody(nil))
}

func (a *SettingApp) GanerateJsonDB(w http.ResponseWriter, r *http.Request) {
	contentDisposition := "attachment; filename=los-db.json"
	w.Header().Set("Content-Disposition", contentDisposition)
	w.WriteHeader(http.StatusOK)
	if err := a.db.Generate(w); err != nil {
		resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
		return
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
