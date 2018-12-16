
exports.up = function(knex, Promise) {
    return knex.schema.createTable('items', table => {
        table.increments();
        table.string('name', 128).notNullable().unique();
        table.string('type', 128);
        table.integer('protein').notNullable();
        table.integer('carbs').notNullable();
        table.integer('fats').notNullable();
        table.integer('calories').notNullable();
        table.decimal('calsperpro');
        table.integer('sodium');
        table.integer('restaurant_id')
          .unsigned()
          .references('id')
          .inTable('restaurants')
          .notNullable();
  
    })
  };
  
  exports.down = function(knex, Promise) {
    return knex.schema.dropTableIfExist('items');
  };
  