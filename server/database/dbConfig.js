const dbEngine = process.env.DB_VERSION || 'development'
const knex = require('knex');
const knexConfig = require('../knexfile.js');

module.exports = knex(knexConfig[dbEngine]);
