(function() {
  angular.module('app')
    .directive('match', function match() {
    return {
      require: 'ngModel',
      restrict: 'A',
      scope: {
        match: '='
      },
      link: function(scope, elem, attrs, ctrl) {
        scope.$watch(function() {
          return (ctrl.$pristine && angular.isUndefined(ctrl.$modelValue)) || scope.match === ctrl.$modelValue
        }, function(currentValue) {
          ctrl.$setValidity('match', currentValue)
        })
      }
    }
  })
})();
