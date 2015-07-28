package jenkinsrss

type Link struct {
	Type string `xml:"type,attr"`
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
}
type Entry struct {
	Title     string `xml:"title"`
	Link      Link   `xml:"link"`
	Id        string `xml:"id"`
	Published string `xml:"published"`
	Updated   string `xml:"updated"`
}
type Feed struct {
	Title   string  `xml:"title"`
	Link    Link    `xml:"link"`
	Id      string  `xml:"id"`
	Updated string  `xml:"updated"`
	Entries []Entry `xml:"entry"`
}
