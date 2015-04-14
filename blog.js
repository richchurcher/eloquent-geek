function Get($scope, $http) {
    $http.get('http://eloquentgeek.com/post/').
        success(function(data) {
            $scope.response = data;
        });
}
