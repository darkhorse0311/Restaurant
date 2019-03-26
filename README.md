# Carbtographer
Application that recommends macro friendly meals from local restaurants 

###### [Backend Deployment Link](https://carbtographer.herokuapp.com/)

##### NOTICE
Recently ported Node.js backend to a golang backend.
Old Node.js back end on `node-backend` branch [here](https://github.com/reynld/carbtographer/tree/node-backend).

## Backend

### Enviroment Variables
`YELP_API_KEY`: Yelp GraphQL API Key
`AWS_HOST`: AWS DB Host URL
`AWS_PORT`: AWS DB Port
`AWS_USER`: AWS DB User
`AWS_DBNAME`: AWS DBNAME
`AWS_PASSWORD`: AWS Password

### Deployment

#### Setup Repo
1. Clone repo outside of `$GOPATH`
`git clone https://github.com/reynld/carbtographer`

2. Initialize go modules
`go mod init`

3. Make sure that any unused modules have been removed from your application:
`go mod tidy`

4. Vendoring dependencies:
`go mod vendor`

#### Setup Heroku
1. Install Heroku CLI
`brew tap heroku/brew && brew install heroku`

2. Login to Heroku
`heroku login`

3. Create App on heroku
`heroku create`

4. Push code up
`git push heroku master`

5. Open app
`heroku open`


### ENDPOINTS

#### GET `/`

- Default endpoint

Response:
```
{
    status: 200,
    message: "server live on port: ####"
}
```

#### GET `/names`

- Returns all restaurant names in database

Response:
```
[
    {
        id: 1
        name: "Name"
    }
]
```

#### GET `/locations/{lat}/{lon}`

- Returns returant info based on lat and long

##### PARAMETERS:  
`lat`: latitude  
`lon`: longitude  

Response:
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

#### GET `/items/{id}`

- Returns all items per restuarant

URL Parameters:  
`id`: Restaurant ID  

Response:
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

#### GET `/*`

- 404 Endpoint

Response:
```
{
    status: 404,
    message: "route not found"
}
```