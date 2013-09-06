var app = angular.module('track', ['ui.bootstrap']);
app.directive('ngBlur', function() {
  return function( scope, elem, attrs ) {
    elem.bind('blur', function() {
      scope.$apply(attrs.ngBlur);
    });
  };
});
function TrackCtrl ($http, $scope) {
    $scope.navType = "tabs";
    $scope.alerts = [];
    $scope.items_to_buy = [];
    $scope.recent_data = [];
    
    $scope.save_thing_to_buy = function() {
        $http.post ("/add",
                { "item" : $scope.item_to_buy }).success (function (response) {
                  $scope.items_to_buy = response;
                  $scope.item_to_buy = '';
                });
    };

    $scope.save = function() {
        $http.post ("/save", 
            { "item" : $scope.item,
              "quantity" : $scope.quantity,
              "price" : $scope.price,
              "vendor" : $scope.vendor }).success (function (data) {
                $scope.alerts = [];
                $scope.alerts.push ( {msg: data.Msg } );
                $scope.getMonthlySpendings();
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

    $scope.getItemsToBuy = function() {
        $http.get ("/get_items_to_buy").then(function (response)
            {
                $scope.items_to_buy = response.data;
            });
    }

    $scope.getMonthlySpendings = function() {
        $http.get ("/get_monthly_spendings").then(function (response)
            {
                $scope.monthTotals = response.data;
            });
    }
    
    $scope.removeFromItemsToBuy = function(index) {
        $http.post ("/remove_item_to_buy", { 
            "item" : $scope.items_to_buy[index] }).success (function (response) 
                { $scope.items_to_buy = response; });
    };
    
    $scope.getItems = function(itemName) {
        return search ("items", itemName);
    };
    
    $scope.getVendors = function(vendorName) {
        return search ("vendors", vendorName);
    };

    $scope.getRecentPurchasesOfItem = function() {
        $http.post ("/recent", 
            { "item" : $scope.item }).success (function (data) {
              $scope.recent_data = data;
            });
    };

    $scope.recalculate = function() {
       for (i = 0; i< $scope.recent_data.length; ++i)
       {
            $scope.recent_data[i].ComparitivePrice = $scope.quantity * $scope.recent_data[i].NormalizedPrice;
       }
    }

    $scope.clear = function() {
        $scope.recent_data = [];
    }
    
    $scope.getItemsToBuy();
    $scope.getMonthlySpendings();
}
