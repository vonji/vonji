import {bs} from "../bookshelf";

const Response = bs.Model.extend({
	tableName: 'responses',
	hasTimestamps: true,
});

export default Response;