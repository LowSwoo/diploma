<template>
  <v-app>
    <TopNav v-on:toggleNavigationDrawer="navigationDrawer = !navigationDrawer"></TopNav>
    <LeftNav :navigationDrawer="navigationDrawer"></LeftNav>
    <!-- <Rightbar></Rightbar> -->
    <v-main>
      <MainSheet></MainSheet>
    </v-main>
  </v-app>
</template>

<script>
import TopNav from "./components/TopNav.vue";
import LeftNav from "./components/LeftNav.vue";
import MainSheet from "./components/MainSheet.vue";
import Rightbar from "./components/Rightbar.vue";
// import Dialogs from "./components/Dialogs.vue";

export default {
  name: "App",
  data: () => ({
    navigationDrawer: true,
    //
  }),
  components: { TopNav, LeftNav, MainSheet, Rightbar },
};
</script>
