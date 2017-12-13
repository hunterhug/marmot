package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/hunterhug/marmot/expert"
	"github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/parrot/util"
)

// Num of miner, We can run it at the same time to crawl data fast
var MinerNum = 5

// You can update this decide whether to proxy
var ProxyAddress interface{}

func main() {
	// You can Proxy!
	// ProxyAddress = "socks5://127.0.0.1:1080"

	fmt.Println(`
		Welcome: Input "url" and picture keep "dir"

		
		`)
	for {
		fmt.Println("---------------------------------------------")
		url := util.Input(`URL(Like: "http://publicdomainarchive.com")`, "http://publicdomainarchive.com")
		dir := util.Input(`DIR(Default: "./picture")`, "./picture")
		fmt.Printf("You will keep %s picture in dir %s\n", url, dir)
		fmt.Println("---------------------------------------------")

		// Start Catch
		err := CatchPicture(url, dir)
		if err != nil {
			fmt.Println("Error:" + err.Error())
		}
	}
}

// Come on!
func CatchPicture(picture_url string, dir string) error {
	// Check valid
	_, err := url.Parse(picture_url)
	if err != nil {
		return err
	}

	// Make dir!
	err = util.MakeDir(dir)
	if err != nil {
		return err
	}

	// New a worker to get url
	worker, _ := miner.New(ProxyAddress)

	result, err := worker.SetUrl(picture_url).SetUa(miner.RandomUa()).Get()
	if err != nil {
		return err
	}

	// Find all picture
	pictures := expert.FindPicture(string(result))

	// Empty, What a pity!
	if len(pictures) == 0 {
		return errors.New("empty")
	}

	// Devide pictures into several worker
	xxx, _ := util.DevideStringList(pictures, MinerNum)

	// Chanel to info exchange
	chs := make(chan int, len(pictures))

	// Go at the same time
	for num, imgs := range xxx {

		// Get pool miner
		worker_picture, ok := miner.Pool.Get(util.IS(num))
		if !ok {
			// No? set one!
			worker_temp, _ := miner.New(ProxyAddress)
			worker_picture = worker_temp
			worker_temp.SetUa(miner.RandomUa())
			miner.Pool.Set(util.IS(num), worker_temp)
		}

		// Go save picture!
		go func(imgs []string, worker *miner.Worker, num int) {
			for _, img := range imgs {

				// Check, May be Pass
				_, err := url.Parse(img)
				if err != nil {
					continue
				}

				// Change Name of our picture
				filename := strings.Replace(util.ValidFileName(img), "#", "_", -1)

				// Exist?
				if util.FileExist(dir + "/" + filename) {
					fmt.Println("File Exist：" + dir + "/" + filename)
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
					e = util.SaveToFile(dir+"/"+filename, imgsrc)
					if e == nil {
						fmt.Printf("SP%d: Keep in %s/%s\n", num, dir, filename)
					}
					chs <- 1
				}
			}
		}(imgs, worker_picture, num)
	}

	// Every picture should return
	for i := 0; i < len(pictures); i++ {
		<-chs
	}

	return nil
}
