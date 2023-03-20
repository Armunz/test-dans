package http

import "test-dans/model"

var dummyJobList = []model.Job{
	{
		ID:          "test1",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "lalala yeyeye Java lalala",
	},
	{
		ID:          "test2",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "lalala yeyeye java yeyeye",
	},
	{
		ID:          "test3",
		Type:        "Freelance",
		Location:    "Surabaya",
		Description: "hehehe hahaha golang yayaya huhuhu",
	},
	{
		ID:          "test4",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "hahaha huhu hahaha javascript hehe",
	},
	{
		ID:          "test5",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "hahaha huhu hahaha java hehe",
	},
	{
		ID:          "test6",
		Type:        "Freelance",
		Location:    "Surabaya",
		Description: "hahaha huhu hahaha ruby hehe",
	},
	{
		ID:          "test7",
		Type:        "Contract",
		Location:    "Jakarta",
		Description: "hahaha huhu hahaha react hehe",
	},
}

var wantResultFullTimeOnly = []model.Job{
	{
		ID:          "test1",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "lalala yeyeye Java lalala",
	},
	{
		ID:          "test2",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "lalala yeyeye java yeyeye",
	},
	{
		ID:          "test4",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "hahaha huhu hahaha javascript hehe",
	},
	{
		ID:          "test5",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "hahaha huhu hahaha java hehe",
	},
}

var wantResultFullTimeJava = []model.Job{
	{
		ID:          "test1",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "lalala yeyeye Java lalala",
	},
	{
		ID:          "test2",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "lalala yeyeye java yeyeye",
	},
	{
		ID:          "test4",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "hahaha huhu hahaha javascript hehe",
	},
	{
		ID:          "test5",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "hahaha huhu hahaha java hehe",
	},
}

var wantResultFullTimeJavaSurabaya = []model.Job{
	{
		ID:          "test1",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "lalala yeyeye Java lalala",
	},
	{
		ID:          "test2",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "lalala yeyeye java yeyeye",
	},
	{
		ID:          "test4",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "hahaha huhu hahaha javascript hehe",
	},
	{
		ID:          "test5",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "hahaha huhu hahaha java hehe",
	},
}

var wantResultFullTimeFalseSurabaya = []model.Job{
	{
		ID:          "test1",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "lalala yeyeye Java lalala",
	},
	{
		ID:          "test2",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "lalala yeyeye java yeyeye",
	},
	{
		ID:          "test3",
		Type:        "Freelance",
		Location:    "Surabaya",
		Description: "hehehe hahaha golang yayaya huhuhu",
	},
	{
		ID:          "test4",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "hahaha huhu hahaha javascript hehe",
	},
	{
		ID:          "test5",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "hahaha huhu hahaha java hehe",
	},
	{
		ID:          "test6",
		Type:        "Freelance",
		Location:    "Surabaya",
		Description: "hahaha huhu hahaha ruby hehe",
	},
}

var wantResultJakartaReact = []model.Job{
	{
		ID:          "test7",
		Type:        "Contract",
		Location:    "Jakarta",
		Description: "hahaha huhu hahaha react hehe",
	},
}

// this is used to test when pagination is 3 and page param is 2
var wantResultLastPage = []model.Job{
	{
		ID:          "test7",
		Type:        "Contract",
		Location:    "Jakarta",
		Description: "hahaha huhu hahaha react hehe",
	},
}

var wantResultMiddlePage = []model.Job{
	{
		ID:          "test4",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "hahaha huhu hahaha javascript hehe",
	},
	{
		ID:          "test5",
		Type:        "Full Time",
		Location:    "Surabaya",
		Description: "hahaha huhu hahaha java hehe",
	},
	{
		ID:          "test6",
		Type:        "Freelance",
		Location:    "Surabaya",
		Description: "hahaha huhu hahaha ruby hehe",
	},
}
