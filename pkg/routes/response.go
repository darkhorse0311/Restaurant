package routes

// DefaultLocation are locations given when geoLocations denied
var DefaultLocation = []byte(`[
	{
	  "id": "r4jG613xCL4ispS14rEBnA",
	  "name": "Subway",
	  "coordinates": {
		"latitude": 40.7163137197495,
		"longitude": -74.004714936018
	  },
	  "photos": [
		"https://s3-media1.fl.yelpcdn.com/bphoto/TWRxx3GueDGljuZWExE5eA/o.jpg"
	  ],
	  "distance": 414.39746949756113,
	  "r_id": 1
	},
	{
	  "id": "p24aXNna282LwdSk0n-4UQ",
	  "name": "Subway",
	  "coordinates": {
		"latitude": 40.71289,
		"longitude": -74.00758
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/y5kSjjmza9da5Y4lGyNHxg/o.jpg"
	  ],
	  "distance": 154.94281540312298,
	  "r_id": 1
	},
	{
	  "id": "jPIZ3FR5LNcwPuUHi2Fe4g",
	  "name": "McDonald's",
	  "coordinates": {
		"latitude": 40.70944,
		"longitude": -74.01012
	  },
	  "photos": [
		"https://s3-media1.fl.yelpcdn.com/bphoto/uCgd-65G_jhB9PiGxY8uEA/o.jpg"
	  ],
	  "distance": 507.50127180505456,
	  "r_id": 4
	},
	{
	  "id": "pymPEcd_X1A1pvXHQeg0_w",
	  "name": "Wendy's",
	  "coordinates": {
		"latitude": 40.71013,
		"longitude": -74.00816
	  },
	  "photos": [
		"https://s3-media3.fl.yelpcdn.com/bphoto/v30Yqk15dCCaeg8RJpOU8A/o.jpg"
	  ],
	  "distance": 349.1855881243587,
	  "r_id": 4
	},
	{
	  "id": "nT3rCBXEeORDUEj0rlWDUA",
	  "name": "Subway",
	  "coordinates": {
		"latitude": 40.71596,
		"longitude": -74.00996
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/fAOvbeERGjiW1vKr2vFf3A/o.jpg"
	  ],
	  "distance": 486.09432143835386,
	  "r_id": 5
	},
	{
	  "id": "mnEMipzEekqLWIogjAB_zA",
	  "name": "Subway",
	  "coordinates": {
		"latitude": 40.707682,
		"longitude": -74.005945
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/kAXa2RNeFZUz6SJaT3vXGg/o.jpg"
	  ],
	  "distance": 569.1138846509674,
	  "r_id": 5
	},
	{
	  "id": "LAo4PdkUMwRF88x5TnG3wA",
	  "name": "Subway",
	  "coordinates": {
		"latitude": 40.7151328772306,
		"longitude": -74.0077042579651
	  },
	  "photos": [
		"https://s3-media3.fl.yelpcdn.com/bphoto/K81z15CGfnNMqRO0onPB-Q/o.jpg"
	  ],
	  "distance": 298.9865852529836,
	  "r_id": 5
	},
	{
	  "id": "xVoy_j2M7KulKHs3NZLqLg",
	  "name": "Subway",
	  "coordinates": {
		"latitude": 40.70933,
		"longitude": -74.00903
	  },
	  "photos": [
		"https://s3-media3.fl.yelpcdn.com/bphoto/XBrYo0HowL0frL9zaRbJDA/o.jpg"
	  ],
	  "distance": 459.7621672468454,
	  "r_id": 5
	},
	{
	  "id": "apD1JSpVX_q4PJh90xJPoQ",
	  "name": "McDonald's",
	  "coordinates": {
		"latitude": 40.715771,
		"longitude": -74.005234
	  },
	  "photos": [
		"https://s3-media3.fl.yelpcdn.com/bphoto/cQ6rCEKzwfXryupYe5VS8g/o.jpg"
	  ],
	  "distance": 336.6089983009257,
	  "r_id": 20
	},
	{
	  "id": "hT229QtA--o0F591IUE7SQ",
	  "name": "McDonald's",
	  "coordinates": {
		"latitude": 40.708279,
		"longitude": -74.004868
	  },
	  "photos": [
		"https://s3-media3.fl.yelpcdn.com/bphoto/juusF-LISHxoy5Bgz1NqiQ/o.jpg"
	  ],
	  "distance": 511.686166983117,
	  "r_id": 20
	},
	{
	  "id": "in7QM4vNUTHVZIi7Outetg",
	  "name": "McDonald's",
	  "coordinates": {
		"latitude": 40.7163213865248,
		"longitude": -74.0106988300204
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/U_gBqmHkjAVm0W2C0XBtww/o.jpg"
	  ],
	  "distance": 556.9173980316313,
	  "r_id": 20
	},
	{
	  "id": "YV0VaJgV2zotzI7XNUaMCQ",
	  "name": "McDonald's",
	  "coordinates": {
		"latitude": 40.71852,
		"longitude": -74.00115
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/dHGOx2gmQmIupGDEMzxRGg/o.jpg"
	  ],
	  "distance": 754.3464336945959,
	  "r_id": 20
	},
	{
	  "id": "2JAoFi1nu3xlrdNzW5FLMQ",
	  "name": "Quiznos",
	  "coordinates": {
		"latitude": 40.7143517,
		"longitude": -74.0070343
	  },
	  "photos": [
		"https://s3-media3.fl.yelpcdn.com/bphoto/None/o.jpg"
	  ],
	  "distance": 200.34185572099815,
	  "r_id": 26
	},
	{
	  "id": "pTv4RDll3N_CGG3BwtEctg",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.711619,
		"longitude": -74.006730001524
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/cNhCmOGY46lvQpUM3WkTrA/o.jpg"
	  ],
	  "distance": 145.02052018349193,
	  "r_id": 28
	},
	{
	  "id": "4jWEL2gMhfwfzej0rfgJiQ",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.712214,
		"longitude": -74.008202
	  },
	  "photos": [
		"https://s3-media4.fl.yelpcdn.com/bphoto/aL3dgTEKMtWSukJvSf2wkw/o.jpg"
	  ],
	  "distance": 196.70122521109678,
	  "r_id": 28
	},
	{
	  "id": "V05D4jAQb3rT3WqEBGq-og",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.714843,
		"longitude": -74.005956
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/1xtC1Mub5-M3Lb83X8nK2Q/o.jpg"
	  ],
	  "distance": 227.20125054409243,
	  "r_id": 28
	},
	{
	  "id": "8uTTEIDkYuEXywiWcP5w-w",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.713782,
		"longitude": -74.009077
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/k0xntGVCiy-TcOWdB-jpeA/o.jpg"
	  ],
	  "distance": 281.3912274165507,
	  "r_id": 28
	},
	{
	  "id": "12xF41OEVDoqsSwnDTEH1g",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.710792,
		"longitude": -74.008533
	  },
	  "photos": [
		"https://s3-media4.fl.yelpcdn.com/bphoto/x6evzCXlkmQrh-5CExjMkw/o.jpg"
	  ],
	  "distance": 308.92386740902833,
	  "r_id": 28
	},
	{
	  "id": "j6nTfbYavGVsJKxz6bte6w",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.710194,
		"longitude": -74.007894
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/M1t60niwqcQr2rwz73lC0w/o.jpg"
	  ],
	  "distance": 330.8368041827513,
	  "r_id": 28
	},
	{
	  "id": "rveP6oXixE2a2sKwMqS-ww",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.715532,
		"longitude": -74.009003
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/eilie99DxiF1qrFEhl_SEQ/o.jpg"
	  ],
	  "distance": 395.4047083287184,
	  "r_id": 28
	},
	{
	  "id": "XzAZVxmGPcysK8hYdcAJKA",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.715754,
		"longitude": -74.003125
	  },
	  "photos": [
		"https://s3-media1.fl.yelpcdn.com/bphoto/wnmygBYHBjPyh66exssYIQ/o.jpg"
	  ],
	  "distance": 408.1757142395983,
	  "r_id": 28
	},
	{
	  "id": "hh2T5aLJsSjpHvPRcQxIJw",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.712117,
		"longitude": -74.01087
	  },
	  "photos": [
		"https://s3-media1.fl.yelpcdn.com/bphoto/I-vhvMK0Iu-227TUUDg82Q/o.jpg"
	  ],
	  "distance": 417.43389864130796,
	  "r_id": 28
	},
	{
	  "id": "CszSAlKolUeJ0-IaRDL0YQ",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.71091,
		"longitude": -74.010363
	  },
	  "photos": [
		"https://s3-media1.fl.yelpcdn.com/bphoto/X_JRUYoLm8Sayz8FLTzBaA/o.jpg"
	  ],
	  "distance": 423.55369157205456,
	  "r_id": 28
	},
	{
	  "id": "60agfQbky4cX8BEApyltIA",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.711018,
		"longitude": -74.000916
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/f6QjfLPOkJ981z8V5SFnpw/o.jpg"
	  ],
	  "distance": 472.10364723740355,
	  "r_id": 28
	},
	{
	  "id": "OuNYQaqEJjkBHRdnC-4HOw",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.708413,
		"longitude": -74.007387
	  },
	  "photos": [
		"https://s3-media1.fl.yelpcdn.com/bphoto/nnOCqsDkhf3zv5BCLWQi1A/o.jpg"
	  ],
	  "distance": 501.62460524372347,
	  "r_id": 28
	},
	{
	  "id": "p8Fygzdm1uONUTu4q-SfJA",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.709590668884,
		"longitude": -74.010550737875
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/6QYI1E8yqtvguCxar6gj8A/o.jpg"
	  ],
	  "distance": 523.9009269439965,
	  "r_id": 28
	},
	{
	  "id": "9SA6MazvA6lXqTpM-ZlYIA",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.708771,
		"longitude": -74.009313
	  },
	  "photos": [
		"https://s3-media1.fl.yelpcdn.com/bphoto/2QL7l8gm9QcXfhOoOhaj2g/o.jpg"
	  ],
	  "distance": 527.905220857281,
	  "r_id": 28
	},
	{
	  "id": "OQEU7KPvgDX9kX2xlHGm5g",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.710372,
		"longitude": -74.011793
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/-Ewagmynd2HIvyhmryyNcg/o.jpg"
	  ],
	  "distance": 557.9388652341933,
	  "r_id": 28
	},
	{
	  "id": "wE3_AwZoZ0ZXbngBH1xRrg",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.718247,
		"longitude": -74.007305
	  },
	  "photos": [
		"https://s3-media3.fl.yelpcdn.com/bphoto/uIr59_YV9r_yVhknYKBzUA/o.jpg"
	  ],
	  "distance": 615.5834744176956,
	  "r_id": 28
	},
	{
	  "id": "FrigBgxfxMxwPJNq-hAr0Q",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.708758,
		"longitude": -74.011055
	  },
	  "photos": [
		"https://s3-media3.fl.yelpcdn.com/bphoto/lJ6EWit_xdIzuTgLBk03gA/o.jpg"
	  ],
	  "distance": 619.3067748020209,
	  "r_id": 28
	},
	{
	  "id": "VGMdcAn34GeasB4arA8XUA",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.707122,
		"longitude": -74.004971
	  },
	  "photos": [
		"https://s3-media3.fl.yelpcdn.com/bphoto/5ZlR8q_l68ROd4fqVqZwog/o.jpg"
	  ],
	  "distance": 637.293588160514,
	  "r_id": 28
	},
	{
	  "id": "_H9AnQkr5VtOiykNnxcQlQ",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.706487,
		"longitude": -74.006685
	  },
	  "photos": [
		"https://s3-media3.fl.yelpcdn.com/bphoto/Kuf3P75m7I0NP3XzcaT8VQ/o.jpg"
	  ],
	  "distance": 704.3432594824154,
	  "r_id": 28
	},
	{
	  "id": "_rgXsb2X_NVwbt34qyHa4A",
	  "name": "Starbucks",
	  "coordinates": {
		"latitude": 40.707437,
		"longitude": -74.010976
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/uKcICK805S97IQpnthaegg/o.jpg"
	  ],
	  "distance": 729.0600197852675,
	  "r_id": 28
	},
	{
	  "id": "uL9Pz4XTLmv9Rial9BKggw",
	  "name": "Subway",
	  "coordinates": {
		"latitude": 40.71012,
		"longitude": -74.0054399
	  },
	  "photos": [
		"https://s3-media3.fl.yelpcdn.com/bphoto/SRqV-Xr4vF2IpbeJvglHYw/o.jpg"
	  ],
	  "distance": 297.8143812585539,
	  "r_id": 29
	},
	{
	  "id": "6XAJEUfXDkWCBxAJTQSZxQ",
	  "name": "Subway",
	  "coordinates": {
		"latitude": 40.7164726406336,
		"longitude": -74.0048296004534
	  },
	  "photos": [
		"https://s3-media3.fl.yelpcdn.com/bphoto/None/o.jpg"
	  ],
	  "distance": 433.13137003457285,
	  "r_id": 29
	},
	{
	  "id": "JYZIJRhWc6KA-SWuFJPAXw",
	  "name": "Subway",
	  "coordinates": {
		"latitude": 40.70871,
		"longitude": -74.01164
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/LNIiBKiwF0W_wP5DUeVXYA/o.jpg"
	  ],
	  "distance": 654.2662587754787,
	  "r_id": 29
	},
	{
	  "id": "KMt3QFk-Bd5FMxxMW2mpnQ",
	  "name": "Subway",
	  "coordinates": {
		"latitude": 40.719177,
		"longitude": -74.004311
	  },
	  "photos": [
		"https://s3-media1.fl.yelpcdn.com/bphoto/Fydkht8NaqZR9Gt2IhR7dg/o.jpg"
	  ],
	  "distance": 723.2862613474417,
	  "r_id": 29
	},
	{
	  "id": "-3Gj4IrIu5Dn0LIZxidwkA",
	  "name": "Subway",
	  "coordinates": {
		"latitude": 40.7181881373376,
		"longitude": -74.000615833625
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/QQIxYB3P43IXAugDw1ETwg/o.jpg"
	  ],
	  "distance": 751.58437834877,
	  "r_id": 29
	},
	{
	  "id": "6eId3eiOMQf5iEhypNBV0g",
	  "name": "Subway",
	  "coordinates": {
		"latitude": 40.70854,
		"longitude": -74.01367
	  },
	  "photos": [
		"https://s3-media3.fl.yelpcdn.com/bphoto/None/o.jpg"
	  ],
	  "distance": 798.1435822448293,
	  "r_id": 29
	},
	{
	  "id": "RGK1ASygW8Q4JeJJ-Az9Xw",
	  "name": "Subway",
	  "coordinates": {
		"latitude": 40.7062784,
		"longitude": -74.011724
	  },
	  "photos": [
		"https://s3-media3.fl.yelpcdn.com/bphoto/W7H0GugJX_GeGRuN_Gndnw/o.jpg"
	  ],
	  "distance": 871.0014781972438,
	  "r_id": 29
	},
	{
	  "id": "XfugTKWi1sI_5Ai1Dr8t9A",
	  "name": "Subway",
	  "coordinates": {
		"latitude": 40.70634,
		"longitude": -74.0038299
	  },
	  "photos": [
		"https://s3-media3.fl.yelpcdn.com/bphoto/None/o.jpg"
	  ],
	  "distance": 879.6994332025077,
	  "r_id": 29
	},
	{
	  "id": "OV66AsjGZZMvugzrxu4aig",
	  "name": "Subway",
	  "coordinates": {
		"latitude": 40.70458,
		"longitude": -74.01048
	  },
	  "photos": [
		"https://s3-media3.fl.yelpcdn.com/bphoto/Iig_LbXdc0lrSM0fQqSN2Q/o.jpg"
	  ],
	  "distance": 991.4172018153475,
	  "r_id": 29
	},
	{
	  "id": "z3n_hpAFLyDHfMrJKV6fkA",
	  "name": "Subway",
	  "coordinates": {
		"latitude": 40.7052841186523,
		"longitude": -74.0142669677734
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/TytLd4qUIKW8a_77vQD24w/o.jpg"
	  ],
	  "distance": 1080.648051847869,
	  "r_id": 29
	},
	{
	  "id": "WdfaJmWIiK80aQd2o2RR-Q",
	  "name": "Subway",
	  "coordinates": {
		"latitude": 40.703255,
		"longitude": -74.01107
	  },
	  "photos": [
		"https://s3-media4.fl.yelpcdn.com/bphoto/eotbiv9b7Pk0sf-jip_e4A/o.jpg"
	  ],
	  "distance": 1144.1888742186607,
	  "r_id": 29
	},
	{
	  "id": "IWZqUl46ogITQyDz3UAXPQ",
	  "name": "Subway",
	  "coordinates": {
		"latitude": 40.7023139,
		"longitude": -74.0114968
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/DajepoLqEE50EsInb7BPFA/o.jpg"
	  ],
	  "distance": 1290.5257346612925,
	  "r_id": 29
	},
	{
	  "id": "FpYVG2jW7SmLhsHBNEGDtg",
	  "name": "Wendy's",
	  "coordinates": {
		"latitude": 40.701009,
		"longitude": -74.0130246
	  },
	  "photos": [
		"https://s3-media2.fl.yelpcdn.com/bphoto/sjZoHjq9YFlR4L-Uz1ekrA/o.jpg"
	  ],
	  "distance": 1438.6034317917224,
	  "r_id": 32
	}
  ]
`)
