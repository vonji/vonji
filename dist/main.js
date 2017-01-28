"use strict";

var _express = require("express");

var _express2 = _interopRequireDefault(_express);

var _bodyParser = require("body-parser");

var _bodyParser2 = _interopRequireDefault(_bodyParser);

var _user = require("./routers/user.router");

var _user2 = _interopRequireDefault(_user);

var _ad = require("./routers/ad.router");

var _ad2 = _interopRequireDefault(_ad);

var _achievement = require("./routers/achievement.router");

var _achievement2 = _interopRequireDefault(_achievement);

var _request = require("./routers/request.router");

var _request2 = _interopRequireDefault(_request);

var _tag = require("./routers/tag.router");

var _tag2 = _interopRequireDefault(_tag);

var _response = require("./routers/response.router");

var _response2 = _interopRequireDefault(_response);

var _notification = require("./routers/notification.router");

var _notification2 = _interopRequireDefault(_notification);

var _transaction = require("./routers/transaction.router");

var _transaction2 = _interopRequireDefault(_transaction);

var _comment = require("./routers/comment.router");

var _comment2 = _interopRequireDefault(_comment);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

var app = (0, _express2.default)();

app.use(_bodyParser2.default.json());
app.use(_bodyParser2.default.urlencoded({ extended: true }));

app.use('/users', _user2.default);
app.use('/achievements', _achievement2.default);
app.use('/ads', _ad2.default);
app.use('/tags', _tag2.default);
app.use('/requests', _request2.default);
app.use('/notifications', _notification2.default);
app.use('/responses', _response2.default);
app.use('/transactions', _transaction2.default);
app.use('/comments', _comment2.default);

app.listen(3000, function () {
	console.log('Application started');
});