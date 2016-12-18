import {bs} from "../bookshelf";
import User from "./user.model";

const Request = bs.Model.extend({
	tableName: 'requests',
	hasTimestamps: true,
});

export default Request;
