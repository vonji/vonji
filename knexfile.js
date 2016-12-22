// Update with your config settings.

module.exports = {

	development: {
		client: 'postgresql',
		connection: {
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

	staging: {
		client: 'postgresql',
		connection: {
			database: 'vonji_api_staging',
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

	production: {
		client: 'postgresql',
		connection: {
			database: 'vonji_api',
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
	}

};
