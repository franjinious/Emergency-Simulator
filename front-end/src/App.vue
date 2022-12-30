<template xmlns:progress="http://java.sun.com/xml/ns/javaee">
  <div class = "app">
    <header>
<!--      <img class="people-icon" style="width: 100%;" src="src/assets/images/ambulance.png" alt=""/>-->
      <div class="header-item date" :style="progress">
        <span>W{{week}} </span> &nbsp; <span>D{{day}}</span>
      </div>
      <div class="header-item patients">
        Nb de patients suivis: {{patientsTraites}}
      </div>

      <div class="image-box">
        <img class="image" src="./assets/images/ambulance.png">
      </div>

      <div>
        <div class="waiting-area">
          <div class="wait">
            <div class="waiter" v-for="wt in waitingList" :key="wt.id">
              <div class="people" :class="wt.state" >
                <img class="people-icon" style="width: 100%;"
                     :src="wt.avatar" alt=""/>
                <div class="label" >
                  Dans l'attente
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class = "input-box" style="margin-top: 5%;margin-left: auto; margin-right: auto ">
          <label for="country">La maladie du nouveau patient</label>
          <select v-model = "maladie" id="country" name="country" style="width: 100%">
            <option value="rhume">Rhume</option>
            <option value="fievre">Fievre</option>
            <option value="rupture des os">Rupture des os</option>
          </select>
          <div class="button" @click="createPatient" style="margin-bottom: auto">Ajouter</div>
        </div>

        <div class="Acceuil ac1" style="margin-top: 2%; margin-left: auto; margin-right: auto">Acceuil</div>
      </div>

    </header>

    <main>

      <div>
        <div class="Room Room1" style="margin-top: 5%">
          Room1
          <div style="display:inline-block;margin-bottom: 0%; margin-top:0; width: 100%;text-align: center; align-items: baseline" >
            <progress
              :value = "progress1" max = 100 style=" width: 60%">
            </progress>
            <br style="float: right">Proges:{{progress1}}%<br>
          </div>

        </div>
        <div class="Room Room2" style="margin-top: 5%">Room2</div>
      </div>


      <div>
        <div class="Room Room3" style="margin-top: 5%; margin-bottom: 5%">Room3</div>
        <div class="Room Room4" style="margin-top: 5%; margin-bottom: 5%">Room4</div>
      </div>

    </main>
  </div>
  <welcome></welcome>
  <img ref="conf0" :src="backgroundImage" style="position:absolute;width: 100%;height: 100%;left: 0;top: 0;z-index: -1" alt="" >

</template>

<script>
import Welcome from "./components/welcomePop";
import {mapState} from 'vuex'
import axios from 'axios'

export default {
  name: 'App',
  components: {
    Welcome
  },
  data(){
    return{
      infos: null,
      backgroundImage: require('/src/assets/images/2382727.jpg'),
      backColorMap: {
        free: 'background: linear-gradient(to right, #dddddd 50%, #aaaaaa 50%);',
        appease: 'background: linear-gradient(to right, #661a00 50%, #401000 50%);',
        safe: 'background:  #00b200;',
        warning: 'background: linear-gradient(to right,#ff9122 50%,#d96d00  50%);',
        danger: 'background: linear-gradient(to right,#ff2020 50%, #b20000 50%);',
        destroy: 'background-color: #535362;text-decoration: line-through;'
      },
      maladie : "init"

    }
  },

  methods: {

    //Fonction d'essai pour pousser l'info du patient au backend
    createPatient() {
      axios.post('http://localhost:8080/api/users', {
        maladie: this.maladie,
      })
          .then(response => {
            // 创建成功，将新用户添加到用户列表中
            console.log(response.data);
          });
    },
    // Fonction pour rafraichir les infos toutes les 0.5 secondes
    mounted() {
      setInterval(() => {
        axios.get('http://example.com/api/data')
            .then(response => {
              this.infos = response.data;

            });
      }, 500);
    },
  },

  computed: {
    ...mapState([
        'hour',
        'minute',
        'day',
        'week',
        'patientsTraites',
        'time1',
        'progress1',
        'waitingList',
        'allPatListList'
    ]),
    progress(){
      let pro = (this.hour*60+this.minute)/240*100
      return 'background: linear-gradient(to right, #ffd24d '+pro+'%, #ffe699 '+pro+'%);'
    },

    // name:'createPatient',
    // data(){
    //   return{
    //     gravite : 0
    //   }
    // },
  }
}
</script>

<style>
#app {
  margin: 0;
  padding:0;
  width: 100%;
  height: 100%;
  box-sizing: border-box;
}

.header-item {
  background: linear-gradient(to bottom, #ffe699 50%, #ffd24d 50%);
  line-height: 1.8rem;
  width: 20%;
  border: 0.3rem solid #8c6900;
  border-radius: 99999px;
  font-size: 1rem;
  /*display: inline-flex;*/
  position: relative;
  margin-bottom: 10%;
}
/*Le temps passe*/
.date{
  margin-left:auto;
  margin-right:auto;
  position: absolute;
  right: 0;
}

/*Le nombre de patients traites*/
.patients{
  margin-left:auto;
  margin-right:auto;
  position: absolute;
  left: 0;
}


.image-box{
  height: 100px;
  width: 100px;
  background-color: transparent;
  margin-left: auto;
  margin-right: auto;
}
.image {
  height: 100px;
  width: 100px;
  /*background-image: url("assets/images/ambulance.png");*/
  background-repeat: no-repeat;
  max-width: 100%;
  max-height: 100%;
  background-size: contain;
  display: inline-block;
  /*background-position: top;*/
}

/*Room*/
main .Room {
  background-color: #ffa64d;
  margin: auto;
  height: 10rem;
  width: 40%;
  /*left: 50%;*/
  /*top: 70%;*/
  /*float: left;*/
  border-radius: 0.8rem;
  border: 0.3rem solid white;
  box-sizing: border-box;
  display: flex;
  flex-wrap: wrap;
  justify-content: left;
  align-items: flex-start;
  text-align: start;
  padding-bottom: 1.5rem;
  /*display: inline-flex;*/

}

main .Room1 {
  float: left;
}

main .Room3 {
  float: left;
}

main .Room2 {
  float: right;
}

main .Room4 {
  float: right;
}

 .Acceuil {
  /*margin-left: auto;*/
  /*margin-left: auto;*/
  /*align-items: center;*/
  margin-top: 50px;
  margin-left: auto;
  margin-right: auto;
  background-color: #ffd24d;
  height: 10rem;
  width: 40%;
  /*left: 50%;*/
  /*top: 70%;*/
  /*float: left;*/
  border-radius: 0.8rem;
  border: 0.3rem solid white;
  box-sizing: border-box;
  display: flex;
  flex-wrap: wrap;
  justify-content: left;
  align-items: normal;
  text-align: start;
  padding-bottom: 1.5rem;

}

 /*Espace pour les entrees*/
.input-box{
  height: 150px;
  width: 10%;
  background-color: #ffd24d;
  margin-left: auto;
  margin-right: auto;
  border-radius: 0.8rem;
  border: 0.3rem solid white;
  padding-bottom: 1.5rem;

}

/*Espaces pour renseigner les infos du noveau patient*/
input[type=text], select {
  width: 40%;
  padding: 12px 20px;
  margin: 8px 0;
  display: inline-block;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
  align-items: center;
  margin-left: auto;
  margin-right: auto;

}
/*Bouton pour envoyer les infos*/
input[type=submit] {
  width: 20%;
  background-color: #4CAF50;
  color: white;
  padding: 14px 20px;
  margin: 8px 0;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  float: right;
}

/*barre de progression*/
.progress-bar{
  height: 2rem;
  background-color: #ac91ff;
}
/*Zone d'attente*/
.waiting-area {
  position: relative;
  width: 95%;
  margin: 0 auto;
}

.waiting-area .wait {
  display: flex;
  flex-direction: row-reverse;
  position: absolute;
  right: 0;
  transform: translateY(-50%);
}

.waiting-area .wait .waiter {
  width: 3.5rem;
  transform: translateX(-50%);
}


/*.button{*/
/*  display: flex;*/
/*  justify-content: center;*/
/*  align-items: center;*/
/*  !*float: center;*!*/
/*  width: 30%;*/
/*  margin-left:auto;*/
/*  margin-right:auto;*/
/*}*/


</style>
