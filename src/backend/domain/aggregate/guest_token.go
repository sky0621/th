package aggregate

import "github.com/sky0621/familiagildo/domain/entity"

type GuestToken struct {
	Root      *entity.GuestToken
	AuditItem *entity.AuditItem
}
