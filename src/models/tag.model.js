import {bs} from "../bookshelf";

const Tag = bs.Model.extend({
	tableName: 'tags',
	hasTimestamps: true,
});

export default Tag;
