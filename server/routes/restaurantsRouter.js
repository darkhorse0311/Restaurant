const router = require('express').Router();
const db = require('../database/dbConfig');

router.get('/names', async (req, res) => {
    const names = await db('restaurants').column('name');
    console.log('names', names);
    res.status(200).json(names);
});

router.get('/items', async (req, res) => {
    const nameSplit = req.query.name.split(' ');

    const name = nameSplit.length > 2 ? nameSplit[0] : nameSplit.join('');

    const restaurant = await db('restaurants').where('name', '=', name).first();
    console.log('restaurant', restaurant);
    const items = await db('items').where('restaurant_id', '=', restaurant.id);
    console.log('items', items);

    res.status(200).json(items)
})

module.exports = router;
