package main

import (
	"encoding/json"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"os"
	"strconv"
)

func downloadFont(family *FontFamily, font *Font) {
	filename := *family.Name + "-" + strconv.Itoa(int(*font.Weight.Start)) + "-"
	if 1.0-*font.Italic.Start < 1e-6 {
		filename += "normal"
	} else {
		filename += "italic"
	}
	filename += ".ttf"
	url := "https://fonts.gstatic.com/s/a/" + hashToString(font.File.Hash) + ".ttf"
	out, err := os.Create("fonts/" + filename)
	if err != nil {
		logrus.Error(err)
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		logrus.Error(err)
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		logrus.Error(err)
	}
	fonts = append(fonts, FontFile{
		Family:   *family.Name,
		Weight:   int(*font.Weight.Start),
		Italic:   1.0-*font.Italic.Start < 1e-6,
		Filename: filename,
	})
	_ = bar.Add(1)
	logrus.Infof("Downloaded %s", filename)
}

func hashToString(bytes []byte) string {
	fileName := ""
	for _, b := range bytes {
		convertedByte := strconv.FormatInt(int64(b), 16)
		convertedByte = fmt.Sprintf("%02s", convertedByte)
		fileName += convertedByte
	}
	return fileName
}

var bar *progressbar.ProgressBar

type FontFile struct {
	Family   string
	Weight   int
	Italic   bool
	Filename string
}

var fonts []FontFile

func main() {
	if _, err := os.Stat("fonts"); os.IsNotExist(err) {
		err = os.Mkdir("fonts", os.ModePerm)
		if err != nil {
			logrus.Panic(err)
		}
	}
	body, err := os.ReadFile("directory007.pb")
	if err != nil {
		logrus.Panic(err)
	}
	var directory Directory
	err = proto.Unmarshal(body, &directory)
	if err != nil {
		logrus.Panic(err)
	}
	size := 0
	for _, family := range directory.Family {
		size = size + len(family.Fonts)
	}
	bar = progressbar.Default(int64(size))
	for _, family := range directory.Family {
		for _, font := range family.Fonts {
			downloadFont(family, font)
		}
	}
	file, err := json.MarshalIndent(fonts, "", " ")
	if err != nil {
		logrus.Error(err)
	}
	err = os.WriteFile("fonts.json", file, 0644)
	if err != nil {
		logrus.Error(err)
	}
}
