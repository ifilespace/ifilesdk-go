package file

import (
	"encoding/json"
	"time"

	"github.com/ifilespace/ifilesdk-go/config"
	"github.com/ifilespace/ifilesdk-go/model"
	"github.com/ifilespace/ifilesdk-go/util"
)

type Files struct {
	config *config.Config
}

func NewFiles(cfg *config.Config) *Files {
	return &Files{cfg}
}

func (file *Files) GetFileList(req *model.GetFileListReq) (ret model.GetFileListRet, err error) {
	now := time.Now().Unix()
	sign := util.GetSignToken(file.config.Keyid, file.config.Keysecret, now)
	req.Keyid = file.config.Keyid
	req.Date = now
	res, err := util.PostJSON(file.config.Url+"/bapi/file/getfilelist?sign="+sign, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &ret)
	return
}
func (file *Files) CreateFolder(req *model.CreateFolderReq) (ret model.CreateFolderRet, err error) {
	now := time.Now().Unix()
	sign := util.GetSignToken(file.config.Keyid, file.config.Keysecret, now)
	req.Keyid = file.config.Keyid
	req.Date = now
	res, err := util.PostJSON(file.config.Url+"/bapi/file/createfolder?sign="+sign, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &ret)
	return
}
func (file *Files) DeleteFolder(req *model.DeleteFolderReq) (ret model.DeleteFolderRet, err error) {
	now := time.Now().Unix()
	sign := util.GetSignToken(file.config.Keyid, file.config.Keysecret, now)
	req.Keyid = file.config.Keyid
	req.Date = now
	res, err := util.PostJSON(file.config.Url+"/bapi/file/deletefolder?sign="+sign, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &ret)
	return
}
func (file *Files) CancelBindFolder(req *model.CancelBindFolderReq) (ret model.CancelBindFolderRet, err error) {
	now := time.Now().Unix()
	sign := util.GetSignToken(file.config.Keyid, file.config.Keysecret, now)
	req.Keyid = file.config.Keyid
	req.Date = now
	res, err := util.PostJSON(file.config.Url+"/bapi/file/cancelbindfolder?sign="+sign, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &ret)
	return
}
