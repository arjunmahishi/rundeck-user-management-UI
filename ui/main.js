const baseURL = "http://localhost:3000/users"
const table = document.querySelector("#user-table")

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
       callback(JSON.parse(res)) 
    })
}

const loadUserTable = (users) => {
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
                    <button class="btn btn-outline-secondary">Edit</button>
                    <button class="btn btn-outline-danger">Delete</button>
                </td>
            </tr>
        `
    })
}

window.onload = () => {
    getUsers(users => loadUserTable(users))
}