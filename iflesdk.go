package ifilesdk

import (
	"github.com/ifilespace/ifilesdk-go/config"
	"github.com/ifilespace/ifilesdk-go/file"
	"github.com/ifilespace/ifilesdk-go/model"
	"github.com/ifilespace/ifilesdk-go/user"
)

type Ifile struct {
	config *config.Config
}

func NewIfile(cfg *config.Config) *Ifile {
	return &Ifile{cfg}
}
func (ifile *Ifile) CheckAuth() (res model.CommonRes, err error) {
	res, err = user.NewUser(ifile.config).CheckAuth()
	return
}
func (ifile *Ifile) GetUserInfo(uid int) (uinfo model.GetUserInfoRet, err error) {
	uinfo, err = user.NewUser(ifile.config).GetUserInfo(uid)
	return
}
func (ifile *Ifile) GetUserToken(uid int) (uinfo model.GetUserTokenRet, err error) {
	uinfo, err = user.NewUser(ifile.config).GetUserToken(uid)
	return
}
func (ifile *Ifile) CreateUser(req *model.CreateUserReq) (res model.CreateUserRet, err error) {
	res, err = user.NewUser(ifile.config).CreateUser(req)
	return
}
func (ifile *Ifile) GetUsedCapacity(uid int) (uinfo model.GetUsedCapacityRet, err error) {
	uinfo, err = user.NewUser(ifile.config).GetUsedCapacity(uid)
	return
}
func (ifile *Ifile) ChangePassword(uid int, password string) (uinfo model.CommonRes, err error) {
	uinfo, err = user.NewUser(ifile.config).ChangePassword(uid, password)
	return
}
func (ifile *Ifile) GetFileList(req *model.GetFileListReq) (uinfo model.GetFileListRet, err error) {
	uinfo, err = file.NewFiles(ifile.config).GetFileList(req)
	return
}
func (ifile *Ifile) CreateFolder(req *model.CreateFolderReq) (uinfo model.CreateFolderRet, err error) {
	uinfo, err = file.NewFiles(ifile.config).CreateFolder(req)
	return
}
func (ifile *Ifile) DeleteFolder(req *model.DeleteFolderReq) (uinfo model.DeleteFolderRet, err error) {
	uinfo, err = file.NewFiles(ifile.config).DeleteFolder(req)
	return
}
func (ifile *Ifile) CancelBindFolder(req *model.CancelBindFolderReq) (uinfo model.CancelBindFolderRet, err error) {
	uinfo, err = file.NewFiles(ifile.config).CancelBindFolder(req)
	return
}
