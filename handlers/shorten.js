function httpRequest(method, rul, body){
    return new Promise(function(resolve, reject){
        var req = new XMLHttpRequest;
        req.open(method, url);

        req.onload = function(){
            if(req.status == 200) {
                resolve(req.response);
            } else{
                reject(req.response);
            }
        };

        req.onerror = function(){
            reject(Error("Fail to httpRequest " + url));
        };

        req.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
        req.send(body);
    });
};

window.onload = function(){
    var origin = document.getElementById("Origin");
    var shorten = document.getElementById('Shorten')
    var shortenbtn = document.getElementsByClassName("shorten-btn");

    shortenbtn.onclick = function(event){
        alert(origin.value);
        return origin.value;
    }

}