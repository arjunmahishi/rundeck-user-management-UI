class Roles {
    constructor(currRoles, renderFunction){
        this.allRoles = ["role1", "role2", "role3", "role4"]
        this.currRoles = currRoles
        this.availRoles = []
        this.renderFunction = renderFunction || (() => {})
    }

    addRole(role) {
        this.currRoles = this.currRoles.concat(role)
        this.updateAvailRoles()
        this.renderFunction(this.getRoles())
    }

    removeRole(role) {
        let i = this.currRoles.indexOf(role);
        if (i > -1) {
            this.currRoles.splice(i, 1);
        }

        this.updateAvailRoles()
        this.renderFunction(this.getRoles())
    }

    updateAvailRoles() {
        this.availRoles = this.allRoles.filter((ele) => this.currRoles.indexOf(ele) < 0)        
    }

    getRoles() {
        return this.currRoles + ""
    }
}