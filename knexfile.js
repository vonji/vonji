// Update with your config settings.

module.exports = {

	development: {
		client: 'postgresql',
		connection: {
			database: 'vonji-api-dev',
			user: 'vonji-api',
			password: 'vonji-api'
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
			database: 'vonji_api-staging',
			user: 'vonji-api',
			password: 'vonji-api'
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
			user: 'vonji-api',
			password: 'vonji-api'
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
