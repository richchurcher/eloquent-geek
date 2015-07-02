angular

  .module('post')

  .directive('postForm', [
    postForm
  ]);


function postForm() {
  return {
    link: function (scope, element, attrs, postCtrl) {
      scope.create = function () {
        postCtrl.createPost();
      }
    },
    require: '^postDisplay',
    scope: false,
    templateUrl: '/js/post/postForm.html'
  };
}

