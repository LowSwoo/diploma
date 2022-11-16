<template>
  <v-app>
    <TopNav v-on:toggleNavigationDrawer="navigationDrawer = !navigationDrawer"></TopNav>
    <LeftNav :navigationDrawer="navigationDrawer" v-on:showDescriptionDialog="setBucketDescription"></LeftNav>
    <!-- <Rightbar></Rightbar> -->
    <v-main >
      <CreateBucketDescription :bucketName="bucketD" :showCreateDescriptionDialog="showDescriptionDialog" v-on:CloseDescriptionDialog="showDescriptionDialog = false"></CreateBucketDescription>
      <MainSheet></MainSheet>
    </v-main>
  </v-app>
</template>

<script>
import TopNav from "./components/TopNav.vue";
import LeftNav from "./components/LeftNav.vue";
import MainSheet from "./components/MainSheet.vue";
import Rightbar from "./components/Rightbar.vue";
import CreateBucketDescription from "./components/dialogs/CreateBucketDescription.vue";
// import Dialogs from "./components/Dialogs.vue";

export default {
  name: "App",
  data: () => ({
    navigationDrawer: false,
    bucketD: "",
    showDescriptionDialog: false,
  }),
  methods: {
    setBucketDescription(bucket) {
      this.bucketD = bucket
      this.showDescriptionDialog = true
    }
  },
  components: { TopNav, LeftNav, MainSheet, Rightbar, CreateBucketDescription },
};
</script>
