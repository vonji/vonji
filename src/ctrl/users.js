import express from 'express';
import piper from '../utils/piper';
import R from 'ramda';
import { createUser, fetchAllUserId, fetchOneUser } from '../utils/api';

const router = express.Router();

router.post('/', piper(async (req, res) => {
	const newUser = await createUser(req.body);
	console.log(newUser);
	req.reply.created();
}));

router.get('/', piper(async (req, res) => {
	const idList = await fetchAllUserId();
	const userify = R.map(userId => fetchOneUser(userId));
	const users = await Promise.all(userify(idList));
	req.reply.ok(users);
}));

export default router;
