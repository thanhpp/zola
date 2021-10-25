package dto

type SignUpReq struct {
	PhoneNumber string `form:"phonenumber"`
	Password    string `form:"password"`
	DeviceUUID  string `form:"uuid"`
}
