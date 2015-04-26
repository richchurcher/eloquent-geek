angular

  .module("eg", [
    "postAPI"
  ])

  .config([
    '$resourceProvider', 
    function($resourceProvider) {
      $resourceProvider.defaults.stripTrailingSlashes = false;
    }
  ]);
