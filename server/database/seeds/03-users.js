const bcrypt = require('bcryptjs');

exports.seed = function(knex, Promise) {
  return knex('users').insert([
    {
      id: 0, 
      email: 'rey@rey.sh',
      password: bcrypt.hashSync('pass', 8),
      name: 'Rey'
    },
  ]);
};