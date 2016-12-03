import {bs} from "./bookshelf";

export const User = bs.Model.extend({
    tableName: 'users',
    hasTimestamps: true,
});