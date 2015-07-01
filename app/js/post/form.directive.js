angular

  .module('post')

  .directive('postForm', [
    postForm
  ]);


function postForm() {
  return {
    link: function (scope, element, attrs, postCtrl) {
      scope.create = function (post) {
        postCtrl.createPost(post);
      }
    },
    require: '^postDisplay',
    scope: {
      post: '='
    },
    templateUrl: '/js/post/postForm.html'
  };
}

