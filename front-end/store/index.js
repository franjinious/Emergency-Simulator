import Vuex from 'vuex'

export const urgences = new Vuex.Store({
    state(){
        return{
            //Semaine
            week : 1,
            //Jour
            day : 1,
            //Patients deja traites
            patientsTraites : 0,
            //Temps total de traitement dans chambre 1
            time1 : 0,
            //Progres
            progress1 : 10,
            //waiting list
            waitingList: [                {
                id:1,
                name:'CUS1',
                avatar: require('./379339-512.png'),
                state: 'waiting',
                patience: 100,
                come:false,
                timer: null
            },
                {
                    id:2,
                    name:'CUS2',
                    avatar: require('./379444-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:3,
                    name:'CUS3',
                    avatar: require('./379446-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
            ],
            //List de patient

            allPatList:[
                {
                    id:1,
                    name:'CUS1',
                    avatar: require('./379339-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:2,
                    name:'CUS2',
                    avatar: require('./379444-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:3,
                    name:'CUS3',
                    avatar: require('./379446-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:4,
                    name:'CUS4',
                    avatar: require('./379448-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:5,
                    name:'CUS5',
                    avatar: require('./iconfinder_Boss-3_379348.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:6,
                    name:'CUS6',
                    avatar: require('./iconfinder_Man-16_379485.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:7,
                    name:'CUS7',
                    avatar: require('./iconfinder_Rasta_379441.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:8,
                    name:'CUS8',
                    avatar: require('./379339-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:9,
                    name:'CUS9',
                    avatar: require('./379444-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:10,
                    name:'CUS10',
                    avatar: require('./379446-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:11,
                    name:'CUS11',
                    avatar: require('./379448-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:12,
                    name:'CUS12',
                    avatar: require('./iconfinder_Boss-3_379348.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:13,
                    name:'CUS13',
                    avatar: require('./iconfinder_Man-16_379485.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:14,
                    name:'CUS14',
                    avatar: require('./iconfinder_Rasta_379441.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:15,
                    name:'CUS15',
                    avatar: require('./iconfinder_Rasta_379441.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:16,
                    name:'CUS16',
                    avatar: require('./iconfinder_Man-16_379485.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:17,
                    name:'CUS17',
                    avatar: require('./iconfinder_Boss-3_379348.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:18,
                    name:'CUS18',
                    avatar: require('./379448-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                ]


        }
    }
})
