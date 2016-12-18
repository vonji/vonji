import {bs} from "../bookshelf";
import User from "./user.model";

const Notification = bs.Model.extend({
	tableName: 'notifications',
	hasTimestamps: true,
	user: () => {
		return this.belongsTo(User)
	},
});

export default Notification;
