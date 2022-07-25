package model

type CommonReq struct {
	Keyid string `json:"keyid"`
	Date  int64  `json:"date"`
}

type CommonRes struct {
	Msg      string `json:"msg"`
	Status   int    `json:"status"`
	Leixinng int    `json:"leixing,omitempty"`
}
type GetUserInfoReq struct {
	CommonReq
	Uid int `json:"uid"`
}

type GetUserInfoRet struct {
	CommonRes
	Data Userinfo `json:"data"`
}

type GetUserTokenRet struct {
	CommonRes
	Data    string `json:"data"`
	Expire  int64  `json:"expire"`
	Driveid string `json:"driveid"`
}

type CreateUserReq struct {
	CommonReq
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
}
type CreateUserRet struct {
	CommonRes
	Data Userinfo `json:"data"`
}

type GetUsedCapacityReq struct {
	CommonReq
	Uid int `json:"uid"`
}

type GetUsedCapacityRet struct {
	CommonRes
	Data int `json:"data"`
}
type ChangePasswordReq struct {
	CommonReq
	Uid      int    `json:"uid"`
	Password string `json:"password"`
}
type Userinfo struct {
	ID            int    `json:"id"`
	Username      string `json:"username"`
	Nickname      string `json:"nickname"`
	Email         string `json:"email"`
	Mobile        string `json:"mobile"`
	Regtime       int    `json:"regtime"`
	Lastlogintime int    `json:"lastlogintime"`
}

type GetFileListReq struct {
	CommonReq
	UID     int    `json:"uid"`
	PFileID string `json:"p_file_id"`
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
}

type GetFileListRet struct {
	CommonRes
	Data []FileList `json:"data"`
}

type FileList struct {
	FileID          string `json:"file_id"`
	PFileID         string `json:"p_file_id"`
	Name            string `json:"name"`
	Addtime         string `json:"addtime"`
	Updatetime      string `json:"updatetime"`
	Status          string `json:"status"`
	PreHash         string `json:"PreHash"`
	Size            int    `json:"size"`
	Type            string `json:"type"`
	Path            string `json:"path"`
	Special         int    `json:"special"`
	Public          int    `json:"public"`
	Share           int    `json:"share"`
	Favorite        int    `json:"favorite"`
	Cover           string `json:"cover"`
	FileExtension   string `json:"file_extension,omitempty"`
	Category        string `json:"category,omitempty"`
	ContentHash     string `json:"content_hash,omitempty"`
	ContentHashName string `json:"content_hash_name,omitempty"`
	ContentType     string `json:"content_type,omitempty"`
}

type CreateFolderReq struct {
	CommonReq
	Uid     int    `json:"uid"`
	PFileID string `json:"p_file_id"`
	Name    string `json:"name"`
	Special int    `json:"special"`
}
type CreateFolderRet struct {
	CommonRes
	Data string `json:"data"`
}
type DeleteFolderReq struct {
	CommonReq
	Uid    int    `json:"uid"`
	FileID string `json:"fileid"`
}
type DeleteFolderRet struct {
	CommonRes
	Data string `json:"data"`
}
type CancelBindFolderReq struct {
	CommonReq
	Uid    int    `json:"uid"`
	FileID string `json:"fileid"`
}
type CancelBindFolderRet struct {
	CommonRes
	Data string `json:"data"`
}
