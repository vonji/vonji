import {bs} from "../bookshelf";

const Transaction = bs.Model.extend({
	tableName: 'transactions',
	hasTimestamps: true,
});

export default Transaction;
