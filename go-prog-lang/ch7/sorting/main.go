package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// sample data
var tracks = []*Track{
	{"Piano Concerto No.1 Op.23 1st Mov", "P.I.Tchaikovsky", "CASIO GP500 CONSERT PLAY", 2012, length("3m38s")},
	{"Prelude Op.28-15 RainDrop", "F.F.Chopin", "CASIO GP500 CONSERT PLAY", 1992, length("3m37s")},
	{"Canon", "J.Pachelbel", "CASIO GP500 CONSERT PLAY", 2007, length("4m36s")},
	{"Jesus, Bleibet Meine Freude", "J.S.Bach", "CASIO GP500 CONSERT PLAY", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 2, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func init() {
	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values))

	sort.Ints(values)
	fmt.Println(values)

	fmt.Println(sort.IntsAreSorted(values))
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
}

func main() {
	fmt.Println("byArtist:")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	fmt.Println("\nReverse(byArtist):")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	fmt.Println("\nbyYear:")
	sort.Sort(byYear(tracks))
	printTracks(tracks)

	fmt.Println("\nCustom:")
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)
}

/**
byArtist:
Title                              Artist           Album                     Year  Length
-----                              ------           -----                     ----  ------
Prelude Op.28-15 RainDrop          F.F.Chopin       CASIO GP500 CONSERT PLAY  1992  3m37s
Canon                              J.Pachelbel      CASIO GP500 CONSERT PLAY  2007  4m36s
Jesus, Bleibet Meine Freude        J.S.Bach         CASIO GP500 CONSERT PLAY  2011  4m24s
Piano Concerto No.1 Op.23 1st Mov  P.I.Tchaikovsky  CASIO GP500 CONSERT PLAY  2012  3m38s

Reverse(byArtist):
Title                              Artist           Album                     Year  Length
-----                              ------           -----                     ----  ------
Piano Concerto No.1 Op.23 1st Mov  P.I.Tchaikovsky  CASIO GP500 CONSERT PLAY  2012  3m38s
Jesus, Bleibet Meine Freude        J.S.Bach         CASIO GP500 CONSERT PLAY  2011  4m24s
Canon                              J.Pachelbel      CASIO GP500 CONSERT PLAY  2007  4m36s
Prelude Op.28-15 RainDrop          F.F.Chopin       CASIO GP500 CONSERT PLAY  1992  3m37s

byYear:
Title                              Artist           Album                     Year  Length
-----                              ------           -----                     ----  ------
Prelude Op.28-15 RainDrop          F.F.Chopin       CASIO GP500 CONSERT PLAY  1992  3m37s
Canon                              J.Pachelbel      CASIO GP500 CONSERT PLAY  2007  4m36s
Jesus, Bleibet Meine Freude        J.S.Bach         CASIO GP500 CONSERT PLAY  2011  4m24s
Piano Concerto No.1 Op.23 1st Mov  P.I.Tchaikovsky  CASIO GP500 CONSERT PLAY  2012  3m38s

Custom:
Title                              Artist           Album                     Year  Length
-----                              ------           -----                     ----  ------
Canon                              J.Pachelbel      CASIO GP500 CONSERT PLAY  2007  4m36s
Jesus, Bleibet Meine Freude        J.S.Bach         CASIO GP500 CONSERT PLAY  2011  4m24s
Piano Concerto No.1 Op.23 1st Mov  P.I.Tchaikovsky  CASIO GP500 CONSERT PLAY  2012  3m38s
Prelude Op.28-15 RainDrop          F.F.Chopin       CASIO GP500 CONSERT PLAY  1992  3m37s
*/
