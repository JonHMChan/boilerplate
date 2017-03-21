const gulp      = require('gulp');
const sass      = require('gulp-ruby-sass');
const babel     = require('gulp-babel');
const pump      = require('pump');
const child     = require('child_process');
const ps        = require('ps-node');

/* ----------------------------------------------------------------------------
 * Instructions
 * ------------------------------------------------------------------------- */

gulp.task('default', () => {
    console.log(" ")
    console.log("You can run the following commands:")
    console.log("-----------------------------------------")
    console.log("    `gulp server` - starts the go server and watches for changes (.go and .html)")
    console.log("    `gulp assets` - processes and watches web asset changes (.sass, .js, and images)")
    console.log(" ")
})

/* ----------------------------------------------------------------------------
 * Assets
 * ------------------------------------------------------------------------- */

// Start watch
gulp.task('assets', ['assets:sass', 'assets:react', 'assets:buildjs'], function() {
    gulp.watch('static/css/src/**/*.scss',['assets:sass']);
    gulp.watch('static/js/src/**/*.js',['assets:buildjs']);
    gulp.watch('static/js/src/**/*.jsx',['assets:react']);
});

// Compile Sass into css
gulp.task('assets:sass', () =>
    sass('static/css/src/**/*.scss')
        .on('error', sass.logError)
        .pipe(gulp.dest('static/css/dist'))
);

gulp.task('assets:react', function () {
    return gulp.src('static/js/src/**/*.jsx')
        .pipe(babel({
            presets: ['es2015', 'react']
        }))
        .pipe(gulp.dest('static/js/dist/'));
});

// Compile ES2015
gulp.task('assets:buildjs', () => {
    return gulp.src('static/js/src/**/*.js')
        .pipe(babel({
            presets: ['es2015']
        }))
        .pipe(gulp.dest('static/js/dist/'));
});

/* ----------------------------------------------------------------------------
 * Server
 * ------------------------------------------------------------------------- */

// Start server with build and start, then watch
gulp.task('server', ['server:build'], () => {
    gulp.watch('**/*.go', ['server:build'])
    gulp.watch('**/*.html', ['server:build'])
})

// Go compilation of the server
gulp.task('server:build', function() {
    var build = child.spawnSync('go', ['install']);
    if (build.stderr.length) {
        var lines = build.stderr.toString()
            .split('\n').filter(function(line) {
                return line.length
        });
        for (var l in lines) console.log(lines[l])
    }
    return build;
});
