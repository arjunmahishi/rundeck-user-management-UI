window.onload = () => {
    getUsers(data => loadUserTable(data.users, data.allowance))
}

const baseURL = "/users"
const table = document.querySelector("#user-table")
var data;

const makeRequest = (type, data, callback) => {
    var xhr = new XMLHttpRequest();
    xhr.withCredentials = true;

    xhr.addEventListener("readystatechange", function () {
        if (this.readyState === 4) {
            callback(this.responseText)
        }
    });

    xhr.open(type, baseURL);
    xhr.send(data);
    console.log(xhr)
}

const getUsers = (callback) => {
    makeRequest("GET", null, (res) => {
        data = JSON.parse(res)
        document.querySelector("#nav-username").innerHTML = `Hi, ${data.users[0].username}`
        callback(data) 
    })
}

const updateUser = (oldUsername, newUser) => {
    var data = {oldUsername: oldUsername, newUser: newUser}
    makeRequest("PUT", JSON.stringify(data), res => getUsers(newData => loadUserTable(newData.users, newData.allowance)))
}

const createUser = (newUser) => makeRequest("POST", JSON.stringify(newUser), res => 
getUsers(newData => loadUserTable(newData.users, newData.allowance)))

const deleteUser = (username) => {
    var data = {username: username}
    if (confirm("Are you sure you want to delete this user from rundeck? this cannot be undone")){
        makeRequest("DELETE", JSON.stringify(data), res => getUsers(newData => loadUserTable(newData.users, newData.allowance)))
    }
}

const loadUserTable = (users, allowance) => {
    // allowance: 0 - add, 1 - edit, 2 - delete
    table.innerHTML = ""
    users.map((user, i) => {
        var roles = ""
        user.roles.sort().map((role) => {
            roles += `<span class="custom-badge badge badge-secondary">${role}</span>`                
        })

        table.innerHTML += `
            <tr class="">
                <td class="">${i+1}</td>
                <td class="">${user.username}</td>
                <td class="">${roles}</td>
                <td>
                    <button class="btn btn-outline-dark custom-button" data-user="${i}"
                    data-toggle="modal"` + ((allowance[1] || i===0) ? "":" disabled ") + `data-target="#user-data">Edit</button>
                    <button class="btn btn-outline-danger custom-button"` + (allowance[2] ? "":" disabled ") + ` 
                    onclick="deleteUser('${user.username}')">Delete</button>
                </td>
            </tr>
        `
    })
    table.innerHTML += `
        <tr>
            <td></td><td></td><td></td>
            <td><button class="btn btn-outline-success custom-button" 
            data-toggle="modal" data-target="#user-data"` + (allowance[0] ? "":" disabled ") + `>Add</button></td>
        </tr>
    `
}

const searchHandler = (query) => {
    loadUserTable(data.users.filter(user => user.username.includes(query.toLowerCase())), data.allowance)
}

const togglePassword = (checked) => {
    if (checked) {
        document.querySelector("#modal-password").type = "text"
    }else {
        document.querySelector("#modal-password").type = "password"
    }
}

const handleCreateSubmit = () => {
    var newUser = {} 
    newUser["username"] = document.querySelector("#modal-username").value.trim().toLowerCase()
    newUser["roles"] = document.querySelector("#modal-roles").value.trim().toLowerCase().split(",")
    newUser["password"] = document.querySelector("#modal-password").value.trim()
    createUser(newUser)
}

const handleUpdateSubmit = (oldUsername) => {
    var newUser = {} 
    newUser["username"] = document.querySelector("#modal-username").value.trim().toLowerCase()
    newUser["roles"] = document.querySelector("#modal-roles").value.trim().toLowerCase().split(",")
    newUser["password"] = document.querySelector("#modal-password").value.trim()
    updateUser(oldUsername, newUser)
}

$('#user-data').on('show.bs.modal', function (event) {
    var button = $(event.relatedTarget)
    var i = button.data('user')

    if (i !== undefined) { // Edit 
        var user = data.users[i]
        var currUser = data.users[0]
        
        if (!currUser.roles.includes("admin")){
            document.querySelector("#modal-roles").disabled = true
        }
        
        if (currUser.username === user.username || currUser.roles.includes("admin")){
            document.querySelector("#modal-password").value = user.password
        }else{
            document.querySelector("#modal-password").disabled = true
        }

        document.querySelector("#modal-roles").value = user.roles
        document.querySelector("#modal-username").value = user.username
        document.querySelector("#modal-submit").setAttribute("onclick", `handleUpdateSubmit("${user.username}")`)
    }else { // Add
        document.querySelector("#modal-username").value = ""
        document.querySelector("#modal-roles").value = "" 
        document.querySelector("#modal-password").value = ""
        document.querySelector("#modal-submit").setAttribute("onclick", "handleCreateSubmit()")
    }
})