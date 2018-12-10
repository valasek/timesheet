// initial state
const state = {
    page: 'Time Sheet',
    monday: getMonday(new Date()),
    sunday: getSunday(new Date()),
}

const getters = {}

const actions = {
    // getPage ({ commit, page }) {
    //     commit('setPage', page)
    // }
}

const mutations = {

    setPage ( state, page ) {
        state.page = page
    },

    setWeek ( state, direction ) {
        var currentMonday = new Date(state.monday)
        var targetMonday = currentMonday
        switch (direction) {
            case 'previous':
                targetMonday.setDate(currentMonday.getDate() - 7)
                break;
            case 'next':
                targetMonday.setDate(currentMonday.getDate() + 7)
                break;
            default:
                return
        }
        var targetSunday = new Date()
        targetSunday.setDate(targetMonday.getDate() + 6)
        state.monday = targetMonday
        state.sunday = targetSunday
    }
}

function getMonday ( date ) {
    var day = date.getDay() || 7
    if( day !== 1 ) 
        date.setHours(-24 * (day - 1))
    return date
}

function getSunday ( date ) {
    var day = date.getDay() || 7
    if( day !== 7 ) 
        date.setHours(-24 * (day - 1))
    return date
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}