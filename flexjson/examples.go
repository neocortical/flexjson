package flexjson

var Examples = []string{
	skinnyJSON,
	fatJSON,
	nullJSON,
	missingJSON,
}

var skinnyJSON = `
{
	"desc": "Sub-object is an ID", 
	"sprocket": 42
}
`
var fatJSON = `
{
	"desc": "Sub-object is an object", 
	"sprocket": {
		"id": 42,
		"size": "large",
		"gears": 15 
	}
}
`

var nullJSON = `
{
	"desc": "Sub-object is null", 
	"sprocket": null
}
`

var missingJSON = `
{
	"desc": "Sub-object is not present"
}
`
