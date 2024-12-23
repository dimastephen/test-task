package model

type refreshInfo struct {
	ip   string
	guid string
}

type RefreshUpdate struct {
	NewRefreshToken string
	Ip              string
	Guid            string
}
