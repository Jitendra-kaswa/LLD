package src

import "strconv"

type AccessLevel int

const (
	Read AccessLevel = iota
	Write
	Owner
)

type Permission struct {
	id          int
	userId      int
	entityId    int
	accessLevel AccessLevel
}

func NewPermission(userId int, id int, entityId int, accessLevel AccessLevel) *Permission {
	return &Permission{
		userId:      userId,
		id:          id,
		entityId:    entityId,
		accessLevel: accessLevel,
	}
}

func (p Permission) GetId() int {
	return p.id
}
func (p Permission) GetUserId() int {
	return p.userId
}
func (p Permission) GetEntityId() int {
	return p.entityId
}
func (p *Permission) ChangeAccessLevel(accessLevel AccessLevel) error {
	p.accessLevel = accessLevel
	return nil
}
func (p *Permission) GetAccessLevel() AccessLevel {
	return p.accessLevel
}

type IPermissionRepository interface {
	CreatePermission(userId int, entityId int, accessLevel AccessLevel) *Permission
	CheckPermission(userId int, entityId int, accessLevel AccessLevel) bool
	ChangePermission(userId int, entityId int, accessLevel AccessLevel) bool
}

type PermissionRepositoryMap struct {
	permissionMap map[string]*Permission
}

func NewPermissionRepository() IPermissionRepository {
	return &PermissionRepositoryMap{
		permissionMap: make(map[string]*Permission),
	}
}

func (prm *PermissionRepositoryMap) CreatePermission(userId int, entityId int, accessLevel AccessLevel) *Permission {
	permission_key := strconv.Itoa(userId) + "_" + strconv.Itoa(entityId)
	if _, exists := prm.permissionMap[permission_key]; exists {
		return nil
	}
	newPermission := NewPermission(userId, 0, entityId, accessLevel)
	prm.permissionMap[permission_key] = newPermission
	return newPermission
}
func (prm PermissionRepositoryMap) CheckPermission(userId int, entityId int, accessLevel AccessLevel) bool {
	permission_key := strconv.Itoa(userId) + "_" + strconv.Itoa(entityId)
	if _, exists := prm.permissionMap[permission_key]; !exists {
		return false
	}
	return true // here we can add check if requested accessLevel is read/view and actual permission is Owner then it's true
}
func (prm PermissionRepositoryMap) ChangePermission(userId int, entityId int, accessLevel AccessLevel) bool {
	permission_key := strconv.Itoa(userId) + "_" + strconv.Itoa(entityId)
	if permission, exists := prm.permissionMap[permission_key]; !exists {
		permission.ChangeAccessLevel(accessLevel)
		return true
	}
	return false
}
