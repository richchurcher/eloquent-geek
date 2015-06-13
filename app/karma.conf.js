module.exports = function(config){
  config.set({

    basePath : './',

    files : [
      '/media/work/src/js/angular.js',
      '/media/work/src/js/angular-route.js',
      '/media/work/src/js/angular-resource.js',
      '/media/work/src/js/angular-mocks.js',
      'js/app.js',
      'js/post/post.js',
      'js/post/apiService.js',
      'js/post/controller.js',
      'js/test/*_tests.js',
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
