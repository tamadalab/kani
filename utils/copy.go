package utils

import "path/filepath"
import "os"
import "io"
import "fmt"

func CopyFile(fromDir, fileName, toDir string) error {
    toFile := filepath.Join(toDir, fileName)
    fromFile := filepath.Join(fromDir, fileName)
    if !ExistFile(fromFile) {
        return fmt.Errorf("%s: file not found", fromFile)
    }
    Mkdirs(filepath.Dir(toFile))
    return copyFile(fromFile, toFile)
}

func copyFile(fromFile, toFile string) error {
    if err := os.Link(fromFile, toFile); err != nil {
        return execCopy(fromFile, toFile)
    }
    return nil
}

func execCopy(fromFile, toFile string) error {
    writer, err1 := os.Create(toFile)
    if err1 != nil {
        return err1
    }
    defer writer.Close()
    reader, err2 := os.Open(fromFile)
    if err2 != nil {
        return err2
    }
    defer reader.Close()

    _, err := io.Copy(writer, reader)
    return err
}
