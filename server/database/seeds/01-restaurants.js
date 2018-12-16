const restaurantData = require('../restaurantData.json');

const resNames = restaurantData.map((res, index) => ({
  id: index,
  name: Object.keys(res)[0]
}));


exports.seed = function(knex, Promise) {
  // Deletes ALL existing entries
  return knex('restaurants').insert(resNames);
};
