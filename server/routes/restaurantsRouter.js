const router = require('express').Router();
const db = require('../database/dbConfig');

router.get('/names', async (req, res) => {
    const names = await db('restaurants').column('name');
    console.log('names', names);
    res.status(200).json(names);
});

router.get('/items', async (req, res) => {
    const { name } = req.query;

    const restaurant = await db('restaurants').where('name', '=', name).first();
    const items = await db('items').where('restaurant_id', '=', restaurant.id);

    res.status(200).json(items)
})

module.exports = router;
