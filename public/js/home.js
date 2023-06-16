async function stateChnage(data) {
    let id = data.id
    let value = data.value
    let data2 = ``;
    // document.getElementById("atcModal").remove()
    // var removedata=document.querySelectorAll(".modal-backdrop")
    // removedata.forEach(box => {
    //     box.remove();
    //   });
    if (value == "hold") {
        // document.getElementById("close").addEventListener("click", function() {       
        //     document.getElementById("atcModal").innerHTML='';
        //     });
        let data = await fetch(`/getupdateddata?id=${id}&value=${value}`)
        let res = await data.json()
        console.log(res)
        console.log(res[0].Id)
        var div = document.getElementById("hold")
        document.getElementById(res[0].Id).remove()
        console.log("div :-", div)
        let imgSrc = "./asset/images/task.png";
        if (res[0].Attachments.length > 0) {
            imgSrc = res[0].Attachments[0].Files;
        }
        data2 += ` <div class="card mb-3 bg-warning" style="width: 15rem;" id="${res[0].Id}">
        <img src="${imgSrc}" class="card-img-top" alt="...">
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
        div.innerHTML += data2
    }
    if (value == "progress") {
        let data = await fetch(`/getupdateddata?id=${id}&value=${value}`)
        let res = await data.json()
        // console.log(res)
        // console.log(res[0].Id)
        // console.log("demo",res[0].Attachments[0].Files)
        var div = document.getElementById("progress")
        document.getElementById(res[0].Id).remove()
        console.log("div :-", div)
        let imgSrc = "./asset/images/task.png";
        if (res[0].Attachments.length > 0) {
            imgSrc = res[0].Attachments[0].Files;
        }
        data2 += ` <div class="card mb-3 bg-info" style="width: 15rem;" id="${res[0].Id}">
        <img src="${imgSrc}" class="card-img-top" alt="...">
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
        div.innerHTML += data2
    }
    if (value == "complete") {
        let data = await fetch(`/getupdateddata?id=${id}&value=${value}`)
        let res = await data.json()
        console.log(res)
        console.log(res[0].Id)
        var div = document.getElementById("complete")
        document.getElementById(res[0].Id).remove()
        console.log("div :-", div)
        let imgSrc = "./asset/images/task.png";
        if (res[0].Attachments.length > 0) {
            imgSrc = res[0].Attachments[0].Files;
        }
        data2 += ` <div class="card mb-3 bg-success" style="width: 15rem;" id="${res[0].Id}">
        <img src="${imgSrc}" class="card-img-top" alt="...">
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
        div.innerHTML += data2
    }

}

async function getMoreData(data) {
    let id = data.id
    let value = "single"
    console.log("id :- ", id)
    let alldatasingleId = await fetch(`/getupdateddata?id=${id}`)
    let res = await alldatasingleId.json()
    console.log(res)
    secondpopupdiv = document.getElementById("atcbody")
    let imgSrc = "./asset/images/task.png";
    var imagearray=[]
    let lengthofimage=res[0].Attachments.length
    if (lengthofimage > 0) {
        imgSrc = res[0].Attachments[0].Files;
        for (let i=0; i<lengthofimage; i++) {
            let url=res[0].Attachments[i].Files
            var imageshow=`<img src="${url}" width="30%"/>`;
            imagearray.push(imageshow);
        }
    }
    let demo = ``
    demo += `<div class="card mb-3" style="width: 15rem;" id="${res[0].Id}" >
        <form action="/addtodo" method="post" enctype="multipart/form-data">
        <img src="${imgSrc}" class="card-img-top" alt="..." data-bs-toggle="modal" data-bs-target="#atcModal" >
        <input type="text" name="todo_id" value="${res[0].Id}" hidden>
        <div class="card-body">
        <h5 class="card-title">
        <div class="mb-3">
        <label for="" class="form-label">Todo Name:</label>
        <input type="text" class="form-control" name="todotopic" id="todotopic" value="${res[0].Todotopic}">
        </div>
        </h5>
        <div class="mb-3">
        <label for="" class="form-label">Todo Desc:</label>
        <input type="text" class="form-control" name="tododescription" id="tododescription" value="${res[0].Tododescription}">
        </div>
        <div class="mb-3">
        <label for="" class="form-label">Select:</label>
        <input type="file" class="form-control" name="upload" multiple="multiple" id="upload">
        </div>
        <div class="mb-3">
            <label for="" class="form-label">Attachments</label>
            <p>${imagearray}</p>
        </div>
        <div class="mb-3">
            <select name="stage" id="${res[0].Id}" onchange="stateChnage(this)">
            <option value="hold" ${res[0].Stage === "todo" ? "selected" : ""}>Todo</option>
            <option value="hold" ${res[0].Stage === "hold" ? "selected" : ""}>Hold</option>
            <option value="progress" ${res[0].Stage === "progress" ? "selected" : ""}>Progress</option>
            <option value="complete" ${res[0].Stage === "complete" ? "selected" : ""}>Complete</option>
        </select>
        </div>
        <div class="mb-3">
        <input type="submit" value="Add Todo" class="form-control btn btn-success">
        </div>
        </div>
    </form>
    </div>`;
    secondpopupdiv.innerHTML = demo
}
