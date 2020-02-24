package structs

// RssParent ... basic RSS structure for an RSS Feed
type RssParent struct {
	ID string `rethinkdb:"id,omitempty"`
	Content RssChannel `xml:"channel" rethinkdb:"content"`
}

//RssChannel ... basic RSS Feed structure. Single <channel> element, which contains information about the channel (metadata) and its contents. More about channels in the following URL: https://validator.w3.org/feed/docs/rss2.html
type RssChannel struct {
	Title       string    `xml:"title" rethinkdb:"title,omitempty"`
	Link        string    `xml:"link" rethinkdb:"link,omitempty"`
	Description string    `xml:"description" rethinkdb:"description,omitempty"`
	Category    string    `xml:"category" rethinkdb:"category,omitempty"`
	Item        []RssItem `xml:"item" rethinkdb:"item,omitempty"`
}

//RssItem ... An item may represent a "story" -- much like a story in a newspaper or magazine; if so its description is a synopsis of the story, and the link points to the full story. An item may also be complete in itself, if so, the description contains the text (entity-encoded HTML is allowed), and the link and title may be omitted. All elements of an item are optional, however at least one of title or description must be present. More can be read on the following URL: https://validator.w3.org/feed/docs/rss2.html
type RssItem struct {
	Title           string `xml:"title" rethinkdb:"title,omitempty"`
	Link            string `xml:"link" rethinkdb:"link,omitempty"`
	Description     string `xml:"description" rethinkdb:"description,omitempty"`
	Author          string `xml:"author" rethinkdb:"author,omitempty"`
	Category        string `xml:"category" rethinkdb:"category,omitempty"`
	Comments        string `xml:"comments" rethinkdb:"comments,omitempty"`
	GUID            string `xml:"guid" rethinkdb:"guid,omitempty"`
	Source          string `xml:"source" rethinkdb:"source,omitempty"`
	PublicationDate string `xml:"pubDate" rethinkdb:"pubDate,omitempty"`
}


// // RssParentRethinkdb ... basic RSS structure for an RSS Feed
// type RssParentRethinkdb struct {
// 	ID string `rethinkdb:"id,omitempty"`
// 	Content RssChannel `rethinkdb:"content"`
// }

// //RssChannelRethinkdb ... basic RSS Feed structure. Single <channel> element, which contains information about the channel (metadata) and its contents. More about channels in the following URL: https://validator.w3.org/feed/docs/rss2.html
// type RssChannelRethinkdb struct {
// 	Title       string    `rethinkdb:"title,omitempty"`
// 	Link        string    `rethinkdb:"link,omitempty"`
// 	Description string    `rethinkdb:"description,omitempty"`
// 	Category    string    `rethinkdb:"category,omitempty"`
// 	Item        []RssItemRethinkdb `rethinkdb:"item,omitempty"`
// }

// //RssItemRethinkdb ... An item may represent a "story" -- much like a story in a newspaper or magazine; if so its description is a synopsis of the story, and the link points to the full story. An item may also be complete in itself, if so, the description contains the text (entity-encoded HTML is allowed), and the link and title may be omitted. All elements of an item are optional, however at least one of title or description must be present. More can be read on the following URL: https://validator.w3.org/feed/docs/rss2.html
// type RssItemRethinkdb struct {
// 	Title           string `rethinkdb:"title,omitempty"`
// 	Link            string `rethinkdb:"link,omitempty"`
// 	Description     string `rethinkdb:"description,omitempty"`
// 	Author          string `rethinkdb:"author,omitempty"`
// 	Category        string `rethinkdb:"category,omitempty"`
// 	Comments        string `rethinkdb:"comments,omitempty"`
// 	GUID            string `rethinkdb:"guid,omitempty"`
// 	Source          string `rethinkdb:"source,omitempty"`
// 	PublicationDate string `rethinkdb:"pubDate,omitempty"`
// }
