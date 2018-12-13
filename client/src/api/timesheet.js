// const _consultants = [
//     {
//         _id: 1,
//         name: 'First last'
//     }
// ]

// const _reportedHours = [
//     {
//         _id: 1,
//         date: '2018-12-04',
//         hours: 8,
//         project: 'Project',
//         description: 'Description',
//         rate: 'Off-site',
//         consultant: 1
//     },
// ]

import axios from 'axios'

export default {
    getReportedHours (rh) {
        axios.get(`http://localhost:3000/api/reported/all`, { crossDomain: true })
            .then(response => {
                rh(response.data)
            })
            .catch(e => {
                console.log(e) /* eslint-disable-line no-console */
            })
    },
    getConsultants (cs) {
        axios.get(`http://localhost:3000/api/consultants/list`, { crossDomain: true })
            .then(response => {
                cs(response.data)
            })
            .catch(e => {
                console.log(e) /* eslint-disable-line no-console */
            })
    }
}
