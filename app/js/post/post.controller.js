angular

  .module("post")

  .controller("PostCtrl", [
    "$scope", 
    "Post",
    PostCtrl
  ])

  .directive("postList", postList)

  .factory("Post", [
    "$resource", 
    postFactory
  ])

  .filter("nlToArray", function() {
    return function (body) {
      return body.split('\n');
    };
  });

function postList() {
  return {
    controller: PostCtrl,
    templateUrl: "/js/post/postList.html",
    link: function ($scope, elt, attrs) {
      if (!$scope.posts) {
        $scope.loadPosts();
      }
    },
  }
}

function postFactory(resource) {
  return resource("/post/:postId");
}

function PostCtrl($scope, Post) {
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