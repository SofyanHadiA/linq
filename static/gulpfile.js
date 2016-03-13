var gulp = require('gulp'),
   uglify = require('gulp-uglify');

gulp.task('minify', function () {
   gulp.src('js/bundle.js')
      .pipe(uglify())
      .pipe(gulp.dest('dist'))
});