<template>
  <v-app-bar app elevation="2">
    <v-container class="py-0 fill-height">
      <v-avatar class="mr-10" color="grey darken-3" size="32">L</v-avatar>

      <v-btn text @click="showCreateDialog = true"> Создать расчёт </v-btn>

      <v-spacer></v-spacer>

      <!-- <v-btn text width="400" v-if="uploadFilesProgress.length != 0" @click="uploadSnackbar=true">
          <v-progress-linear v-bind:value="uploadFilesProgress.reduce((sum,elem)=>{return sum + elem.uploaded},0)/uploadFilesProgress.length"></v-progress-linear>
        </v-btn> -->
      <v-btn width="400" text>
        <v-progress-linear value="40"></v-progress-linear>
      </v-btn>
      <!-- <v-btn v-if="bucketName" class="outlined x-small elevation-0 white" @click="RemoveBucket">
                    {{bucketName}}
                    <span class="material-icons">delete</span>
        </v-btn> -->
      <v-btn text> Загузить данные </v-btn>

      <v-btn v-if="bucketName" class="outlined x-small elevation-0 " @click="RemoveBucket">
                    {{bucketName}}
                    <v-icon color="red" class="mx-2">mdi-close-circle</v-icon>
        </v-btn>
      <CreateBucket :showCreateDialog="showCreateDialog" v-on:CloseCreateDialog="showCreateDialog = false"></CreateBucket>
    </v-container>
  </v-app-bar>
</template>
<script>
  import CreateBucket from "./dialogs/CreateBucket.vue";
  export default {
    data: () => ({
      showCreateDialog: false,
      bucketName: '',
    }),
    methods: {
      RemoveBucket() {
        this.$http
        .post("http://localhost:8080" + "/api/bucket/remove", {
          bucketName: this.bucketName,
        })
        .then(
          (response) => {
            this.$root.$emit("UpdateBucketList")
          },
          (response) => {
            console.log("err:", response);
          }
        );
      }
    },
    mounted() {
      this.$root.$on('ChangeCurrentBucket', (bucketName) => this.bucketName = bucketName)
    },
    components: { CreateBucket },
  };
</script>
