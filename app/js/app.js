angular

  .module('eg', [
    'post',
  ])

  .config([
    '$resourceProvider', 
    function($resourceProvider) {
      $resourceProvider.defaults.stripTrailingSlashes = false;
    }
  ])

  .controller('EGCtrl', ['$scope', function ($scope) {
    $scope.style = {
      css: ''
    };
  }]);
