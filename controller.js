
app.controller('StudentCtrl', ['$scope', 'studentservice',
function ($scope, studentservice) {
    var baseurl = 'http://localhost:8080/';
    $scope.button = "Save";
  
    $scope.SaveUpdate = function () {
       
        var student = {
            id:$scope.id,
            name: $scope.name,
            phone: $scope.phone,
          
        }
        if ($scope.button == "Save") {
            var api = baseurl + 'savestudent/';
            var saveStudent = studentservice.post(api, student);
            saveStudent.then(function (response) {
             
                if (response.data != "") {
                    alert("saved");
                    $scope.GetStudents();
                    $scope.clear();
                } else {
                    alert("Some error");
                } });
        }

        else

        {              
            var stud = {
                id:$scope.id,
                name: $scope.name,
                phone: $scope.phone, 
            }
            var api = baseurl + 'student/';
            var UpdateStudent = studentservice.put(api,stud,$scope.id);
            UpdateStudent.then(function (response) {
                if (response.data != "") {
                    alert("data updated");
                    $scope.GetStudents();
                    $scope.clear();

                } else {
                    alert("Some error");
                }

            });
        }
    }


    $scope.GetStudents = function () {
        var api = baseurl + 'GetStudents/';
        var student = studentservice.getAll(api);
        student.then(function (response) {
            
            $scope.studnets = response.data;

        });
    }
    $scope.GetStudents();
    
    $scope.getstdbyid = function (sdata)
    {
        
        var api = baseurl + 'getstdbyid/';
        var student = studentservice.getbyID(api, sdata.id);
        student.then(function (response) {
            $scope.id = response.data.id;
            $scope.name = response.data.name;
            $scope.phone = response.data.phone;
            
            $scope.button = "Update";
        });
    }

    $scope.deletestudent = function (sdata)
    {
        
        var api = baseurl + 'deletestudent/' + sdata.id;
        var deleteStudent = studentservice.delete(api);
        deleteStudent.then(function (response) {
            if (response.data != "") {
                alert("deleted");
                $scope.GetStudents();
                $scope.clear();

            } else {
                alert("Some error");
            }

        });
    }

    $scope.clear = function () {
        //$scope.id = "";
        $scope.name = "";
        $scope.phone = "";
        $scope.button = "Save";
    }

}]);