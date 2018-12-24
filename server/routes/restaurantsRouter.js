require('dotenv').config();

const router = require('express').Router();
const { request, GraphQLClient } =  require('graphql-request')
const db = require('../database/dbConfig');

const yelpApi = "https://api.yelp.com/v3/graphql"

const client = new GraphQLClient(yelpApi, {
    headers: {
      Authorization: `Bearer ${process.env.YELP_API_KEY}`,
    },
})

router.get('/names', async (req, res) => {
    const names = await db('restaurants').column('name');
    // console.log('names', names);
    res.status(200).json(names);
});

router.get('/items', async (req, res) => {
    const nameSplit = req.query.name.split(' ');

    const name = nameSplit.length > 2 ? nameSplit[0] : nameSplit.join('');

    const restaurant = await db('restaurants').where('name', '=', name).first();
    // console.log('restaurant', restaurant);
    const items = await db('items').where('restaurant_id', '=', restaurant.id);
    // console.log('items', items);

    res.status(200).json(items)
})

router.get('/locations', async (req, res) => {
    const namesRes = await db('restaurants').column('name');

    const names = namesRes.map(nameObj => nameObj.name);

    let chainedQuery = "{";
    names.forEach(name => {
        const reg = /-|\.|'|\//gi;
        let n = name.replace(reg, "");
        const query = `
            ${n}: search(
                term:"${name}",
                longitude: -74.0060,
                latitude: 40.7128,
                radius: 500
            ) {
                ...businessInfo
            }
        `
        chainedQuery = chainedQuery + query;
    })
    chainedQuery = chainedQuery + `
    }
    fragment businessInfo on Businesses{
        total
            business {
                id
                name
                coordinates {
                    latitude
                    longitude
                }
                photos
                distance
        }
    }
    `;

    const yelpRes = await client.request(chainedQuery);

    const resArray = Object.values(yelpRes);

    const ids = [];
    const uniqueBusinesses = [];

    resArray.forEach(busArr => {
        busArr.business.forEach(bus => {
            const id = bus.id
            if(!ids.includes(id)) {
                ids.push(id);
                uniqueBusinesses.push(bus)
            }
        })
    })

    const validBusinesess = [];

    uniqueBusinesses.forEach(bus => {
        if(names.includes(bus.name)){
            validBusinesess.push(bus)
        }
    })

    res.status(200).json(validBusinesess)
})

module.exports = router;
