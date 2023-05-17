var totalError
var logEmailerr = document.getElementById("loginEmail");
var logPasserr = document.getElementById("loginPass");
function LoginFields() {
    const input = document.querySelectorAll("input[type=text]");
    input.forEach(ele => {
        if (ele.value == "") {
            console.log("please fill up the form fields");
            logEmailerr.innerHTML = "please fill up Email ";
            logPasserr.innerHTML = "please fill up Password ";
            totalError++;
            return false
        } else {
            logEmailerr.innerHTML = "";
            logPasserr.innerHTML = "";
            return true
        }
    })
}

var logtimer = 0

var follow = document.getElementById("email");
follow.addEventListener("keyup", async (e) => {
    console.log("keyup");
    if (e.target.value == "") {

        logEmailerr.innerHTML = "Enter Email"
        totalError++;
        return
    }
    logEmailerr.innerHTML = ""
    clearTimeout(logtimer)
    logtimer = setTimeout(() => debounceSecond(e.target.value.trim()), 1000);
})

function debounceSecond(find_value) {
    fetch(`http://localhost:8081/emailCheck/${find_value}`)
        .then(res => res.json())
            .then((arr) => {
                console.log(arr);
                if (arr == "Match") {
                    logEmailerr.innerHTML = "";
                    return true

                } else {

                    logEmailerr.innerHTML = "Enter Valid Mail"
                    totalError++;
                    return false
                }
            })
}

function FormValidation(){
    totalError=0;
    LoginFields();
    if(totalError>0){
        return false;
    }else{
        return true;
    }
}