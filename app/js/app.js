angular

  .module("eg", [
    "post",
    "style"
  ])

  .config([
    '$resourceProvider', 
    function($resourceProvider) {
      $resourceProvider.defaults.stripTrailingSlashes = false;
    }
  ]);
