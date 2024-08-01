package middleware

import "os"

type FileInfo struct {
	Name  string
	IsDir bool
}

func GetFilesStruct(files []os.FileInfo) []FileInfo {
	var fileInfos []FileInfo

	for _, entry := range files {
		fileInfos = append(fileInfos, FileInfo{
			Name:  entry.Name(),
			IsDir: entry.IsDir(),
		})
	}

	fileInfos = append(fileInfos, FileInfo{
		Name:  "..",
		IsDir: true,
	})

	return fileInfos
}

func DeleteFile(dir string, file string) {
  err := os.Remove(dir + "/" + file)
  if err != nil {
    panic(err)
  }
}
