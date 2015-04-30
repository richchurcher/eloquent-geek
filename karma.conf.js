module.exports = function(config){
  config.set({

    basePath : './',

    files : [
      '../src/js/angular.js',
      '../src/js/angular-route.js',
      '../src/js/angular-resource.js',
      '../src/js/angular-mocks.js',
      'js/app.js',
      'js/post/post.js',
      'js/post/post.controller.js',
      'js/test/post.tests.js',
    ],

    autoWatch : true,

    frameworks: ['jasmine'],

    browsers : ['Chrome'],

    plugins : [
            'karma-chrome-launcher',
            'karma-firefox-launcher',
            'karma-jasmine',
            'karma-junit-reporter'
            ],

    junitReporter : {
      outputFile: 'test_out/unit.xml',
      suite: 'unit'
    }

  });
};
