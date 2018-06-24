package main

type question struct {
	Body   string
	Answer string
	ID     int32
}

// Questions is an array of questions to be loaded in by the server
var Questions = []question{
	{
		"gRPC",
		"gRPC Remote Procedure Call",
		1,
	},
	{
		"CORBA",
		"Common Object Request Broker Architecture",
		2,
	},
	{
		"SOAP",
		"Simple Object Access Protocol",
		3,
	},
	{
		"REST",
		"Representational State Transfer",
		4,
	},
	{
		"WSDL",
		"Web Services Description Language",
		5,
	},
	{
		"RMI",
		"Remote Method Invocation",
		6,
	},
	{
		"XML",
		"Extensible Markup Language",
		7,
	},
}
