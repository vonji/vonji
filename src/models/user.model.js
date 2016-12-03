import {bs} from "../bookshelf";

const User = bs.Model.extend({
    tableName: 'users',
    hasTimestamps: true,
});

export default User;