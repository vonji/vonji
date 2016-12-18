import {bs} from "../bookshelf";

const Ad = bs.Model.extend({
	tableName: 'ads',
	hasTimestamps: true,
});

export default Ad;