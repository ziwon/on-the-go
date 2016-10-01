package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

type Track struct {
	Title  string
	Year   int
	Artist string
}

var tracks = []*Track{
	{"AA", 2010, "ZZZ"},
	{"AA", 2005, "YYY"},
	{"BB", 2002, "ZZZ"},
	{"BB", 2002, "XXX"},
	{"BC", 2003, "YYY"},
	{"CB", 2006, "XXX"},
	{"CB", 2002, "ZZZ"},
	{"AA", 2006, "YYY"},
	{"CB", 2007, "ZZZ"},
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 2, 2, ' ', 0)
	fmt.Fprintf(tw, format, "title", "year", "artist")
	fmt.Fprintf(tw, format, "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Year, t.Artist)
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

type Next func(int, []Handler, *Track, *Track) interface{}

func NewSort(handlers ...Handler) Chain {
	return Chain{append(([]Handler)(nil), handlers...)}
}

func (c Chain) Apply(tracks []*Track) []*Track {
	sort.Sort(Sortable{tracks, func(x, y *Track) bool {
		var fn Next
		fn = func(i int, h []Handler, x, y *Track) interface{} {
			if i > len(h)-1 {
				return false
			}

			if v := h[i](x, y); v != nil {
				return v.(bool)
			} else {
				i++
				return fn(i, h, x, y)
			}
		}
		return fn(0, c.handlers, x, y).(bool)
	}})
	return tracks
}

func main() {
	title := func(x, y *Track) interface{} {
		if x.Title != y.Title {
			return x.Title < y.Title
		} else {
			return nil
		}
	}

	year := func(x, y *Track) interface{} {
		if x.Year != y.Year {
			return x.Year < y.Year
		} else {
			return nil
		}
	}

	artist := func(x, y *Track) interface{} {
		if x.Artist != y.Artist {
			return x.Artist < y.Artist
		} else {
			return nil
		}
	}

	fmt.Println("# original")
	printTracks(tracks)
	fmt.Println()

	fmt.Println("# sort by title, year")
	printTracks(NewSort(title, year).Apply(tracks))
	fmt.Println()

	fmt.Println("# sort by year, title, artist")
	printTracks(NewSort(year, title, artist).Apply(tracks))
	fmt.Println()

	fmt.Println("# sort by artist, title, year")
	printTracks(NewSort(artist, title, year).Apply(tracks))
}

/*

# original
title  year  artist
-----  ----  ------
AA     2010  ZZZ
AA     2005  YYY
BB     2002  ZZZ
BB     2002  XXX
BC     2003  YYY
CB     2006  XXX
CB     2002  ZZZ
AA     2006  YYY
CB     2007  ZZZ

# sort by title, year
title  year  artist
-----  ----  ------
AA     2005  YYY
AA     2006  YYY
AA     2010  ZZZ
BB     2002  ZZZ
BB     2002  XXX
BC     2003  YYY
CB     2002  ZZZ
CB     2006  XXX
CB     2007  ZZZ

# sort by year, title, artist
title  year  artist
-----  ----  ------
BB     2002  XXX
BB     2002  ZZZ
CB     2002  ZZZ
BC     2003  YYY
AA     2005  YYY
AA     2006  YYY
CB     2006  XXX
CB     2007  ZZZ
AA     2010  ZZZ

# sort by artist, title, year
title  year  artist
-----  ----  ------
BB     2002  XXX
CB     2006  XXX
AA     2005  YYY
AA     2006  YYY
BC     2003  YYY
AA     2010  ZZZ
BB     2002  ZZZ
CB     2002  ZZZ
CB     2007  ZZZ

*/
