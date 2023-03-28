let stat=0

const usernameRegex = /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/;
var email=false
var pass=false

function validate(){
    if (!usernameRegex.test(document.getElementById("logintext").value)) {
        usernamevalid.style.display="block"
        buttonid.disabled = true;
      } else {
        usernamevalid.style.display="none"
        buttonid.disabled = false;
      }

      if ((document.getElementById("loginpassword").value).length < 7){
        passwordvalid.style.display="block"
        buttonid.disabled = true;
      }else{
        passwordvalid.style.display="none"
        buttonid.disabled = false;
      }


}

function openHome(){
    students.style.display="none";
    admin.style.display="none"
    departments.style.display="none"
    home.style.display="inline";
  
}

function openStudents(){
    home.style.display="none"
    admin.style.display="none"
    departments.style.display="none"
    students.style.display="inline";
    search()
}

function openDepartment(){
    home.style.display="none"
    admin.style.display="none"
    students.style.display="none";
    departments.style.display="inline"
   
}

function add(){
    home.style.display="none"
    students.style.display="none";
    departments.style.display="none"
    admin.style.display="inline"
}

function addDepartment(){
    adddepartment.style.display="block"
}

function clos(){
    adddepartment.style.display="none"
    addadmin.style.display="none"
    editstudent.style.display="none"
    addstudent.style.display="none"
    editdepartment.style.display="none"
}


function addAdmin(){
    addadmin.style.display="block"
}

function editDepartment(){
    editdepartment.style.display="block"
}



function check(){
    var checkbox1 = document.getElementById("checkbox1");
    var checkbox2 = document.getElementById("checkbox2");

   if (stat==0){
    student.style.display="none"
    admin.style.display="block"
    checkbox2.checked = true;
    stat=1
   }else if (stat==1){
    
    admin.style.display="none"
    student.style.display="block"
    checkbox1.checked = false;
    stat=0
   }




}

function addStudent(){
    addstudent.style.display="block"
}

function editStudent(id){
    editstudent.style.display="block"
   
    var ID = parseInt(id)

    $.ajax({
        url:"/ajax",
        type:"POST",
        contentType: "application/json",
        data: JSON.stringify(ID),
        success: function(item) {
            console.log(item)
            // Display the item data to the user.
            $("#edid").val(item.Id);
            $("#edfname").val(item.Fname);
            $("#edlname").val(item.Lname);
            $("#edemail").val(item.Email);
            $("#edphone").val(item.Phone);
            $("#eddate").val(item.Dob);
            $("#edplace").val(item.Place);
            $("#edusername").val(item.Username);
            
            console.log(item.Lname)

        },
        error: function(xhr, textStatus, errorThrown) {
            alert("Error: " + xhr.responseText);
        }
    })
}


function search(){
    let value = document.getElementById("search").value
    
    $.ajax({
        url:"/search",
        type:"POST",
        contentType: "application/json",
        data: JSON.stringify(value),
        success:function(data){
            console.log("hello")
            console.log(data)
           
           
            var table = '';
            for (let i=0; i<data.length;i++){
                table += "<tr>"
                table += "<td>" + data[i].Id + "</td>"
                table += "<td>" + data[i].Fname + "</td>"
                table += "<td>" + data[i].Lname + "</td>"
                table += "<td>" + data[i].Email + "</td>"
                table += "<td>" + data[i].Phone + "</td>"
                table += "<td>" + data[i].Place + "</td>"
                table += "<td>" + data[i].Dob + "</td>"
                table += "<td>" + data[i].Username + "</td>"
                table += "<td>" + data[i].Dep_id+ "</td>"
                table+='<td> <div class="d-flex flex-row justify-content-center mb-3">' +
                '<div style="width: 70px; margin-right: 15px;" ><button onclick="editStudent('+data[i].Id+')" class="btn btn-success w-100">' +
                    '<i class="fa-solid fa-pen-to-square"></i>' +
                '</button></div>' +
                '<div style="width: 70px;" ><a href="/delete/'+data[i].Id+'/stu" class="btn btn-danger w-100">' +
                    '<i class="fa-solid fa-trash"></i>' +
                '</a></div>' +
            '</div></td>';
                table += "</tr>"
            }
            document.getElementById("searchTable").innerHTML=table;

               
        },
        error: function(xhr, textStatus, errorThrown) {
            alert("Error: " + xhr.responseText);
        }
    })


    


}