
exports.seed = function(knex, Promise) {
  // Deletes ALL existing entries
  return knex('restaurants').insert([
    {id: 1, name: 'testName'},
  ]);
};
