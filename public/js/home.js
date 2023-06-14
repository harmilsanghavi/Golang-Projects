async function stateChnage(data){
    let id=data.id
    let value=data.value
    let data2=``;
    if(value=="hold"){
        let data= await fetch(`/getupdateddata?id=${id}&value=${value}`)
        let res=await data.json()
        console.log(res)
        console.log(res[0].Id)
        var div=document.getElementById("hold")
        document.getElementById(res[0].Id).remove()
        console.log("div :-",div)
        data2+=` <div class="card mb-3 bg-warning" style="width: 15rem;" id="${res[0].Id}">
        <img src="./asset/images/task.png" class="card-img-top" alt="...">
        <div class="card-body">
            <h5 class="card-title">${res[0].Todotopic}</h5>
            <p class="card-text">${res[0].Tododescription}</p>
            <select name="stage" id="${res[0].Id}" onchange="stateChnage(this)">
                <option value="">Hold</option>
                <option value="progress">Progress</option>
                <option value="complete">Complete</option>
            </select>
        </div>
    </div>`
    div.innerHTML+=data2
    }
    if(value=="progress"){
        let data=await fetch(`/getupdateddata?id=${id}&value=${value}`)
        let res=await data.json()
        console.log(res)
        console.log(res[0].Id)
        var div=document.getElementById("progress")
        document.getElementById(res[0].Id).remove()
        console.log("div :-",div)
        data2+=` <div class="card mb-3 bg-info" style="width: 15rem;" id="${res[0].Id}">
        <img src="./asset/images/task.png" class="card-img-top" alt="...">
        <div class="card-body">
            <h5 class="card-title">${res[0].Todotopic}</h5>
            <p class="card-text">${res[0].Tododescription}</p>
            <select name="stage" id="${res[0].Id}" onchange="stateChnage(this)">
                <option value="">Progress</option>
                <option value="hold">Hold</option>
                <option value="complete">Complete</option>
            </select>
        </div>
    </div>`
    div.innerHTML+=data2
    }
    if(value=="complete"){
        let data=await fetch(`/getupdateddata?id=${id}&value=${value}`)
        let res=await data.json()
        console.log(res)
        console.log(res[0].Id)
        var div=document.getElementById("complete")
        document.getElementById(res[0].Id).remove()
        console.log("div :-",div)
        data2+=` <div class="card mb-3 bg-success" style="width: 15rem;" id="${res[0].Id}">
        <img src="./asset/images/task.png" class="card-img-top" alt="...">
        <div class="card-body">
            <h5 class="card-title">${res[0].Todotopic}</h5>
            <p class="card-text">${res[0].Tododescription}</p>
            <select name="stage" id="${res[0].Id}" onchange="stateChnage(this)">
                <option value="">Complete</option>
                <option value="hold">Hold</option>
                <option value="progress">Progress</option>
            </select>
        </div>
    </div>`
    div.innerHTML+=data2
    }
    
}

async function getMoreData(data){
    let id=data.id
    let demo=``
    let value="single"
    console.log("id :- ",id)
    let alldatasingleId=await fetch(`/getupdateddata?id=${id}`)
    let res=await alldatasingleId.json()
    console.log(res)
    secondpopupdiv=document.getElementById("atcbody")
    demo+=` <div class="card mb-3" style="width: 15rem;" id="${res[0].Id}" onclick="getMoreData(this)"><img src="./asset/images/task.png" class="card-img-top" alt="..."  data-bs-toggle="modal" data-bs-target="#atcModal" >
    <div class="card-body">
        <h5 class="card-title">${res[0].Todotopic}</h5>
        <p class="card-text">${res[0].Tododescription}</p>
        <select name="stage" id="${res[0].Id}" onchange="stateChnage(this)">
            <option value="" selected>ToDo</option>
            <option value="hold">Hold</option>
            <option value="progress">Progress</option>
            <option value="complete">Complete</option>
        </select>
    </div>`
    secondpopupdiv.innerHTML=demo


    //It is for
    // var btn = document.getElementById("close");
    // var myModal = new bootstrap.Modal(document.getElementById("atcModal"));  

    // let outerdiv=document.getElementsByClassName("unique")
    // outerdiv.id=res[0].Id
    // console.log(outerdiv)

    // let innerdiv=document.getElementsByClassName("unique2")
    // innerdiv.id=res[0].Id
    // console.log(innerdiv)

    // let title=document.getElementById("title")
    // title.innerHTML=res[0].Todotopic

    // let desc=document.getElementById("desc")
    // desc.innerHTML=res[0].Tododescription

    // myModal.show();
  


    // btn.addEventListener("click", function() {       
    //     myModal.hide();
    // });
}
