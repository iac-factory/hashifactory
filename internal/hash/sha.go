package main

import (
    "crypto/sha256"
    "encoding/base64"
    "fmt"
    "io/fs"
    "log"
    "os"
)

func Walker() []string {
    var hashes []string
    var directory = "/Users/io" // os.TempDir()
    filesystem := os.DirFS(directory)

    if exception := fs.WalkDir(filesystem, ".", func(path string, d fs.DirEntry, exception error) error {
        fmt.Println(path, d.Name(), d.IsDir())
        if exception != nil && os.IsPermission(exception) {
            return nil
        } else {
            panic(exception)
        }

        hashable := !d.IsDir()

        if hashable {
            sha, exception := hash(filesystem, path)
            if exception != nil {
                log.Fatal(exception)
            }

            log.Println(path)
            hashes = append(hashes, sha)
        }

        return nil
    }); exception != nil {
        log.Fatal(exception)
    }

    return hashes
}

func hash(system fs.FS, file string) (string, error) {
    data, exception := fs.ReadFile(system, file)
    if exception != nil {
        panic(exception)
    }

    h256 := sha256.New()
    h256.Write([]byte(data))
    sha256base64 := base64.StdEncoding.EncodeToString(h256.Sum(nil)[:])

    return sha256base64, nil
}

func main() {
    Walker()
}
