package libcommon

import (
	"time"
	"container/list"
)

// include all ORM struct

// table files
type FileDO struct {
	Id         int64  `gorm:"column:id;auto_increment;primary_key" json:"id"`
	Md5        string `gorm:"column:md5" json:"md5"`
	PartNumber int    `gorm:"column:parts_num" json:"parts_num"`
	Group      string `gorm:"column:grop" json:"group"`
	Instance   string `gorm:"column:instance" json:"instance"`
	Finish     int    `gorm:"column:finish" json:"finish"`
}

func (FileDO) TableName() string {
	return "file"
}

// table clients
type StorageClientDO struct {
	Uuid       string `gorm:"primary_key" json:"uuid"`
	Host       string `gorm:"column:host" json:"host""`
	Port       int    `gorm:"column:port" json:"port"`
	Status     int    `gorm:"column:status" json:"status"`
	TrackerId  int64  `gorm:"column:tracker" json:"tracker"`
	TotalFiles int    `gorm:"column:total_files" json:"total_files"`
	Group      string `gorm:"column:grop" json:"group"`
	InstanceId string `gorm:"column:instance_id" json:"instance_id"`
	HttpPort   int    `gorm:"column:http_port" json:"http_port"`
	HttpEnable bool   `gorm:"column:http_enable" json:"http_enable"`
	StartTime  int64  `gorm:"column:start_time" json:"start_time"`
	Download   int64  `gorm:"column:downloads" json:"downloads"`
	Upload     int64  `gorm:"column:uploads" json:"uploads"`
	Disk       int64  `gorm:"column:disk" json:"disk"`
	ReadOnly   bool   `gorm:"column:read_only" json:"read_only"`
	Finish     int    `gorm:"column:finish" json:"finish"`
	IOin       int64  `gorm:"column:ioin" json:"ioin"`
	IOout      int64  `gorm:"column:ioout" json:"ioout"`
}

func (StorageClientDO) TableName() string {
	return "storage_client"
}

// table part
type PartDO struct {
	Id   int64  `gorm:"column:id;auto_increment;primary_key" json:"id"`
	Md5  string `gorm:"column:md5" json:"md5"`
	Size int64  `gorm:"column:size" json:"size"`
}

func (PartDO) TableName() string {
	return "part"
}

// table relation_file_part
type FilePartRelationDO struct {
	Id     int64 `gorm:"column:id;auto_increment;primary_key" json:"id"`
	FileId int64 `gorm:"column:fid" json:"fid"`
	PartId int64 `gorm:"column:pid" json:"pid"`
}

func (FilePartRelationDO) TableName() string {
	return "relation_file_part"
}

// table sys
type SysDO struct {
	Key   string `gorm:"column:key;primary_key" json:"key"`
	Value string `gorm:"column:value" json:"value"`
}

func (SysDO) TableName() string {
	return "sys"
}

// table tracker
type TrackerDO struct {
	Uuid          string    `gorm:"column:uuid;primary_key" json:"uuid"`
	TrackerSyncId int64     `gorm:"column:tracker_sync_id" json:"tracker_sync_id"`
	LastRegTime   time.Time `gorm:"column:last_reg_time" json:"last_reg_time"`
	LocalPushId   int64     `gorm:"column:local_push_id" json:"local_push_id"`
}

func (TrackerDO) TableName() string {
	return "tracker"
}

// table web_storage_log
type WebStorageLogsDO struct {
	Id        int64 `gorm:"column:id;auto_increment;primary_key" json:"id"`
	StorageId int64 `gorm:"column:storage" json:"storage"`
	LogTime   int64 `gorm:"column:log_time" json:"log_time"`
	IOin      int64 `gorm:"column:ioin" json:"ioin"`
	IOout     int64 `gorm:"column:ioout" json:"ioout"`
	Disk      int64 `gorm:"column:disk" json:"disk"`
	Memory    int64 `gorm:"column:mem" json:"mem"`
	Download  int64 `gorm:"column:download" json:"download"`
	Upload    int64 `gorm:"column:upload" json:"upload"`
}

func (WebStorageLogsDO) TableName() string {
	return "web_storage_log"
}

// table web_storage
type WebStorageDO struct {
	Id         int64  `gorm:"column:id;auto_increment;primary_key" json:"id"`
	Host       string `gorm:"column:host" json:"host""`
	Port       int    `gorm:"column:port" json:"port"`
	Status     int    `gorm:"column:status" json:"status"`
	TrackerId  int64  `gorm:"column:tracker" json:"tracker"`
	Uuid       string `gorm:"column:uuid" json:"uuid"`
	TotalFiles int    `gorm:"column:total_files" json:"total_files"`
	Group      string `gorm:"column:grop" json:"group"`
	InstanceId string `gorm:"column:instance_id" json:"instance_id"`
	HttpPort   int    `gorm:"column:http_port" json:"http_port"`
	HttpEnable bool   `gorm:"column:http_enable" json:"http_enable"`
	IOin       int64  `gorm:"column:ioin" json:"ioin"`
	IOout      int64  `gorm:"column:ioout" json:"ioout"`
	Disk       int64  `gorm:"column:disk" json:"disk"`
	StartTime  int64  `gorm:"column:start_time" json:"start_time"`
	Download   int64  `gorm:"column:downloads" json:"downloads"`
	Upload     int64  `gorm:"column:uploads" json:"uploads"`
	ReadOnly   bool   `gorm:"column:read_only" json:"read_only"`
	Finish     int    `gorm:"column:finish" json:"finish"`
}

func (WebStorageDO) TableName() string {
	return "web_storage"
}

// table web_tracker
type WebTrackerDO struct {
	Id      int64     `gorm:"column:id" json:"id"`
	Uuid    string    `gorm:"column:uuid" json:"uuid"`
	Host    string    `gorm:"column:host" json:"host"`
	Port    int       `gorm:"column:port" json:"port"`
	Status  int       `gorm:"column:status" json:"status"` // 0: disabled,  1:enabled, 3: deleted
	Secret  string    `gorm:"column:secret" json:"secret"`
	TotalFiles int    `gorm:"column:files" json:"files"`
	Remark  string    `gorm:"column:remark" json:"remark"`
	AddTime time.Time `gorm:"column:add_time" json:"add_time"`
}
func (WebTrackerDO) TableName() string {
	return "web_tracker"
}

// table files
type FileVO struct {
	Id         int64    `gorm:"column:id;auto_increment;primary_key" json:"id"`
	Md5        string   `gorm:"column:md5" json:"md5"`
	PartNumber int      `gorm:"column:parts_num" json:"parts_num"`
	Group      string   `gorm:"column:grop" json:"group"`
	Instance   string   `gorm:"column:instance" json:"instance"`
	Finish     int      `gorm:"column:finish" json:"finish"`
	Parts      []PartDO `gorm:"-" json:"parts"`
}

func (FileVO) TableName() string {
	return "file"
}

func (vo *FileVO) From(fileDO *FileDO) {
	if fileDO == nil {
		return
	}
	vo.Id = fileDO.Id
	vo.Md5 = fileDO.Md5
	vo.PartNumber = fileDO.PartNumber
	vo.Group = fileDO.Group
	vo.Instance = fileDO.Instance
	vo.Finish = fileDO.Finish
}

// set file parts of fileVO,
// list member must be *PartDO
func (vo *FileVO) SetParts(parts *list.List) {
	if parts == nil {
		return
	}
	temp := make([]PartDO, parts.Len())
	index := 0
	for ele := parts.Front(); ele != nil; ele = ele.Next() {
		temp[index] = *ele.Value.(*PartDO)
		index++
	}
	vo.Parts = temp
}

// set file parts of fileVO,
// list member must be *PartDO
func (vo *FileVO) SetPartsFromVO(parts *list.List) {
	if parts == nil {
		return
	}
	temp := make([]PartDO, parts.Len())
	index := 0
	for ele := parts.Front(); ele != nil; ele = ele.Next() {
		pdo := &PartDO{}
		item := *ele.Value.(*PartVO)
		pdo.Id = item.Id
		pdo.Md5 = item.Md5
		pdo.Size = item.Size
		temp[index] = *pdo
		index++
	}
	vo.Parts = temp
}

// table part
type PartVO struct {
	Id     int64  `gorm:"column:id;auto_increment;primary_key" json:"id"`
	FileId int64  `gorm:"column:fid" json:"fid"`
	Md5    string `gorm:"column:md5" json:"md5"`
	Size   int64  `gorm:"column:size" json:"size"`
}

func (PartVO) TableName() string {
	return "part"
}

// result count struct
type Total struct {
	Count int `gorm:"column:count" json:"count"`
}

// result count struct
type Statistic struct {
	FileCount   int   `gorm:"column:files" json:"files"`
	FinishCount int   `gorm:"column:finish" json:"finish"`
	DiskSpace   int64 `gorm:"column:disk" json:"disk"`
}









