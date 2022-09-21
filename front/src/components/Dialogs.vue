<template>
  <v-container>
    <CreateBucket></CreateBucket>
    <UploadFile></UploadFile>
    <UploadSnackbar></UploadSnackbar>
  </v-container>
</template>

<script>
import CreateBucket from "./dialogs/CreateBucket.vue";
import UploadFile from "./dialogs/UploadFile.vue";
import UploadSnackbar from "./dialogs/UploadSnackbar.vue";

export default {
    name: 'Dialogs'
}

</script>
