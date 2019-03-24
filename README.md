# Carbtographer
Application that recommends macro friendly meals from local restaurants 


## Server Side

### Enviroment Variables

`YELP_CLIENT_ID`: Yelp GraphQL API Client ID

`YELP_API_KEY`: Yelp GraphQL API Key

`JWT_SECRET`: JWT Secret for genrating JWT

## ENDPOINTS

##### GET `/`

- Default endpoint

Response:
```
{
    status: 200,
    message: "server live on port: ####"
}
```

##### GET `/names`

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

##### GET `/items/{id}`

- Returns all items per restuarant

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

##### GET `/*`

- 404 Endpoint

Response:
```
{
    status: 404,
    message: "route not found"
}
```