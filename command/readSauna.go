package command

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"

	"github.com/hackverket/swedish-embassy-broadcasting/motuavb"
	"github.com/hackverket/swedish-embassy-broadcasting/polly"
	uuid "github.com/satori/go.uuid"
)

func ReadSauna() {
	saunaT := getTemperature()
	saunaString := "The sauna is now " + saunaT + " degrees celsius. Time to get in there!"
	ttsFile := polly.GetTTS(saunaString, "Joey")
	// sc := motuavb.Connect("10.44.22.107")
	// sc.FadeChannelVolume(8, 0.05)
	u,_ := uuid.NewV4()
	wavpath := path.Join(os.Getenv("DUMP_PATH"), u.String()) + ".wav"
	lol := exec.Command("ffmpeg", "-i", ttsFile, "-ar", "44100", "-ac", "2", wavpath)
	g, berr := lol.Output()
	fmt.Println(g, berr)

	paplayArgs := append([]string{
		"-s", "127.0.0.1",
		"--channel-map=aux0,aux1",
		wavpath,
	})

	fmt.Println(paplayArgs)

	cmd := exec.Command("paplay", paplayArgs...)
	o, err := cmd.Output()
	fmt.Println(o, err)
	// sc.FadeChannelVolume(8, 0.8)
}

func getTemperature() string {

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	resp, err := http.Get("http://www.tarlab.fi/sensors/temperature1")
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	f := string(bytes)

	return f
}
