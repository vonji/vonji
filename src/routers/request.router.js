import Request from "../models/request.model";
import express from "express";
import _ from "lodash";
import {simpleRouting} from '../utils/simpleRouting';
import {piper} from '../utils/piper';
import {knex} from "../bookshelf";

const router = express.Router();

const tagsForRequest = async request_id =>
  await knex('requests_tags')
    .join('tags', 'requests_tags.tag_id', 'tags.id')
    .where({ request_id })
    .pluck('name');

const pluckUser = async promise => {
  const object = await promise;
  if ('user_id' in object) {
    const user = await knex('users')
      .where({ id: object.user_id }).first()
      .select('id', 'displayed_name');
    return { ..._.omit(object, 'user_id'), user };
  }
  throw new Error(`Can not pluck user from object ${object}`);
}

const fetchRequest = async id => {
	const request = await knex('requests').where({ id }).first();
	if (!request) {
		throw { name: 'NotFoundError', message: `Request with id ${id} does not exist.` };
	}
};

router.get('/:id', piper(async (req, res) => {
	const user = await pluckUser(fetchRequest(req.params.id));
	const tags = await tagsForRequest(req.params.id);
	return req.reply.ok({ ...user, tags });
}));

export default router;
