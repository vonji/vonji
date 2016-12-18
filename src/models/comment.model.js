import {bs} from "../bookshelf";

const Comment = bs.Model.extend({
	tableName: 'comments',
	hasTimestamps: true,
});

export default Comment;
