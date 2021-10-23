package dto

type SignUpReq struct {
	PhoneNumber string `json:"phonenumber"`
	Password    string `json:"password"`
	DeviceUUID  string `json:"uuid"`
}
