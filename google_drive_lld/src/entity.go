package src

// Entity can be folder of file
type IEntity interface {
	GetID() int
	GetName() string
	GetPath() string
	GetParent() *Folder
	SetParent(parent *Folder) error
	IsFolder() bool
}

type File struct {
	id      int
	name    string
	path    string
	parent  *Folder
	content []byte
}

func NewFile(id int, name string, parent *Folder, content []byte) *File {
	path := name
	if parent != nil {
		path = parent.GetPath() + path
	}
	return &File{
		id:      id,
		name:    name,
		parent:  parent,
		content: content,
		path:    path,
	}
}

func (f File) GetID() int {
	return f.id
}
func (f File) GetName() string {
	return f.name
}
func (f File) GetPath() string {
	return f.path
}
func (f File) GetParent() *Folder {
	return f.parent
}
func (f *File) SetParent(parent *Folder) error {
	path := f.name
	if parent != nil {
		path = parent.GetPath() + path
	}
	f.path = path
	f.parent = parent
	return nil
}
func (f File) IsFolder() bool {
	return false
}

type Folder struct {
	id       int
	name     string
	path     string
	parent   *Folder
	Children []IEntity
}

func NewFolder(id int, name string, parent *Folder) *Folder {
	path := name
	if parent != nil {
		path = parent.GetPath() + path
	}
	return &Folder{
		id:     id,
		name:   name,
		parent: parent,
		path:   path,
	}
}

func (f Folder) GetID() int {
	return f.id
}
func (f Folder) GetName() string {
	return f.name
}
func (f Folder) GetPath() string {
	return f.path
}
func (f Folder) GetParent() *Folder {
	return f.parent
}
func (f *Folder) SetParent(parent *Folder) error {
	path := f.name
	if parent != nil {
		path = parent.GetPath() + path
	}
	f.path = path
	f.parent = parent
	return nil
}

func (f Folder) IsFolder() bool {
	return true
}

func (f *Folder) AddChildren(child IEntity) error {
	f.Children = append(f.Children, child)
	child.SetParent(f)
	return nil
}

func (f *Folder) DeleteChildren(child IEntity) error {
	newChilds := make([]IEntity, 0)
	for _, entity := range f.Children {
		if entity.GetID() != child.GetID() {
			newChilds = append(newChilds, entity)
		}
	}
	f.Children = newChilds
	return nil
}

func (f *Folder) GetChildern() []IEntity {
	return f.Children
}

type IEntityRepository interface {
	GetEntityById(id int) IEntity
	CreateFolder(parent *Folder, name string) *Folder
	CreateFile(parent *Folder, name string, content []byte) *File
	DeleteEntity(id int) error
	MoveEntityToPath(id int, newParent *Folder) error
}

type EntityRepository struct {
	currentEntityId int
	entityMap       map[int]IEntity
}

func NewEntityRepository() IEntityRepository {
	return &EntityRepository{
		entityMap:       make(map[int]IEntity),
		currentEntityId: 1,
	}
}

func (er EntityRepository) GetEntityById(id int) IEntity {
	if eneity, exists := er.entityMap[id]; exists {
		return eneity
	}
	return nil
}

func (er *EntityRepository) CreateFolder(parent *Folder, name string) *Folder {
	newFolder := NewFolder(er.currentEntityId, name, parent)
	er.entityMap[er.currentEntityId] = newFolder
	if parent != nil {
		parent.AddChildren(newFolder)
	}
	er.currentEntityId++
	return newFolder
}
func (er *EntityRepository) CreateFile(parent *Folder, name string, content []byte) *File {
	newFile := NewFile(er.currentEntityId, name, parent, content)
	er.entityMap[er.currentEntityId] = newFile
	if parent != nil {
		parent.AddChildren(newFile)
	}
	er.currentEntityId++
	return newFile
}
func (er EntityRepository) DeleteEntity(id int) error {
	entity := er.entityMap[id]
	if entity.IsFolder() {
		folder := entity.(*Folder)
		for _, child := range folder.GetChildern() {
			er.DeleteEntity(child.GetID())
		}
	}
	delete(er.entityMap, id)
	return nil
}
func (er EntityRepository) MoveEntityToPath(id int, newParent *Folder) error {
	entity := er.entityMap[id]
	entity.GetParent().DeleteChildren(entity)
	entity.SetParent(newParent)
	entity.GetParent().AddChildren(entity)
	if entity.IsFolder() {
		folder := entity.(*Folder)
		for _, child := range folder.GetChildern() {
			er.MoveEntityToPath(child.GetID(), entity.(*Folder))
		}
	}
	return nil
}
