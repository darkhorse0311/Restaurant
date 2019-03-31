# Carbtographer
Application that recommends macro friendly meals from local restaurants.

[![Netlify Status](https://api.netlify.com/api/v1/badges/a49fb016-d1b4-4fdb-af1e-9f4e371b2252/deploy-status)](https://app.netlify.com/sites/carbtographer/deploys)
 
[![Build Status](https://travis-ci.com/reynld/carbtographer.svg?branch=master)](https://travis-ci.com/reynld/carbtographer)
[![CodeFactor](https://www.codefactor.io/repository/github/reynld/carbtographer/badge/master)](https://www.codefactor.io/repository/github/reynld/carbtographer/overview/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/reynld/carbtographer)](https://goreportcard.com/report/github.com/reynld/carbtographer)
[![Coverage Status](https://coveralls.io/repos/github/reynld/carbtographer/badge.svg?branch=coverage-setup)](https://coveralls.io/github/reynld/carbtographer?branch=coverage-setup)

Frontend Link: https://carbtographer.com

Backend Link: https://carbtographer.herokuapp.com

#### NOTICE
Recently ported Node.js backend to a Golang backend.
 
Node.js back end on `node-backend` branch [here](https://github.com/reynld/carbtographer/tree/node-backend).

# Frontend
Frontend built in React.js, Redux and Styled Components

## Enviroment Variables
You will require a Mapbox Access Token to make authenticated requests.
 
You can generate one following the intructions here: [Mapbox: Access tokens](https://docs.mapbox.com/help/how-mapbox-works/access-tokens/)

- ``REACT_APP_MAP_BOX_KEY`` - your Mapbox Access Token (API Key)
- ``REACT_APP_BACKEND_URL`` - your backend link

## Usage
 With all enviroment variables in place you can now run it locally

#### CD into client
`cd client`
 
#### Install node modules
 `npm i`
 
#### Run
 `npm run start`


# Backend

## Enviroment Variables
You will require a Yelp API Key to make authenticated requests.
 
You can generate one following the intructions here: [Creating an app on Yelp's Developers site](https://www.yelp.com/developers/documentation/v3/authentication)
- ``YELP_API_KEY`` - your Yelp Key (API Key)
- ``AWS_HOST`` - your database Host URL
- ``AWS_PORT``- your database PORT
- ``AWS_USER`` - your database user
- ``AWS_DBNAME`` - your database name
- ``AWS_PASSWORD`` - your database password

## Usage
 With all enviroment variables in place you can now run the api locally

#### Install
`go install`
 
#### Build
 `go build`
 
#### Run
 `./carbtographer`

<br/>

## ENDPOINTS

#### GET `/`
 
```
{
    status: 200,
    message: "server live on port: ####"
}
```
<br/>

#### GET `/names`

Returns all restaurant names in database
 
```
[
    {
        id: 1
        name: "Name"
    }
]
```
<br/>

#### GET `/locations/{lat}/{lon}`
 
Returns restaurant info based on lat and long
 
`lat`: latitude  
`lon`: longitude  
 
```
[
  {
    "id": "id",
    "name": "Name",
    "coordinates": {
      "latitude": 0.0,
      "longitude": 0.0
    },
    "photos": [
      "url"
    ],
    "distance": 20.0
  }
]
```
<br/>

#### GET `/items/{id}`

Returns all items for any given restaurant ID
 
`id`: restaurant ID  
 
```
[
    {
        "id": 0,
        "name": "item name",
        "type": "type",
        "protein": 0,
        "carbs": 0,
        "fats": 0,
        "calories": 0,
        "calperpro": 0.0
        "sodium": 0,
        "r_id": 0
    }
]
```
<br/>

#### GET `/*`
 
404 Endpoint
 
```
{
    status: 404,
    message: "route not found"
}
```
