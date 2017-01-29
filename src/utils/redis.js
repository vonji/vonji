import redis from 'redis';
import Promise from 'bluebird';
Promise.promisifyAll(redis.RedisClient.prototype);
Promise.promisifyAll(redis.Multi.prototype);

const r = redis.createClient({
	host: '192.168.99.100',
	port: 32768,
});

export default r;
