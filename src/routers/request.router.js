import Request from "../models/request.model";
import express from "express";
import _ from "lodash";
import {simpleRouting} from "./utils";
import {knex} from "../bookshelf";

const router = express.Router();

const tagsForRequest = async request_id =>
  await knex('requests_tags')
    .join('tags', 'requests_tags.tag_id', 'tags.id')
    .where({ request_id })
    .pluck('name');

const pluckUser = async object => {
  if ('user_id' in object) {
    const user = await knex('users')
      .where({ id: object.user_id }).first()
      .select('id', 'displayed_name');
    return { ..._.omit(object, 'user_id'), user };
  }
  throw `Can not pluck user from object ${object}`;
}

const getRequest = async id =>
  await knex('requests')
    .where({ id })
    .first();

router.get('/:id', async (req, res) => {
  try {
    res.json({
      ...(await pluckUser(await getRequest(req.params.id))),
      tags: await tagsForRequest(req.params.id)
    });
  } catch (e) {
    console.error(e);
  }
});

export default router;
