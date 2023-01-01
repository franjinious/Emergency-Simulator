<template xmlns:progress="http://java.sun.com/xml/ns/javaee">
  <div class = "app">
    <header style="margin-bottom: 20px">
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

        <div style="display: flex; flex-wrap: wrap">
          <div class="waiting-area" style="float: right">
            <div class="wait" style="width: 60%">
              <div class="waiter" v-for="wt in waitingList" :key="wt.id">
                <div class="people" :class="wt.state" >
                  <img class="people-icon" style="width: 100%;"
                       :src="wt.avatar" alt=""/>
                  <div class="label" :style="calculateWaitStyle(wt)">
                    Waiting
                  </div>
                </div>
              </div>
            </div>
          </div>
        <div class = "input-box" style="margin-top: 5%;margin-left: 0; margin-right: auto ">
          <label for="country">La maladie du nouveau patient</label>
          <select v-model = "maladie"  style="width: 100%">
            <option value=1>Rhume</option>
            <option value=2>Fievre</option>
            <option value=3>Mal à l'estomac</option>
            <option value=4>Diarrhée</option>
            <option value=5>Rupture des os</option>
<!--            心绞痛-->
            <option value=6>Angine</option>
<!--            中风-->
            <option value=7>Accident vasculaire cérébral</option>
            <option value=8>Cracher du sang</option>
<!--            休克-->
            <option value=9>Choc</option>
<!--            心脏骤停-->
            <option value=10>Arrêt cardiaque</option>
          </select>
          <div class="button" @click="createPatient" style="margin-bottom: auto">Ajouter</div>
        </div>
        </div>

      <div>
        <div >
          <div style="width: 25%; float: left">
            <div class="Accueil ac1" v-if="nbAccueil>=1" style="margin-top: 2%;height: 6rem; width:50%; float: left">
              Accueil 1
              <div class="people p1 free" >
                <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
                <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
                <div class="redC"></div>
                <div class="redR"></div>
              </div>

              <!--              <div class="people cus" :class="cus.state">-->
<!--                <img class="people-icon" style="width: 100%" :src="cus.avatar"/>-->
<!--              </div>-->

            </div>
            <div class="Accueil ac2" v-if="nbAccueil>=2" style="margin-top: 2%;height: 6rem; width:50%;float: left">Accueil 2          <div class="people p1 free" >
              <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
              <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
              <div class="redC"></div>
              <div class="redR"></div>
            </div>
            </div>
            <div class="Accueil ac3" v-if="nbAccueil>=3" style="margin-top: 2%;height: 6rem; width:50%;float: left">Accueil 3          <div class="people p1 free" >
              <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
              <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
              <div class="redC"></div>
              <div class="redR"></div>
            </div>
            </div>
            <div class="Accueil ac4" v-if="nbAccueil>=4" style="margin-top: 2%;height: 6rem; width:50%;float: left">Accueil 4          <div class="people p1 free" >
              <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
              <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
              <div class="redC"></div>
              <div class="redR"></div>
            </div>
            </div>
            <div class="Accueil ac5" v-if="nbAccueil>=5" style="margin-top: 2%;height: 6rem; width:50%;float: left">Accueil 5          <div class="people p1 free" >
              <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
              <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
              <div class="redC"></div>
              <div class="redR"></div>
            </div>
            </div>
            <div class="button activerAccueil" @click="activerAccueil" style="float: left; margin-top: auto; margin-bottom: auto">➕</div>
            <div class="button desactiverAccueil" @click="desactiverAccueil"  style="float: left; margin-top: auto; margin-bottom: auto">➖</div>
          </div>

          <div></div>

          <div style="width: 30%; float: right; text-align: end">
            <div class="Accueil infi1" v-if="nbInfirmier>=1" style="margin-top:1.5%;height: 5.5rem;float:right;display: block">
              <div class="people p1 free" v-if="nbInfirmier>=1" style="float: left;display: inline-block" >
                <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
                <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
                <div class="redC"></div>
                <div class="redR"></div>
              </div>
              Reception Infirmiere 1
            </div>

            <div class="Accueil infi2" v-if="nbInfirmier>=2" style="margin-top:1.5%;height:5.5rem;float: right; display: block">
              <div class="people p1 free" v-if="nbInfirmier>=2" style="float: left;display: inline-block" >
                <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
                <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
                <div class="redC"></div>
                <div class="redR"></div>
              </div>

              Reception Infirmiere 2
            </div>
<!--            <div class="Accueil infi2" v-else style="margin-top:1.5%;height:3rem; {background: #aaaaaa}" >Reception Infirmiere 2</div>-->

            <div class="Accueil infi3" v-if="nbInfirmier>=3" style="margin-top:1.5%;height: 5.5rem;float: right;display: block">
              <div class="people p1 free" v-if="nbInfirmier>=3" style="float: left;display: inline-block" >
                <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
                <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
                <div class="redC"></div>
                <div class="redR"></div>
              </div>
              Reception Infirmiere 3
            </div>

            <div class="Accueil infi4" v-if="nbInfirmier>=4" style="margin-top:1.5%;height: 5.5rem;float:right;display: block">
              <div class="people p1 free" v-if="nbInfirmier>=4" style="float: left;display: inline-block" >
                <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
                <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
                <div class="redC"></div>
                <div class="redR"></div>
              </div>
              Reception Infirmiere 4
            </div>

            <div class="Accueil infi5" v-if="nbInfirmier>=5" style="margin-top: 1.5%;height: 5.5rem;float:right;display: block">
              <div class="people p1 free" v-if="nbInfirmier>=1" style="float: left;display: inline-block" >
                <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
                <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
                <div class="redC"></div>
                <div class="redR"></div>
              </div>
              Reception Infirmiere 5
            </div>
            <div class="button activerInfirmier" @click="activerInfirmier" style="float: left">➕</div>
            <div class="button desactiverInfirmier" @click="desactiverInfirmier" style="float: left">➖</div>
          </div>
        </div>

      </div>

    </header>

    <main style="margin-top: 20px; width: 100%">
      <div>
        <div class="salleAttente" style="width: 100%; margin-top: 2rem">
          Salle d'attente
        </div>
      </div>
      <div style="display: flex; flex-wrap: wrap">
        <div style="width:70%; margin-left: 0;">
        <div class="Room Room1" v-if="nbSalle>=1" style="float:left;">
          Room1
          <div class="people p1 free" >
            <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
            <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
            <div class="redC"></div>
            <div class="redR"></div>
          </div>
        </div>
        <div class="Room Room2" v-if="nbSalle>=2" style="float:left;margin-left: 2rem">
          Room2
          <div class="people p1 free" >
            <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
            <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
            <div class="redC"></div>
            <div class="redR"></div>
          </div>
        </div>
        <div class="Room Room3" v-if="nbSalle>=3" style="float:left;margin-left: 2rem">
          Room3
          <div class="people p1 free" >
            <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
            <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
            <div class="redC"></div>
            <div class="redR"></div>
          </div>
        </div>
        <div class="Room Room4" v-if="nbSalle>=4" style="float:left;margin-left: 2rem">
          Room4
          <div class="people p1 free" >
            <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
            <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
            <div class="redC"></div>
            <div class="redR"></div>
          </div>
        </div>
        <div class="Room Room5" v-if="nbSalle>=5" style="float:left;margin-left: 2rem">
          Room5
          <div class="people p1 free" >
            <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
            <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
            <div class="redC"></div>
            <div class="redR"></div>
          </div>
        </div>
        <div class="Room Room6" v-if="nbSalle>=6" style="float:left;margin-left: 2rem">
          Room6
          <div class="people p1 free" >
            <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
            <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
            <div class="redC"></div>
            <div class="redR"></div>
          </div>
        </div>
        <div class="Room Room7" v-if="nbSalle>=7" style="float:left;margin-left: 2rem">
          Room7
          <div class="people p1 free" >
            <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
            <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
            <div class="redC"></div>
            <div class="redR"></div>
          </div>
        </div>
        <div class="Room Room8" v-if="nbSalle>=8" style="float:left;margin-left: 2rem">
          Room8
          <div class="people p1 free" >
            <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
            <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
            <div class="redC"></div>
            <div class="redR"></div>
          </div>
        </div>
        <div class="Room Room9" v-if="nbSalle>=9" style="float:left;margin-left: 2rem">
          Room9
          <div class="people p1 free" >
            <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
            <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
            <div class="redC"></div>
            <div class="redR"></div>
          </div>
        </div>
        <div class="Room Room10" v-if="nbSalle>=10" style="float:left;margin-left: 2rem">
          Room10
          <div class="people p1 free" >
            <img style="width: 100%;filter: grayscale(100%) brightness(500%);" class="people-icon">
            <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
            <div class="redC"></div>
            <div class="redR"></div>
          </div>
        </div>
        <div class="button activerSalle" @click="activerSalle" style="float: left">➕</div>
        <div class="button desactiverSalle" @click="desactiverSalle" style="float: left">➖</div>
        </div>
        <div style="width: 30%">
          <div style="width: 100%; justify-items: center; float: right">
          <div class="Accueil nbMedecin" style=" margin-right: auto; margin-left: auto;background: #ffa64d;height: 8rem">
            Le nombre de medecin actifs:{{nbMedecin}}
            <div class="button activerDoc" @click="activerDoc" style="float: left">➕</div>
            <div class="button desactiverDoc" @click="desactiverDoc" style="float: left">➖</div>

          </div>
        </div>
        </div>
      </div>


      <div>
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
      //Les infos ensemble passees
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
      maladie : 0,
      nbAccueil : 1,
      nbInfirmier : 1,
      nbSalle : 1,
      nbMedecin:9,
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
    // 计算顾客耐心值
    calculateWaitStyle(wt) {
      return 'background: linear-gradient(to right, #006dd9 ' + (100 - wt.patience) + '%, #2693ff ' + (100 - wt.patience) + '%);'
    },
    activerAccueil(){
      if (this.nbAccueil == 5){
        this.nbAccueil = 5
      }else
        this.nbAccueil ++
    },

    desactiverAccueil(){
      if (this.nbAccueil == 1){
        this.nbAccueil = 1
      }else
        this.nbAccueil --
    },
    activerInfirmier(){
      if (this.nbInfirmier == 5){
        this.nbInfirmier = 5
      }else
        this.nbInfirmier ++
    },

    desactiverInfirmier(){
      if (this.nbInfirmier == 1){
        this.nbInfirmier = 1
      }else
        this.nbInfirmier --
    },

    activerSalle(){
      if (this.nbSalle == 10){
        this.nbSalle = 10
      }else
        this.nbSalle ++
    },
    desactiverSalle(){
      if (this.nbSalle == 1){
        this.nbSalle = 1
      }else
        this.nbSalle --
    },
    activerDoc(){

    },
    desctiverDoc(){

    }


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
        'allPatListList',
        'cntInfirmier',
        'color1',
        'color2',
        'nbMedecin'
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
  height: 5rem;
  width: 8%;  /*left: 50%;*/
  /*top: 70%;*/
  /*float: left;*/
  border-radius: 0.8rem;
  border: 0.3rem solid white;
  box-sizing: border-box;
  display: inline-block;
  /*flex-wrap: wrap;*/
  justify-content: left;
  align-items: flex-start;
  text-align: start;
  margin-bottom: 1%;
  /*padding-bottom: 2rem;*/
  /*display: inline-flex;*/

}

/*main .Room1 {*/
/*  float: left;*/
/*}*/

/*main .Room3 {*/
/*  float: left;*/
/*}*/

/*main .Room2 {*/
/*  float: right;*/
/*}*/

/*main .Room4 {*/
/*  float: right;*/
/*}*/
.buttonArea{
  margin-top: 3rem;
  width: 40%;
  display: inline-block;
}
 .Accueil {
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
 .salleAttente{
   background: chocolate;
   margin: auto;
   height: 5rem;
   width: 10%;  /*left: 50%;*/
   /*top: 70%;*/
   /*float: left;*/
   border-radius: 0.8rem;
   border: 0.3rem solid white;
   box-sizing: border-box;
   display: inline-block;
   /*flex-wrap: wrap;*/
   justify-content: left;
   align-items: flex-start;
   text-align: start;
   margin-bottom: 1%;

 }

 .activerAccueil{
   margin-top: 3rem;
 }

 .desactiverAccueil{
   margin-left: 1rem;
 }

 .activerDoc{
   height: 40%;
 }
 .desactiverDoc{
  height: 40%;
}

 /*Espace pour les entrees*/
.input-box{
  height: 80px;
  width: 20%;
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
  width: 80%;
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
