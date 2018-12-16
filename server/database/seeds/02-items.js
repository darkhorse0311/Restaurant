
exports.seed = function(knex, Promise) {
  // Deletes ALL existing entries
  return knex('items').insert([
    {
      id: 1, 
      name: 'name',
      type: 'type',
      protein: 10,
      carbs: 20,
      fats: 5,
      calories: 125,
      calsperpro: 1.25,
      sodium: 200,
      restaurant_id: 1,
    },
  ]);
};
