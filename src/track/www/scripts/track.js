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
                $scope.alerts.push ( {msg: data.Msg } );
        });
    };
    
    $scope.closeAlert = function(index) {
        $scope.alerts.splice(index, 1);
        $scope.item = $scope.quantity = $scope.price = $scope.vendor = '';
    };
    
    function search(type, name) {
        return $http.get ("/" + type + "?name=" + name).then(function (response)
        {
          return response.data;
        });

    }
    $scope.getItems = function(itemName) {
        return search ("items", itemName);
    };
    
    $scope.getVendors = function(vendorName) {
        return search ("vendors", vendorName);
    };
}
