package handlers

import (
	"github.com/yujie0121/canvas/libhttp"
	"html/template"
	"net/http"
	"time"
	"crypto/md5"
	"io"
	"fmt"
	"os"
	"strconv"
	"io/ioutil"
	"github.com/yujie0121/canvas/model"
	"strings"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法

	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		//fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./static/files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}

	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/uploadfile.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}



	images :=getAllImages( r)

	data := model.ImageData{
		Token: token,
		Images: images,
	}

	tmpl.Execute(w, data)
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	imageName := r.FormValue("name")
	f, err := os.Open("./static/files/"+imageName)
	handleErr(err, w)
	buffer, err := ioutil.ReadAll(f)
	handleErr(err, w)
	defer f.Close()
	w.Write(buffer)
}

func handleErr(err error, w http.ResponseWriter)  {
	if  err != nil {
		w.Write([]byte(err.Error()))
	}
}

func getAllImages(r *http.Request) []model.Image {
	var images []model.Image
	files, _ := ioutil.ReadDir("./static/files/")
	for _, f := range files {
		namesli := strings.Split(f.Name(), ".")
		suffix := strings.ToUpper(namesli[1])
		if(len(namesli) == 2 && (suffix == "JPG" || suffix == "JPEG" ||suffix == "PNG")){
			images = append(images, model.Image{
				Name: namesli[0],
				Type: suffix,
				Directory: "static/files/",
				Url: template.URL("http://"+ r.Host+"/image?name="+f.Name()),
			})
		}

	}

	return images
}