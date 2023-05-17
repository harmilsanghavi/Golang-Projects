var totalError;
var errfname=document.getElementById("errfname");
var errlname=document.getElementById("errlname");
var errpassword=document.getElementById("errpassword");
var errnumber=document.getElementById("errnumber");
var erremail=document.getElementById("erremail");
const previewImage = (event) => {
    const array_of_allowed_files = ['png', 'jpeg', 'jpg'];
    var tweet_image = document.getElementById("image");
    const file_extension = tweet_image.value.slice(((tweet_image.value.lastIndexOf('.') - 1) >>> 0) + 2);
    var imgsize = Math.round(tweet_image.files[0].size / 1048576);
    if (imgsize > 2) {
        document.getElementById("err2").innerHTML = "upload image whose size is less than 2mb";
        totalError++;
    }
    else {
        if (array_of_allowed_files.includes(file_extension)) {
            document.getElementById("err2").innerHTML = "";
            /**
            * Get the selected files.
            */
            const imageFiles = event.target.files;
            /**
            * Count the number of files selected.
            */
            const imageFilesLength = imageFiles.length;
            /**
            * If at least one image is selected, then proceed to display the preview.
            */
            if (imageFilesLength > 0) {
                /**
                 * Get the image path.
                 */
                const imageSrc = URL.createObjectURL(imageFiles[0]);
                /**
                 * Select the image preview element.
                 */
                const imagePreviewElement = document.querySelector("img");
                /**
                 * Assign the path to the image preview element.
                 */
                imagePreviewElement.src = imageSrc;
                /**
                 * Show the element by changing the display value to "block".
                 */
                imagePreviewElement.style.display = "block";
            }
        } else {
            document.getElementById("err2").innerHTML = "upload only jpg,jpeg and png formate files";
            totalError++;
        }

    }

};

function checkname(e) {
   
    var regName = /^[a-zA-Z]+$/;
    var name=e.value.trim()
    if (regName.test(name)) {
        if(e.id=="fname"){
            errfname.innerHTML=""
        }else{
            errlname.innerHTML=""
        }
    } else if(name.trim()==""){
        if(e.id=="fname"){
            errfname.innerHTML="enter first name"
        }else{
            errlname.innerHTML="enter Last Name"
        }
        totalError++;
    }else{
        if(e.id=="fname"){
            errfname.innerHTML="enter valid first name"
        }else{
            errlname.innerHTML="enter valid Last Name"
        }
        totalError++;
    }
}   

var timer=0
var search=document.getElementById("email1")
search.addEventListener("keyup",async(e)=>{
    if(e.target.value==""){
        
        erremail.innerHTML="enter email"
        return
    }
    erremail.innerHTML=""
    clearTimeout(timer)
    timer=setTimeout(() => debounceFirst(e.target.value.trim()),1000);
})
function debounceFirst(search_value) {
    fetch(`http://localhost:8081/emailCheck/${search_value}`)
    .then(res => res.json())
    .then((arr) => {
        console.log(arr)
        if(arr=="Match"){
            erremail.innerHTML="enter valid email"
        }else{
            erremail.innerHTML=""
        }
    })
}

function FillValidation(){
    const input = document.querySelectorAll("input[type=text]")
    input.forEach(ele=>{
        if(ele.value == " "){
            erremail.innerHTML="please fill the Email"
            errpassword.innerHTML="please fill the password"
            errfname.innerHTML="please fill the first name"
            errlname.innerHTML="please fill the last name"
            errnumber.innerHTML="please fill the number"
            totalError++;
        }else{
            erremail.innerHTML=" "
            errpassword.innerHTML=" "
            errfname.innerHTML=" "
            errlname.innerHTML=" "
            errnumber.innerHTML=" "
        }
    })
}
function FormValidation(){
    totalError=0;
    FillValidation();
    if(totalError>0){
        return false;
    }else{
        return true;
    }
}









































