import knex from './db';
import R from 'ramda';

export const exists = async (resource, id) => {
	return await knex(resource).where({ id }).first();
}

export const fetchOne = async (resource, id) => {
	const result = await knex(resource).where({ id }).first();
	if (!result) {
		throw { name: 'NotFoundError', message: `${resource} with id ${id} does not exist.` };
	}
	return result;
};

export const fetchAll = async (resource, id) => {
	return await knex(resource).select();
};

export const save = async (resource, attributes, id) => {
	if (id) {
		await fetchOne(resource, id);
		await knex(resource)
			.where({ id })
			.update(attributes);
	} else {
		await knex(resource)
			.insert(attributes);
	}
}

export const destroy = async (resource, id) => {
	await fetchOne(resource, id);
	return await knex(resource)
		.where({ id })
		.delete();
}
