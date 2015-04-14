function Get($scope, $http) {
    $http.get('http://eloquentgeek.com/post/').
        success(function(data) {
            $scope.response = data;
        });
}

function Post($scope, $http) {
    $http.post('http://eloquentgeek.com/post/create',
        {
              title: "Title",
              body: "Body",
              tags: ["one", "two", "three"],
        }).
        success(function(data) {
            $scope.response = data;
        }).
        error(function(data) {
            $scope.response = data;
        });
}
