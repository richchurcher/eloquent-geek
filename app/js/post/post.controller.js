angular

  .module('post')

  .controller('PostCtrl', [
    '$scope', 
    'postApiService',
    PostCtrl
  ])

  .directive('postList', postList)

  .filter('trusted', ['$sce', function($sce) {
    return function (text) {
      return $sce.trustAsHtml(text);
    };
  }]);

  //.filter('nlToArray', function() {
    //return function (body) {
      //return body.split('\n');
    //};
  //});


function postList() {
  return {
    controller: PostCtrl,
    templateUrl: '/js/post/postList.html',
    link: function ($scope, elt, attrs) {
      if (!$scope.posts) {
        $scope.loadPosts();
      }
    },
  }
}

function PostCtrl($scope, postApiService) {
  $scope.loadPosts = function() {
    var converter = new showdown.Converter();
    return postApiService.query(function (data) {
      // Markdown
      for (var i = 0; i < data.length; i++) {
        data[i].body = converter.makeHtml(data[i].body);  
      }
      $scope.posts = data;
    });
  };
  
  $scope.deletePost = function(id, i) {
    return postApiService.delete({
      postId: id,
    }).$promise.then(function () {
        $scope.posts.splice(i, 1);
    }, function (error) {
        // TODO: handle error
    });
  };

  $scope.createPost = function(post) {
    if (!post.tags) post.tags = '';
    return postApiService.save({
      title: post.title,
      body: post.body,
      tags: post.tags.split(' '),
    }, function (response) {
      $scope.posts.push(response);
    });
  };
}
