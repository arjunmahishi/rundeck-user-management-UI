window.onload = () => {
    getUsers(users => loadUserTable(users))
}

const baseURL = "http://localhost:3000/users"
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

const loadUserTable = (users) => {
    table.innerHTML = ""
    table.innerHTML += users.map((user, i) => {
        var roles = ""
        user.roles.map((role) => {
            roles += `<span class="custom-badge badge badge-dark">${role}</span>`                
        })

        return `
            <tr class="">
                <td class="">${i+1}</td>
                <td class="">${user.username}</td>
                <td class="">${roles}</td>
                <td>
                    <button class="btn btn-outline-secondary" data-user="${i}"
                    data-toggle="modal" data-target="#user-data">Edit</button>
                    <button class="btn btn-outline-danger">Delete</button>
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

$('#user-data').on('show.bs.modal', function (event) {
    var button = $(event.relatedTarget)
    var i = button.data('user')

    if (i !== undefined) {
        var user = usersList[i]
        document.querySelector("#modal-username").value = user.username
        document.querySelector("#modal-roles").value = user.roles
        document.querySelector("#modal-password").value = user.password
    }else {
        document.querySelector("#modal-username").value = ""
        document.querySelector("#modal-roles").value = "" 
        document.querySelector("#modal-password").value = ""
    }
})