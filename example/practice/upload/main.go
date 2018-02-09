package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println(32 << 20)
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":1789", nil)
}

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprintln(w, "upload ok!")
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(tpl))
}

const tpl = `<html>  
<head>  
<title>上传文件</title>  
</head>  
<body>  
<form enctype="multipart/form-data" action="/upload" method="post">  
 <input type="file" name="uploadfile" />  
 <input type="hidden" name="token" value="xxx"/>  
 <input type="submit" value="upload" />  
</form>  
</body>  
</html>`
