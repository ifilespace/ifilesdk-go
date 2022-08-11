package user

import (
	"encoding/json"
	"time"

	"github.com/ifilespace/ifilesdk-go/config"
	"github.com/ifilespace/ifilesdk-go/model"
	"github.com/ifilespace/ifilesdk-go/util"
)

type User struct {
	config *config.Config
}

func NewUser(cfg *config.Config) *User {
	return &User{cfg}
}
func (user *User) CheckAuth() (ret model.CommonRes, err error) {
	now := time.Now().Unix()
	sign := util.GetSignToken(user.config.Keyid, user.config.Keysecret, now)
	req := &model.CommonReq{Keyid: user.config.Keyid, Date: now}
	res, err := util.PostJSON(user.config.Url+"/bapi/user/checkauth?sign="+sign, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &ret)
	return
}
func (user *User) GetUserInfo(uid int) (ret model.GetUserInfoRet, err error) {
	now := time.Now().Unix()
	sign := util.GetSignToken(user.config.Keyid, user.config.Keysecret, now)
	req := &model.GetUserInfoReq{Uid: uid}
	req.Keyid = user.config.Keyid
	req.Date = now
	res, err := util.PostJSON(user.config.Url+"/bapi/user/getuserinfo?sign="+sign, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &ret)
	return
}
func (user *User) GetProjectToken(uid int, fileid string) (ret model.GetUserTokenRet, err error) {
	now := time.Now().Unix()
	sign := util.GetSignToken(user.config.Keyid, user.config.Keysecret, now)
	req := &model.GetProjectTokenReq{Uid: uid, Fileid: fileid}
	req.Keyid = user.config.Keyid
	req.Date = now
	res, err := util.PostJSON(user.config.Url+"/bapi/user/getprojecttoken?sign="+sign, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &ret)
	return
}
func (user *User) GetUserToken(uid int) (ret model.GetUserTokenRet, err error) {
	now := time.Now().Unix()
	sign := util.GetSignToken(user.config.Keyid, user.config.Keysecret, now)
	req := &model.GetUserInfoReq{Uid: uid}
	req.Keyid = user.config.Keyid
	req.Date = now
	res, err := util.PostJSON(user.config.Url+"/bapi/user/getusertoken?sign="+sign, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &ret)
	return
}
func (user *User) CreateUser(req *model.CreateUserReq) (ret model.CreateUserRet, err error) {
	now := time.Now().Unix()
	sign := util.GetSignToken(user.config.Keyid, user.config.Keysecret, now)
	req.Keyid = user.config.Keyid
	req.Date = now
	res, err := util.PostJSON(user.config.Url+"/bapi/user/createuser?sign="+sign, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &ret)
	return
}

func (user *User) GetUsedCapacity(uid int) (ret model.GetUsedCapacityRet, err error) {
	now := time.Now().Unix()
	sign := util.GetSignToken(user.config.Keyid, user.config.Keysecret, now)
	req := &model.GetUsedCapacityReq{Uid: uid}
	req.Keyid = user.config.Keyid
	req.Date = now
	res, err := util.PostJSON(user.config.Url+"/bapi/user/getusedcapacity?sign="+sign, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &ret)
	return
}

func (user *User) ChangePassword(uid int, password string) (ret model.CommonRes, err error) {
	now := time.Now().Unix()
	sign := util.GetSignToken(user.config.Keyid, user.config.Keysecret, now)
	req := &model.ChangePasswordReq{Uid: uid, Password: password}
	req.Keyid = user.config.Keyid
	req.Date = now
	res, err := util.PostJSON(user.config.Url+"/bapi/user/changepassword?sign="+sign, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &ret)
	return
}
func (user *User) JoinProject(uid int, fileid string) (ret model.CommonRes, err error) {
	now := time.Now().Unix()
	sign := util.GetSignToken(user.config.Keyid, user.config.Keysecret, now)
	req := &model.JLProjectReq{Uid: uid, Fileid: fileid}
	req.Keyid = user.config.Keyid
	req.Date = now
	res, err := util.PostJSON(user.config.Url+"/bapi/user/joinproject?sign="+sign, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &ret)
	return
}

func (user *User) LeaveProject(uid int, fileid string) (ret model.CommonRes, err error) {
	now := time.Now().Unix()
	sign := util.GetSignToken(user.config.Keyid, user.config.Keysecret, now)
	req := &model.JLProjectReq{Uid: uid, Fileid: fileid}
	req.Keyid = user.config.Keyid
	req.Date = now
	res, err := util.PostJSON(user.config.Url+"/bapi/user/leaveproject?sign="+sign, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &ret)
	return
}
