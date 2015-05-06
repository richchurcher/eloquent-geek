angular

  .module("style")

  .controller("StyleCtrl", [
    "$scope", 
    "State",
    StyleCtrl
  ])

  .directive("styleSwitcher", styleSwitcher);

function styleSwitcher() {
  return {
    controller: StyleCtrl,
    templateUrl: "/js/style/styleSwitcher.html",
    //link: function ($scope, elt, attrs) {
      //// ...
    //},
  }
}

function StyleCtrl($scope, State) {
  $scope.layoutFile = State.style;
}
