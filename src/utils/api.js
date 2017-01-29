import r from './redis';

export const createUser = async (data) => {
	const userId = await r.incrAsync('next_user_id');
	const userKey = `user:${userId}`;
	await r.multi()
		.sadd('users', userId)
		.hmset(userKey, data)
		.execAsync();
	return await r.hgetallAsync(userKey);
};

export const fetchOneUser = async id => {
	return await r.hgetallAsync(`user:${id}`);
};

export const fetchAllUserId = async () => {
	return await r.smembersAsync('users');
}
