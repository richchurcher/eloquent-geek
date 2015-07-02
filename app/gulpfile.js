var gulp = require('gulp');
var concat = require('gulp-concat');
var ngAnnotate = require('gulp-ng-annotate');
var stripDebug = require('gulp-strip-debug');
var uglify = require('gulp-uglify');
var sourcemaps = require('gulp-sourcemaps');
 
gulp.task('js', function () {
  return gulp.src([
      // Naming scheme maintains load order
      'js/app.js', 
      'js/**/*.module.js', 
      'js/**/*.controller.js',
      'js/**/*.directive.js',
      'js/**/*.service.js'
    ])
    .pipe(sourcemaps.init())
      .pipe(concat('js/all.js'))
      //.pipe(stripDebug())
      .pipe(ngAnnotate())
      .pipe(uglify(true))
    .pipe(sourcemaps.write())
    .pipe(gulp.dest('.'));
});

gulp.task('watch', ['js'], function () {
  gulp.watch('js/**/*.js', ['js']);
});
