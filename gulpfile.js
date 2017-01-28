const gulp = require('gulp');
const babel = require('gulp-babel');

gulp.task('default', () => {
	return gulp.src('src/main.js')
		.pipe(babel({
			presets: ['es2015', 'stage-2'],
      plugins: ['transform-runtime'], 
		}))
		.pipe(gulp.dest('dist'));
});
