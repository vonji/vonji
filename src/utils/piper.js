export const piper = fn => (req, res, next) => fn(req, res, next).catch(next);
