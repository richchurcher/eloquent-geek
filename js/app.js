var app = angular.module('eg', ['ngResource']);
app.factory("Post", function($resource) {
    return $resource("/post");
});

app.controller("PostIndexCtrl", function ($scope, Post) {
    Post.query(function(data) {
        $scope.response = data;
    });
});

app.controller("CreateCtrl", function ($scope, Post) {
    Post.save({
        title: "Title",
        body: "Body",
        tags: ["one", "two", "three"],
    });
});

