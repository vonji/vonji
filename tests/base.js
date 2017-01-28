import test from 'ava'
import * as defaultRequest from 'request';
// https://stackoverflow.com/questions/36720335/ava-syntaxerror-unexpected-token-import
export const request = defaultRequest.defaults({
	baseUrl: 'http://localhost:1618',
	headers: { 'User-Agent': 'Request-Promise' },
	resolveWithFullResponse: true,
	json: true
});

export function sanitize(object) {
	delete object.CreatedAt;
	delete object.UpdatedAt;
	delete object.ID;
	return object;
}

export async function get(t, objectRequest, model) {
	let id = await objectRequest.post('', { body: model }).then(response => response.body.ID);
	let object = await objectRequest.get('' + id).then(response => {
		t.is(response.statusCode, 200);
		sanitize(response.body);
		return response.body;
	});
	t.deepEqual(model, object);
}

export async function create(t, objectRequest, model) {
	let created = await objectRequest.post('', { body: model }).then(response => {
		t.is(response.statusCode, 201);
		sanitize(response.body);
		return response.body;
	});
	t.deepEqual(model, created);
}