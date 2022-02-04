package main

const (
	//nolint: lll
	physicsgg          = "https://physicsgg.me/category/%ce%b1%cf%83%cf%84%cf%81%ce%bf%cf%86%cf%85%cf%83%ce%b9%ce%ba%ce%b7/feed/"
	spaceCom           = "https://www.space.com/feeds/all"
	astronioAstroNews  = "https://www.astronio.gr/archives/category/astronews/feed"
	astronioDiastimiki = "https://www.astronio.gr/archives/category/diastimiki/feed"
	cnnDiastimaTag     = "https://www.cnn.gr/tag/diastima?format=feed&type=rss"
	earthSky           = "https://earthsky.org/space/feed/"
	universeToday      = "https://www.universetoday.com/feed/"
	newScientist       = "https://www.newscientist.com/subject/space/feed/"
	everyDayAstronaut  = "https://everydayastronaut.com/feed/"
	astronomyNow       = "https://astronomynow.com/feed/"
	sciTechDaily       = "https://scitechdaily.com/news/space/feed/"
	skyAndTelescope    = "https://skyandtelescope.org/astronomy-news/feed/"
	astronomyNews      = "https://news.google.com/rss/search?q=site:astronomy.com" // Aπό το Google, γιατί δεν έχει RSS
)

// nolint: gochecknoglobals
var feeds = []string{
	astronomyNews,
	skyAndTelescope,
	sciTechDaily,
	astronomyNow,
	everyDayAstronaut,
	newScientist,
	universeToday,
	earthSky,
	cnnDiastimaTag,
	spaceCom,
	astronioAstroNews,
	astronioDiastimiki,
	physicsgg,
}
