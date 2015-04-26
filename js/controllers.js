var app = angular.module("eg", ["ngResource", "postAPI"]);

app.config(['$resourceProvider', function($resourceProvider) {
  $resourceProvider.defaults.stripTrailingSlashes = false;
}]);

var postAPI = angular.module("postAPI", ["ngResource"]);

postAPI.factory("Post", ["$resource", 
  function postFactory(resource) {
    return resource("/post/:postId");
  }
]);

postAPI.controller("PostCtrl", ["$scope", "Post",
  function($scope, Post) {
    $scope.loadPosts = function() {
      Post.query(function (data) {
        $scope.posts = data;
      });
    };
    
    $scope.deletePost = function(id, i) {
      Post.delete({
        postId: id,
      }).$promise.then(function () {
          $scope.posts.splice(i, 1);
      }, function (error) {
          // TODO: handle error
      });
    };

    $scope.createPost = function(post) {
      if (!post.tags) post.tags = "";
      Post.save({
        title: post.title,
        body: post.body,
        tags: post.tags.split(" "),
      }, function (response) {
        $scope.posts.push(response);
      });
    };
  }
])

postAPI.directive("postList", function () {
  return {
    controller: postAPI.PostCtrl,
    link: function ($scope, elt, attrs) {
      if (!$scope.posts) {
        $scope.loadPosts();
      }
    },
  }
});
