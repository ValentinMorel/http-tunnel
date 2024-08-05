package handlers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
)

var (
	Tunnels = map[int]chan Tunnel{}
)

type Tunnel struct {
	W    io.Writer
	Done chan struct{}
}

func List(w http.ResponseWriter, r *http.Request) {
	idQuery := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idQuery)
	_, ok := Tunnels[id]
	if !ok {
		w.Write([]byte("tunnel not found"))
		return
	}

	dirPath := r.URL.Query().Get("dir")
	if dirPath == "" {
		dirPath = "."
	}

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		http.Error(w, "Unable to list directory", http.StatusInternalServerError)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Fprintf(w, "<a href=\"/list?id=%d&dir=%s\">%s/</a><br>", id, filepath.Join(dirPath, file.Name()), file.Name())
		} else {
			fmt.Fprintf(w, "<a href=\"/download?id=%d&file=%s\" target=\"_blank\">%s</a><br>", id, filepath.Join(dirPath, file.Name()), file.Name())
		}
	}
}
