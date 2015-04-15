var app = angular.module("eg", ["ngResource", "postAPI"]);

var postAPI = angular.module("postAPI", ["ngResource"]);

postAPI.factory("Post", ["$resource", 
    function ($resource) {
        return $resource("/post", {}, {});
    }
]);

postAPI.controller("PostIndexCtrl", ["$scope", "Post",
    function($scope, Post) {
        Post.query(function(data) {
            $scope.response = data;
        });
    }
]);

//postAPI.controller("CreateCtrl", ["$scope", "Post",
    //function($scope, Post) {
        //Post.save({
            //title: "Title",
            //body: "Body",
            //tags: ["one", "two", "three"],
        //});
    //}
//]);
