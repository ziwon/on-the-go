package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title string
	Year  int
}

var tracks = []*Track{
	{"AA", 2010},
	{"AA", 2005},
	{"BB", 2002},
	{"BA", 2002},
	{"BC", 2003},
	{"CB", 2006},
	{"CB", 2002},
	{"AA", 2006},
	{"CB", 2007},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 2, 2, ' ', 0)
	fmt.Fprintf(tw, format, "title", "year")
	fmt.Fprintf(tw, format, "-----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Year)
	}
	tw.Flush()
}

type Sortable struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (s Sortable) Len() int           { return len(s.t) }
func (s Sortable) Less(i, j int) bool { return s.less(s.t[i], s.t[j]) }
func (s Sortable) Swap(i, j int)      { s.t[i], s.t[j] = s.t[j], s.t[i] }

type Handler func(x, y *Track) interface{}

type Chain struct {
	handlers []Handler
}

func NewSort(handlers ...Handler) Chain {
	return Chain{append(([]Handler)(nil), handlers...)}
}

func (c Chain) Apply(tracks []*Track) []*Track {
	sort.Sort(Sortable{tracks, func(x, y *Track) bool {
		for i := 0; i < len(c.handlers)-1; i++ {
			v := c.handlers[i](x, y)
			switch v {
			case nil:
				i++
				return c.handlers[i](x, y).(bool)
			}
			return v.(bool)
		}
		return false

	}})
	return tracks
}

func main() {
	Title := func(x, y *Track) interface{} {
		if x.Title != y.Title {
			return x.Title < y.Title
		} else {
			return nil
		}
	}

	Year := func(x, y *Track) interface{} {
		if x.Year != y.Year {
			return x.Year < y.Year
		} else {
			return nil
		}
	}

	fmt.Println("# original")
	printTracks(tracks)
	fmt.Println()

	fmt.Println("# sort by title, year")
	printTracks(NewSort(Title, Year).Apply(tracks))
	fmt.Println()

	fmt.Println("# sort by year, title")
	printTracks(NewSort(Year, Title).Apply(tracks))
}
