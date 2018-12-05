package model

type Miniocon struct {
	ID int
	//Endpoint 对象存储服务的URL
	Endpoint string
	//AccessKeyID Access key是唯一标识你的账户的用户ID。
	AccessKeyID string
	//SecretAccessKey 	Secret key是你账户的密码。
	SecretAccessKey string
	// Secure true代表使用HTTPS
	Secure bool
}
