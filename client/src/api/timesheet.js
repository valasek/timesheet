// const _consultants = [    
//     {
//         _id: 1,
//         name: 'Stanislav Valášek'
//     },
//     {
//         _id: 2,
//         name: 'Jakub Barták'
//     },
//     {
//         _id: 3,
//         name: 'Jan Ferdinand Polášek'
//     },
//     {
//         _id: 4,
//         name: 'Jaroslav Barannikov'
//     },
//     {
//         _id: 5,
//         name: 'Zoroslav Kunčák'
//     },
//     {
//         _id: 6,
//         name: 'Radek Ocelák'
//     },
// ]

const _reportedHours = [    
    {
        _id: 1,
        date: '2018-12-04',
        hours: 8,
        project: 'Voya',
        description: 'Member mastering and deployment',
        rate: 'Off-site',
        consultant: 1
    },
    {
        _id: 2,
        date: '2018-12-04',
        hours: 1,
        project: '_Training',
        description: 'English lession',
        rate: 'Off-site',
        consultant: 1
    },
    {
        _id: 3,
        date: "2018-12-05",
        hours: 2,
        project: 'CCNC',
        description: 'Member - NULL attributes on master',
        rate: 'Off-site',
        consultant: 2
    },
    {
        _id: 4,
        date: '2018-12-06',
        hours: 3,
        project: 'Servus',
        description: 'Credential - names cleansing',
        rate: 'Off-site',
        consultant: 2
    },
    {
        _id: 5,
        date: '2018-12-06',
        hours: 8,
        project: 'Target',
        description: 'Fix rule XZ bla',
        rate: 'Off-site',
        consultant: 3
    }
]

import axios from 'axios'

export default {
    getReportedHours ( rh ) {
        setTimeout(() => rh(_reportedHours), 100)
    },
    getConsultants ( cs ) {
        axios.get(`http://localhost:3000/api/consultants/list`, {crossDomain: true})
            .then(response => {
                cs(response.data)
            })
            .catch(e => {
                console.log(e) /* eslint-disable-line no-console */
            })
    }
}