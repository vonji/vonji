import express from "express";
import piper from '../utils/piper';
import * as Req from '../services/request';

const router = express.Router();

router.get('/:id', piper(async (req, res) => {
	return req.reply.ok(Req.fetchRequest(req.params.id));
}));

router.post('/', piper(async (req, res) => {
	return req.reply.created(Req.createRequest(req.body));
}));

export default router;
