
exports.up = function(knex, Promise) {
  return knex.schema.createTable('restaurants', table => {
      table.increments();
      table.string('name', 128).notNullable().unique();
  })
};

exports.down = function(knex, Promise) {
  return knex.schema.dropTableIfExist('restaurants');
};
