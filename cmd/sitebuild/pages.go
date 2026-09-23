package main

var staticPages = []page{
	{
		Source: "templates/index.html",
		Target: "index.html",
	},
	{
		Source: "templates/404.html",
		Target: "404.html",
	},

	{
		Source: "templates/pilot.html",
		Target: "pilot/index.html",
	},

	{
		Source: "templates/museum.html",
		Target: "museum/index.html",
	},
}
