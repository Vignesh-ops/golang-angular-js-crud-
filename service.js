app.service('studentservice', function ($http) {

    
    this.post = function (api, student) {
        var request = $http({
            method: "post",
            url: api,
            data: student,
          //  headers:{'content-Type':'application/json'},
          // 
        });
        return request;
    }
    this.put = function (api, stud,id) {
        var request = $http({
            method: "put",
            url: api+id,
            data: stud
        });
        return request;
    }
    this.delete = function (api) {
        var request = $http({
            method: "delete",
            url: api
        });
        return request;
    }
    this.getAll = function (api) {

        url = api;
        return $http.get(url);
    }
    this.getbyID = function (api,id) {
        url = api + id;
        return $http.get(url);
    }
});