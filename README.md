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

##### GET `/*`

- 404 Endpoint

Response:
```
{
    status: 404,
    message: "route not found"
}
```