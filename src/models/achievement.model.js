import {bs} from "../bookshelf";

const Achievement = bs.Model.extend({
    tableName: 'achievements',
    hasTimestamps: true,
});

export default Achievement;