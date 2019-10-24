package config

import (
	"V2RayA/models"
	"github.com/mohae/deepcopy"
)

var tr *models.TouchRaw

/*返回的是一份深拷贝*/
func GetTouchRaw() models.TouchRaw {
	if tr == nil {
		tr = new(models.TouchRaw)
		_ = tr.ReadFromFile()
		if tr.Subscriptions == nil {
			tr.Subscriptions = make([]models.SubscriptionRaw, 0)
		}
		if tr.Servers == nil {
			tr.Servers = make([]models.TouchServerRaw, 0)
		}
	}
	return *deepcopy.Copy(tr).(*models.TouchRaw)
}

func SetTouchRaw(newTr *models.TouchRaw) {
	tr = newTr
}