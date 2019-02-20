class Roles {
    constructor(renderFunction){
        this.allRoles = ["user", "voice-basic", "voice-adv", "api-basic", "api-adv"]
        this.currRoles = []
        this.availRoles = []
        this.renderFunction = renderFunction || (() => {})
    }

    addRole(role) {
        this.currRoles = this.currRoles.concat(role)
        this.updateAvailRoles()
        this.renderFunction(this.currRoles, this.availRoles)
    }

    removeRole(role) {
        let i = this.currRoles.indexOf(role);
        if (i > -1) {
            this.currRoles.splice(i, 1);
        }

        this.updateAvailRoles()
        this.renderFunction(this.currRoles, this.availRoles)
    }

    updateAvailRoles() {
        this.availRoles = this.allRoles.filter((ele) => this.currRoles.indexOf(ele) < 0)        
    }

    getRoles() {
        return this.currRoles
    }

    init(currRoles) {
        this.currRoles = currRoles
        this.updateAvailRoles()
        this.renderFunction(currRoles, this.availRoles)
    }
}