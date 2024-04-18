package entity

type Config struct {
	DbUrl         string
	UploadFolder  string
	JWTSignString string
	MailFrom      string
	MailPassword  string
	MailAddress   string
	MailPort      int64
}
