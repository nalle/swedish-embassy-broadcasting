package command

import (
	"github.com/hackverket/swedish-embassy-broadcasting/fetch"
	"github.com/hackverket/swedish-embassy-broadcasting/mpd"
)

func QueueSong(url string) {
	au, _ := fetch.DownloadAudio(url)

	mpd.M.Add(au.Path)
}
