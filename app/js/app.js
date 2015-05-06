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
  ])

  .factory('State', [StateFactory]);

function StateFactory() {
  var state = {
    style: "layout.css",
  }

  return state;
}
