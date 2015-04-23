var app = angular.module("eg", ["ngResource", "postAPI"]);

app.config(['$resourceProvider', function($resourceProvider) {
  // Don't strip trailing slashes from calculated URLs
  $resourceProvider.defaults.stripTrailingSlashes = false;
}]);

var postAPI = angular.module("postAPI", ["ngResource"]);

postAPI.factory("Post", ["$resource", 
  function postFactory(resource) {
    return resource("/post/:postId");
  }
]);

postAPI.controller("PostIndexCtrl", ["$scope", "Post",
  function($scope, Post) {
    Post.query(function(data) {
      $scope.posts = data;
    });

    $scope.deletePost = function(id) {
      Post.delete({
        postId: id,
      }), function (response) {
        $scope.deleteResponse = "API DELETE response: " + angular.toJson(response);
      }
    };
  }
]);

postAPI.controller("CreateCtrl", ["$scope", "Post",
  function($scope, Post) {
    $scope.createPost = function(post) {
      Post.save({
        title: post.title,
        body: post.body,
        tags: post.tags.split(" "),
      }, function (response) {
        $scope.response = "API POST response: " + angular.toJson(response);
      });
    };
  }
]);

postAPI.directive("removeButton",
  function () {
    return {
      link: function ($scope, elt, attrs) {
        $scope.remove = function() {
          elt.parent().remove();
        };
      }
    }
  }
);
