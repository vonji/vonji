// Update with your config settings.

module.exports = {
	development: {
		client: 'postgresql',
		connection: {
			host: 'localhost',
			port: 5432,
			database: 'vonji_api_dev',
			user: 'vonji_api',
			password: 'vonji_api'
		},
		pool: {
			min: 2,
			max: 10
		},
		migrations: {
			tableName: 'knex_migrations'
		}
	},
};
