const router = require('express').Router();
const db = require('../database/dbConfig');

router.get('/names', async (req, res) => {

    const names = await db('restaurants').column('name');
    console.log('names', names);
    res.status(200).json(names);
});

module.exports = router;
