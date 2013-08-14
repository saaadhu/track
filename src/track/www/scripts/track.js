angular.module('track', ['ui.bootstrap']);
function TrackCtrl ($http, $scope) {
    $scope.navType = "tabs";
    $scope.alerts = [];

    $scope.save = function() {
        $http.post ("/save", 
            { "item" : $scope.item,
              "quantity" : $scope.quantity,
              "price" : $scope.price,
              "vendor" : $scope.vendor }).success (function (data) {
                $scope.alerts = [];
                $scope.alerts.push (data.Msg);
        });
    };
    $scope.getItems = function(itemName) {
        return $http.get ("/items?name=" + itemName).then(function (response)
        {
          return response.data;
        });
    };
}
