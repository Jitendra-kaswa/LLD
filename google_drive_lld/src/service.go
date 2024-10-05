package src

type IService interface {
	CreateFolder(name string, parentId int) *Folder
	CreateFile(name string, parentId int, content []byte) *File

	AllChildsOfFolder(id int) []string
	MoveFolderToNewDest(id int, newParentId int) bool
}

type Service struct {
	userRepository       IUserRepository
	entityRepository     IEntityRepository
	permissionRepository IPermissionRepository
}

func NewService() IService {
	return &Service{
		userRepository:       NewUserRepositoryMap(),
		entityRepository:     NewEntityRepository(),
		permissionRepository: NewPermissionRepository(),
	}
}

func (s *Service) CreateFolder(name string, parentId int) *Folder {
	parentFolder := s.entityRepository.GetEntityById(parentId)
	if parentFolder != nil {
		return s.entityRepository.CreateFolder(parentFolder.(*Folder), name)
	}
	return s.entityRepository.CreateFolder(nil, name)
}
func (s *Service) CreateFile(name string, parentId int, content []byte) *File {
	parentFolder := s.entityRepository.GetEntityById(parentId)
	if parentFolder != nil {
		return s.entityRepository.CreateFile(parentFolder.(*Folder), name, content)
	}
	return s.entityRepository.CreateFile(nil, name, content)
}
func (s *Service) AllChildsOfFolder(id int) []string {
	folder := s.entityRepository.GetEntityById(id)
	if folder == nil {
		return nil
	}
	childs := folder.(*Folder).GetChildern()
	names := make([]string, 0)
	for _, child := range childs {
		names = append(names, child.GetName())
	}
	return names
}
func (s *Service) MoveFolderToNewDest(id int, newParentId int) bool {
	entity := s.entityRepository.GetEntityById(id)
	newParent := s.entityRepository.GetEntityById(newParentId)
	if newParent == nil {
		return false
	}
	s.entityRepository.MoveEntityToPath(entity.GetID(), newParent.(*Folder))
	return true
}
