angular

  .module('post')

  .factory('postApiService', [
    '$resource', 
    postFactory
  ]);

function postFactory(resource) {
  return resource('/post/:postId');
}
