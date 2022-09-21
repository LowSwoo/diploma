<template>
  <v-app>
    <TopNav></TopNav>
    <v-main class="">
      <MainView></MainView>
    </v-main>
    <!-- <Dialogs></Dialogs> -->
  </v-app>
</template>

<script>
import MainView from "./views/MainView.vue";
import TopNav from "./components/TopNav.vue";
// import Dialogs from "./components/Dialogs.vue";

export default {
  name: "App",
  data: () => ({
    //
  }),
  components: { MainView, TopNav,  },
};
</script>
