package main
import (
	"encoding/xml"
	//"io"
	"io/ioutil"
	//"os"
	"fmt"
)

type Query struct {
	Series Show
	// Have to specify where to find episodes since this
	// doesn't match the xml tags of the data that needs to go into it
	EpisodeList []Episode `xml:"Episode"`
}

type Show struct {
	//have to specify where to find series title since 
	// the field of this struct doesn't match the xml tag
	Title string `xml:"SeriesName"`
	SeriesID int
	//keyword map[string]bool
	
}

type Episode struct {
	SeasonNumber	int
	EpisodeNumber 	int
	EpisodeName		string
	FirstAired		string
}

func (s Show) String() string {
	return fmt.Sprintf("%s - %d", s.Title, s.SeriesID)
}

func (e Episode) String() string {
	return fmt.Sprintf("S%02dE%02d - %s - %s", e.SeasonNumber, e.EpisodeNumber, e.EpisodeName, e.FirstAired)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	
	data, err := ioutil.ReadFile("xmltest.xml")
	check(err)
	
	var q Query
	xml.Unmarshal(data, &q)
	
	fmt.Println(q.Series)
	
	//fmt.Println(q.keyword[""])
	for _, episode := range q.EpisodeList {
		fmt.Printf("%s\n", episode)
	}
}
