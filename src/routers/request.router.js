import Request from "../models/request.model";
import express from "express";
import _ from "lodash";
import {fetchOne, fetchAll, remove, save} from "./utils";
import {knex} from "../bookshelf";

const router = express.Router();

router.get('/:id', (req, res) => {
	// fetchOne(res, new Request({id: req.params.id}));

	knex
		.from('requests')
		.where({id: req.params.id})
		.first().then(data => {
		data.tags = [];
		return knex('tags_requests')
			.innerJoin('tags', 'tags.id', 'tags_requests.tag_id')
			.where({'request_id': data.id})
			.select('name').then(tags => {
				_.forEach(tags, tag => {
					data.tags.push(tag.name);
				});
				return data;
			})
			.then(data => {
				res.send(data);
			});
	});
});

router.get('/', (req, res) => {
	fetchAll(res, Request);
});

router.post('/', (req, res) => {
	save(res, req.body, new Request());
});

router.put('/:id', (req, res) => {
	fetchOne(res, new Request({id: req.params.id}), {
		done(model) {
			save(res, req.body, model);
		},
	});
});

router.delete('/:id', (req, res) => {
	fetchOne(res, new Request({id: req.params.id}), {
		done(model) {
			remove(res, model);
		},
	});
});

export default router;
