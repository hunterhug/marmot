package tool

import (
	"errors"
	"fmt"
	"github.com/hunterhug/marmot/expert"
	"github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/marmot/util"
	"net/url"
	"strings"
)

// Download one HTML page's all pictures
// @URL: http://image.baidu.com
// @SaveDir /home/images
// @ProxyAddress : "socks5://127.0.0.1:1080"
func DownloadHTMLPictures(URL string, SaveDir string, MinerNum int, ProxyAddress interface{}) error {

	// Check valid
	_, err := url.Parse(URL)
	if err != nil {
		return err
	}

	// New a worker to get url
	worker, err := miner.New(ProxyAddress)
	if err != nil {
		return err
	}

	result, err := worker.SetUrl(URL).SetUa(miner.RandomUa()).Get()
	if err != nil {
		return err
	}

	// Find all picture
	pictures := expert.FindPicture(string(result))

	return DownloadURLPictures(pictures, SaveDir, MinerNum, worker)
}

// Download pictures faster!
func DownloadURLPictures(PictureUrls []string, SaveDir string, MinerNum int, initWorker *miner.Worker) error {
	// Empty, What a pity!
	if len(PictureUrls) == 0 {
		return errors.New("empty")
	}

	// Make dir!
	err := util.MakeDir(SaveDir)
	if err != nil {
		return err
	}

	// Divide pictures into several worker
	xxx, _ := util.DivideStringList(PictureUrls, MinerNum)

	// Chanel to info exchange
	chs := make(chan int, len(PictureUrls))

	// Go at the same time
	for num, pictureList := range xxx {
		// Clone new worker
		workerPicture := initWorker.Clone()
		workerPicture.SetUa(miner.RandomUa())

		// Go save picture!
		go func(images []string, worker *miner.Worker, num int) {
			for _, img := range images {

				// Check, May be Pass
				_, err := url.Parse(img)
				if err != nil {
					continue
				}

				// Change Name of our picture
				filename := strings.Replace(util.ValidFileName(img), "#", "_", -1)

				// Exist?
				if util.FileExist(SaveDir + "/" + filename) {
					fmt.Println("File Exist：" + SaveDir + "/" + filename)
					chs <- 0
				} else {

					// Not Exsit?
					imgsrc, e := worker.SetUrl(img).Get()
					if e != nil {
						fmt.Println("Download " + img + " error:" + e.Error())
						chs <- 0
						return
					}

					// Save it!
					e = util.SaveToFile(SaveDir+"/"+filename, imgsrc)
					if e == nil {
						fmt.Printf("SP%d: Keep in %s/%s\n", num, SaveDir, filename)
					}
					chs <- 1
				}
			}
		}(pictureList, workerPicture, num)
	}

	// Every picture should return
	for i := 0; i < len(PictureUrls); i++ {
		<-chs
	}

	return nil
}
