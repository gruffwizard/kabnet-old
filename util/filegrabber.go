package util

import (
	"fmt"
	"log"
	"time"

	"github.com/cavaliercoder/grab"
)

func FetchFile(dir string, url string,file string) {

  if FileExists(dir+"/"+file) {
    log.Printf("image file %s already exists",file)
    return 
  }
	client := grab.NewClient()
	req, _ := grab.NewRequest(dir, url+"/"+file)

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(5000 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size(),
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
    log.Fatalf("Download failed: %v\n", err)
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)

}
