angular

  .module('post')

  .directive('postForm', [
    postForm
  ]);


function postForm() {
  return {
    require: '^postDisplay',
    scope: false,
    templateUrl: '/js/post/postForm.html'
  };
}

