<template>
  <v-container>
    <v-row>
      <LeftNav></LeftNav>
      <MainSheet></MainSheet>
    </v-row>
  </v-container>
</template>

<script>
import LeftNav from "@/components/LeftNav.vue";
import MainSheet from "@/components/MainSheet.vue";

export default {
  name: "MainView",

  components: {
    LeftNav,
    MainSheet,
  },
};
</script>
