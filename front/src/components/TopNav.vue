<template>
  <v-app-bar app elevation="2">
    <v-container class="py-0 fill-height">
      <v-avatar class="mr-10" color="grey darken-3" size="32">L</v-avatar>

      <v-btn text @click="showCreateDialog = true"> Создать расчёт </v-btn>

      <v-spacer></v-spacer>

      <v-btn width="400" text>
        <v-progress-linear value="40"></v-progress-linear>
      </v-btn>
      <v-btn icon @click="showUploadDialog = true">
        <v-icon>mdi-cloud-upload</v-icon>
      </v-btn>

      <v-btn v-if="bucketName" class="outlined x-small elevation-0 " @click="RemoveBucket" text>
                    {{bucketName}}
                    <v-icon color="red darken-4" class="mx-2">mdi-delete-forever</v-icon>
        </v-btn>
      <CreateBucket :showCreateDialog="showCreateDialog" v-on:CloseCreateDialog="showCreateDialog = false"></CreateBucket>
      <UploadFile :bucketName="bucketName" :showUploadDialog="showUploadDialog" v-on:CloseUploadDialog="showUploadDialog = false"></UploadFile>
    </v-container>
  </v-app-bar>
</template>
<script>
  import CreateBucket from "./dialogs/CreateBucket.vue";
import UploadFile from "./dialogs/UploadFile.vue";
  export default {
    data: () => ({
      showCreateDialog: false,
      showUploadDialog: false,
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
    components: { CreateBucket, UploadFile },
  };
</script>
