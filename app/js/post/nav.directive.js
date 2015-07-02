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
      scope.previousPost = function (current) {
        postCtrl.loadPost(current, 'previous');
      };
      scope.nextPost = function (current) {
        postCtrl.loadPost(current, 'next');
      };
    },
    require: '^postDisplay',
    scope: {
      id: '@'
    },
    templateUrl: '/js/post/postNav.html'
  };
}

