angular

  .module('post')

  .directive('postNav', [
    postNav
  ]);


function postNav() {
  return {
    link: function (scope, element, attrs, postCtrl) {
      scope.firstPost = function () {
        postCtrl.loadPost('first');
      };
      scope.latestPost = function () {
        postCtrl.loadPost('latest');
      };
      scope.previousPost = function () {
      };
      scope.nextPost = function () {
      };
    },
    require: '^postDisplay',
    scope: {
      post: '='
    },
    templateUrl: '/js/post/postNav.html'
  };
}

