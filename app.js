angular

  .module("eg", [
    "post"
  ])

  .config([
    '$resourceProvider', 
    function($resourceProvider) {
      $resourceProvider.defaults.stripTrailingSlashes = false;
    }
  ]);
