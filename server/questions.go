package main

type question struct {
	Question string
	Answer   string
}

// Questions is an array of questions to be loaded in by the server
var Questions = []question{
	{
		"gRPC",
		"gRPC Remote Procedure Call",
	},
	{
		"CORBA",
		"Common Object Request Broker Architecture",
	},
	{
		"SOAP",
		"Simple Object Access Protocol",
	},
	{
		"REST",
		"Representational State Transfer",
	},
	{
		"WSDL",
		"Web Services Description Language",
	},
	{
		"RMI",
		"Remote Method Invocation",
	},
	{
		"XML",
		"Extensible Markup Language",
	},
}
