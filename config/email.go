package config

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUser     string
	SMTPPassword string
}

func GetConfig() *EmailConfig {
	return &EmailConfig{
		SMTPHost:     "smtp.gmail.com",
		SMTPPort:     587,
		SMTPUser:     "sam.malakjan@gmail.com",
		SMTPPassword: "pvmlbkkiysajmgjo",
	}
}
