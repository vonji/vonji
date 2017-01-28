import test from 'ava';
import * as base from './base';

const userModel = {
	"DeletedAt": null,
	"Email": "gordonfreeman@free.fr",
	"Password": "mesa",
	"DisplayedName": "Gordon Freeman",
	"RealName": "Gordon Freeman",
	"Description": "",
	"Motto": "...",
	"FacebookLink": "",
	"TwitterLink": "",
	"LinkedInLink": "",
	"Phone": "",
	"Birthday": "",
	"Location": "City 17",
	"VCoins": 179,
	"VActions": 15,
	"Avatar": "https://i.imgur.com/KDiqEoj.jpg",
	"Gender": "man",
	"Achievements": [],
	"Tags": []
};

const userRequest = base.request.defaults({
	baseUrl: base.request.baseUrl + '/users'
});

test('getUser', async t => {
	base.get(t, userRequest, userModel)
});

test('createUser', async t => {
	base.create(t, userRequest, userModel)
});