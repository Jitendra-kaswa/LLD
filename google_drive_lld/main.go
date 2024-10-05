package main

import (
	"fmt"
	"google-drive-lld/src"
)

func main() {
	service := src.NewService()

	// Creating a more complex nested structure.
	root := service.CreateFolder("Root", -1)

	// Level 1 Folders under Root.
	folderA := service.CreateFolder("Subfolder A", root.GetID())
	folderB := service.CreateFolder("Subfolder B", root.GetID())
	folderC := service.CreateFolder("Subfolder C", root.GetID())

	// Level 2 Folders under Subfolder A.
	subfolderA1 := service.CreateFolder("Subfolder A1", folderA.GetID())
	subfolderA2 := service.CreateFolder("Subfolder A2", folderA.GetID())

	// Level 3 Folders under Subfolder A1.
	subfolderA1_1 := service.CreateFolder("Subfolder A1.1", subfolderA1.GetID())
	service.CreateFolder("Subfolder A1.2", subfolderA1.GetID())

	// Level 3 Folders under Subfolder A2.
	service.CreateFolder("Subfolder A2.1", subfolderA2.GetID())

	// Level 2 Folders under Subfolder B.
	subfolderB1 := service.CreateFolder("Subfolder B1", folderB.GetID())

	// Level 3 Folders under Subfolder B1.
	service.CreateFolder("Subfolder B1.1", subfolderB1.GetID())

	// Files in various folders.
	service.CreateFile("File 1", root.GetID(), []byte("Content of File 1"))
	service.CreateFile("File 2", folderA.GetID(), []byte("Content of File 2"))
	service.CreateFile("File A1.1 Document", subfolderA1_1.GetID(), []byte("Content of A1.1 Document"))

	service.CreateFile("File B Document", subfolderB1.GetID(), []byte("Content of B Document"))

	// Print expected and actual outputs for various operations.

	fmt.Printf("Childs of Root (expected [Subfolder A, Subfolder B, Subfolder C, File 1]): %v\n", service.AllChildsOfFolder(root.GetID()))

	fmt.Printf("Childs of Subfolder A (expected [Subfolder A1, Subfolder A2, File 2]): %v\n", service.AllChildsOfFolder(folderA.GetID()))

	fmt.Printf("Childs of Subfolder A1 (expected [Subfolder A1.1, Subfolder A1.2]): %v\n", service.AllChildsOfFolder(subfolderA1.GetID()))

	fmt.Printf("Childs of Subfolder B (expected [Subfolder B1]): %v\n", service.AllChildsOfFolder(folderB.GetID()))

	fmt.Printf("Childs of Subfolder B1 (expected [Subfolder B1.1, File B Document]): %v\n", service.AllChildsOfFolder(subfolderB1.GetID()))

	// Move File 2 to Subfolder C.
	file2 := service.CreateFile("Temporary File 2", folderC.GetID(), []byte{})
	service.MoveFolderToNewDest(file2.GetID(), folderC.GetID())

	fmt.Printf("\nAfter moving Temporary File 2 to Subfolder C:\n")

	fmt.Printf("Childs of Root (expected [Subfolder A, Subfolder B, Subfolder C, File 1]): %v\n", service.AllChildsOfFolder(root.GetID()))

	fmt.Printf("Childs of Subfolder C (expected [Temporary File 2]): %v\n", service.AllChildsOfFolder(folderC.GetID()))

	fmt.Printf("Childs of Subfolder A (expected [Subfolder A1, Subfolder A2, File 2]): %v\n", service.AllChildsOfFolder(folderA.GetID()))
}
