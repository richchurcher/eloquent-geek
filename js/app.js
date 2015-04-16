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
            $scope.response = data;
        });
    }
]);

postAPI.controller("CreateCtrl", ["$scope", "Post",
    function($scope, Post) {
        Post.save({
            title: "Title",
            body: "Body",
            tags: ["one", "two", "three"],
        }, function (response) {
            $scope.response = "API POST response: " + angular.toJson(response);
        });
    }
]);

postAPI.controller("DeleteCtrl", ["$scope", "Post",
    function($scope, Post) {
        Post.delete({
            postId: 5629499534213120,
        }), function (response) {
            $scope.response = "API DELETE response: " + angular.toJson(response);
        };
    }
]);
