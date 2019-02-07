window.onload = () => {
    getUsers(users => loadUserTable(users))
}

const baseURL = "/users"
const table = document.querySelector("#user-table")
var usersList;

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
}

const getUsers = (callback) => {
    makeRequest("GET", null, (res) => {
        usersList = JSON.parse(res)
        callback(usersList) 
    })
}

const updateUser = (oldUsername, newUser) => {
    var data = {oldUsername: oldUsername, newUser: newUser}
    makeRequest("PUT", JSON.stringify(data), res => getUsers(newList => loadUserTable(newList)))
}

const createUser = (newUser) => makeRequest("POST", JSON.stringify(newUser), res => getUsers(newList => loadUserTable(newList)))

const deleteUser = (username) => {
    var data = {username: username}
    if (confirm("Are you sure you want to delete this user from rundeck? this cannot be undone")){
        makeRequest("DELETE", JSON.stringify(data), res => getUsers(newList => loadUserTable(newList)))
    }
}

const loadUserTable = (users) => {
    table.innerHTML = ""
    users.map((user, i) => {
        var roles = ""
        user.roles.map((role) => {
            roles += `<span class="custom-badge badge badge-secondary">${role}</span>`                
        })

        table.innerHTML += `
            <tr class="">
                <td class="">${i+1}</td>
                <td class="">${user.username}</td>
                <td class="">${roles}</td>
                <td>
                    <button class="btn btn-outline-dark custom-button" data-user="${i}"
                    data-toggle="modal" data-target="#user-data">Edit</button>
                    <button class="btn btn-outline-danger custom-button" onclick="deleteUser('${user.username}')">Delete</button>
                </td>
            </tr>
        `
    })
}

const searchHandler = (query) => {
    loadUserTable(usersList.filter(user => user.username.includes(query.toLowerCase())))
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

    if (i !== undefined) {
        var user = usersList[i]
        document.querySelector("#modal-username").value = user.username
        document.querySelector("#modal-roles").value = user.roles
        document.querySelector("#modal-password").value = user.password
        document.querySelector("#modal-submit").setAttribute("onclick", `handleUpdateSubmit("${user.username}")`)
    }else {
        document.querySelector("#modal-username").value = ""
        document.querySelector("#modal-roles").value = "" 
        document.querySelector("#modal-password").value = ""
        document.querySelector("#modal-submit").setAttribute("onclick", "handleCreateSubmit()")
    }
})