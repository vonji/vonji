import R from 'ramda';
import knex from '../utils/db';
import { fetchOne, save, exists } from '../utils/dbHelpers';

const _tagsForRequest = async request_id =>
  await knex('requests_tags')
    .join('tags', 'requests_tags.tag_id', 'tags.id')
    .where({ request_id })
		.select('id', 'name');

/*
const _pluckUser = async promise => {
  const object = await promise;
  if ('user_id' in object) {
    const user = await knex('users')
      .where({ id: object.user_id }).first()
      .select('id', 'displayed_name');
    return { ..._.omit(object, 'user_id'), user };
  }
  throw new Error(`Can not pluck user from object ${object}`);
};
*/

const bindTag = async (tag, tag_id, request_id) => {
	await save('tags', tag, tag_id);
	const bound = await knex('requests_tags').where({ tag_id, request_id });
	if (!bound) {
		await knex('requests_tags').insert({ tag_id, request_id });
	}
}

export const fetchRequest = async request_id => {
	const request = await fetchOne('requests', request_id);
	const tags = await _tagsForRequest(request_id);
	return { ...request, tags };
};

export const createRequest = async request => {
	await Promise.all(R.map(tag => bindTag(tag, tag.id, request.id), request.tags));
};
