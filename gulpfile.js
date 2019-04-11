
var gulp   = require('gulp')
    sass   = require('gulp-ruby-sass') 
    notify = require("gulp-notify") 
    bower  = require('gulp-bower')
    uglify = require('gulp-uglify')
    rjs    = require('gulp-requirejs');
    browserSync = require('browser-sync').create();
    critical = require('critical');

var config = {
     sassPath: './resources/sass',
    jsPath:   './resources/js',
     bowerDir: './bower_components' ,
    nodeDir: './node_modules' 
}


gulp.task('icons', function() { 
    return gulp.src(config.bowerDir + '/fontawesome/fonts/**.*') 
        .pipe(gulp.dest('./web/dist/fonts')); 
});

gulp.task('glyphicons', function() { 
    return gulp.src(config.bowerDir + '/bootstrap-sass-official/assets/fonts/bootstrap/***.*') 
        .pipe(gulp.dest('./web/dist/fonts/bootstrap')); 
});



gulp.task('css', function() { 
    return gulp.src(config.sassPath + '/style.scss')
         .pipe(sass({
             style: 'compressed',
             loadPath: [
                 './resources/sass',
                 config.bowerDir + '/bootstrap-sass-official/assets/stylesheets',
                 config.bowerDir + '/fontawesome/scss',
             ]
         }) 
            .on("error", notify.onError(function (error) {
                 return "Error: " + error.message;
             }))) 
         .pipe(gulp.dest('./web/dist/css')); 
});

gulp.task('bower', function() { 
    return bower()
         .pipe(gulp.dest(config.bowerDir)) 
});


gulp.task('js', function() {
  return gulp.src([


    config.jsPath + '/*.js', config.nodeDir + '/requirejs/require.js'])

    .pipe(uglify())
    .pipe(gulp.dest('./web/dist/js'));
});

gulp.task('requireJs', function() {
    rjs({
        baseUrl: config.jsPath + '/app',
        shim: {
            // standard require.js shim options
        },
        paths: {
            jQuery: '../../../' + config.bowerDir + '/jquery/dist/jquery.min',
            'fastclick': '../../../' + config.nodeDir + '/fastclick/lib/fastclick',
            'basket'   : '../lib/basket',
            'bootstrap': '../../../' + config.bowerDir + '/bootstrap-sass-official/assets/javascripts/bootstrap.min'
        },
        out: 'main.js',
        name: "main",
        findNestedDependencies: true
        // ... more require.js options
    })
    .pipe(uglify())
    .pipe(gulp.dest('./web/dist/js')); // pipe it to the output DIR
});

gulp.task('browser-sync', function() {
    browserSync.init({
        proxy: "http://localhost:8888"
    });
});

gulp.task('critical', function (cb) {
  critical.generate({
    inline: false,
    base: 'templates',
    src: 'layout.html',
    css: ['web/dist/css/style.css'],
    dimensions: [{
      width: 320,
      height: 480
    },{
      width: 768,
      height: 1024
    },{
      width: 1280,
      height: 960
    }],
    dest: 'templates/critical.css',
    minify: true,
    extract: true,
    ignore: ['@font-face', '@import']
  });
});

// Rerun the task when a file changes
 gulp.task('watch', function() {
     gulp.watch(config.sassPath + '/**/*.scss', ['css']); 
});

  gulp.task('default', ['bower', 'icons', 'glyphicons', 'css', 'js', 'requireJs', 'critical']);