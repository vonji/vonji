import knexFactory from "knex";
import bookshelfFactory from "bookshelf";

const knex = knexFactory({
	client: 'pg',
	connection: {
		host: '127.0.0.1',
		user: 'vonji_api',
		password: 'vonji_api',
		database: 'vonji_api_dev',
		charset: 'utf8'
	}
});

const bs = bookshelfFactory(knex);

export { bs, knex };
