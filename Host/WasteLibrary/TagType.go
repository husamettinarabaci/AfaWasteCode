package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//TagType
type TagType struct {
	TagID         float64
	CustomerId    float64
	DeviceId      float64
	UID           string
	Epc           string
	ContainerNo   string
	ContainerType string
	Latitude      float64
	Longitude     float64
	Statu         string
	ImageStatu    string
	Active        string
	ReadTime      string
	CheckTime     string
	CreateTime    string
}

//ToId String
func (res TagType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.TagID)
}

//ToCustomerId String
func (res TagType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToDeviceId String
func (res TagType) ToDeviceIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res TagType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res TagType) ToString() string {
	return string(res.ToByte())

}

//Byte To TagType
func ByteToTagType(retByte []byte) TagType {
	var retVal TagType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To TagType
func StringToTagType(retStr string) TagType {
	return ByteToTagType([]byte(retStr))
}

//SelectSQL
func (res TagType) SelectSQL() string {
	return fmt.Sprintf(`SELECT 
	CustomerId    ,
	DeviceId      ,
	UID           ,
	Epc           ,
	ContainerNo   ,
	ContainerType ,
	Latitude      ,
	Longitude     ,
	Statu         ,
	ImageStatu    ,
	Active        ,
	ReadTime      ,
	CheckTime     ,
	CreateTime    
	FROM public.tags
		   WHERE TagID=%f ;`, res.TagID)

}

//InsertTagDataSQL
func (res TagType) InsertTagDataSQL() string {
	return fmt.Sprintf(`INSERT INTO public.tagdata 
	(TagID,CustomerId,DeviceId,Epc, 
	UID,ContainerNo,ContainerType,Latitude,Longitude,  
	Statu,ImageStatu, 
	ReadTime,CheckTime)  
	   VALUES (%f,%f,%f,'%s''%s''%s','%s',%f,%f,'%s','%s','%s','%s')   
	   RETURNING DataId;`,
		res.TagID, res.CustomerId, res.DeviceId,
		res.Epc, res.UID, res.ContainerNo, res.ContainerType,
		res.Latitude, res.Longitude,
		res.Statu, res.ImageStatu,
		res.ReadTime, res.CheckTime)

}

//InsertSQL
func (res TagType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.tags(
		Epc,DeviceId,CustomerId,UID,Latitude,Longitude)
	   VALUES ('%s',%f,%f,'%s',%f,%f)  
	   RETURNING TagID;`, res.Epc, res.DeviceId,
		res.CustomerId, res.UID,
		res.Latitude, res.Longitude)
}

//UpdateSQL
func (res TagType) UpdateSQL() string {
	if res.Latitude == 0 || res.Longitude == 0 {
		return fmt.Sprintf(`UPDATE public.tags
	SET UID='%s',ReadTime='%s',Statu='%s',DeviceId=%f
   WHERE Epc='%s' AND CustomerId=%f 
   RETURNING TagID;`, res.UID, res.ReadTime, res.Statu,
			res.DeviceId, res.Epc, res.CustomerId)
	} else {
		return fmt.Sprintf(`UPDATE public.tags
	SET UID='%s',ReadTime='%s',Statu='%s',Latitude=%f,Longitude=%f,DeviceId=%f
   WHERE Epc='%s' AND CustomerId=%f 
   RETURNING TagID;`, res.UID, res.ReadTime, res.Statu,
			res.Latitude, res.Longitude,
			res.DeviceId, res.Epc, res.CustomerId)

	}
}
