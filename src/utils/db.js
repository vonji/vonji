import knexFactory from "knex";
import dbConfig from '../../knexfile';

const knex = knexFactory({
	client: 'pg',
	connection: dbConfig.development.connection,
});

export default knex;
